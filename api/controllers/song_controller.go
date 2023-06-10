package controllers

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/dhowden/tag"
	"github.com/jbenet/go-base58"
	"github.com/science-engineering-art/gotify/api/models"
	"github.com/science-engineering-art/gotify/api/net"
	"github.com/science-engineering-art/gotify/api/responses"

	trackerUtils "github.com/science-engineering-art/gotify/tracker/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateSong(c *fiber.Ctx) error {
	fmt.Println("Init CreateSong()")

	// get file from the multipart-form
	fileForm, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
		return c.Status(http.StatusCreated).
			JSON(
				responses.SongResponse{
					Status:  http.StatusBadRequest,
					Message: "You must send in the form data the music file, with the key `file`.",
					Data:    &fiber.Map{"success": false},
				},
			)
	}
	filename := fileForm.Filename

	// create a temporal file with the received file
	err = c.SaveFile(fileForm, fmt.Sprintf("./%s", filename))
	if err != nil {
		return c.Status(http.StatusCreated).
			JSON(
				responses.SongResponse{
					Status:  http.StatusInternalServerError,
					Message: "failed",
					Data:    &fiber.Map{"success": false},
				},
			)
	}
	// then remove it
	defer os.Remove(filename)

	// open the temporal file
	file, err := os.Open(filename)
	if err != nil {
		return c.Status(http.StatusCreated).
			JSON(
				responses.SongResponse{
					Status:  http.StatusInternalServerError,
					Message: "failed",
					Data:    &fiber.Map{"success": false},
				},
			)
	}
	// when it finish, close it
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)

	// keep in a buffer the file information
	file.Read(buffer)

	key, err := net.Peer.Store(&buffer)
	if err != nil {
		return c.Status(http.StatusCreated).
			JSON(
				responses.SongResponse{
					Status:  http.StatusInternalServerError,
					Message: "failed",
					Data:    &fiber.Map{"success": false},
				},
			)
	}

	// if net.Peer == nil {
	// 	fmt.Printf("?????\n?????\n?????\n?????\n?????\n?????\n?????\n")
	// }
	// // check if was correctly created
	// bufferSong, _ := net.Peer.GetValue(key, 0, 0)
	// fmt.Printf("?????\n?????\n?????\n?????\n?????\n?????\n?????\nPeer Saved a value with len: %d?????\n?????\n?????\n?????\n?????\n?????\n?????\n", len(bufferSong))
	// os.WriteFile("received_song.mp3", bufferSong, 0600)

	songBytes := bytes.NewReader(buffer)

	m, err := tag.ReadFrom(songBytes)
	if err != nil {
		return err
	}

	// Store metadata section
	jsonMap := make(map[string]string)
	jsonMap["artist"] = m.Artist()
	jsonMap["album"] = m.Album()
	jsonMap["genre"] = m.Genre()
	jsonMap["title"] = m.Title()

	jsonString, _ := json.Marshal(jsonMap)
	fmt.Println(string(jsonString))
	// Get datahash
	hash := sha1.Sum(buffer)
	songDataHash := base58.Encode(hash[:])

	net.Tracker.StoreSongMetadata(string(jsonString), songDataHash)

	fmt.Printf("Song ID Created: %s\n", key)

	return c.Status(http.StatusCreated).
		JSON(
			responses.SongResponse{
				Status:  http.StatusCreated,
				Message: "success",
				Data:    &fiber.Map{"success": true},
			},
		)
}

func GetSongById(c *fiber.Ctx) error {
	fmt.Println("==> INIT GetSongById()")
	defer fmt.Println("==> EXIT GetSongById()")

	// get the song ID
	songId := c.Params("songId")

	fmt.Println("SongId:", songId)

	rangeHeader := c.Get("Range")

	rangePattern := `bytes=(\d+)-(\d*)`
	re := regexp.MustCompile(rangePattern)
	matches := re.FindStringSubmatch(rangeHeader)

	var start int64
	var end int64

	if len(matches) == 3 {
		start, _ = strconv.ParseInt(matches[1], 10, 64)
		end, _ = strconv.ParseInt(matches[2], 10, 64)
	} else {
		fmt.Println("==> ERROR `Invalid or missing Range header`")
		return errors.New("invalid or missing range header")
	}

	if end == 0 {
		end = math.MaxInt64
	}
	fmt.Printf("Range of %s Start: %d <==> End: %d\n", songId, start, end)

	song := []byte{}

	for i := start; i < end; i += 1048576 {
		j := math.Min(float64(i+1048576), float64(end))

		fmt.Println("Before Peer.GetValue()")
		songChunck, err := net.Peer.GetValue(songId, int64(i), int64(j))
		if err != nil {
			fmt.Printf("==> ERROR %s\n", err)
			return nil
		}
		fmt.Printf("After Peer.GetValue(%s, %d, %d) and Received: %d\n", songId, i, int(j), len(songChunck))
		song = append(song, songChunck...)

		if int64(len(songChunck)) < int64(j)-i {
			break
		}
	}

	// // Get the unique identifier from the request context
	// requestId, ok := c.Context().UserValue("requestId").(uuid.UUID)
	// if !ok {
	// 	// Handle error if unique identifier cannot be obtained
	// 	fmt.Println("==> ERROR `unique identifier cannot be obtained`")
	// 	return fiber.NewError(fiber.StatusInternalServerError, "Unique identifier could not be obtained")
	// }

	// fileName := "./public/tmp.mp3"

	// os.WriteFile(fileName, song, 0600)
	// defer os.Remove(fileName)

	// c.Response().Header.SetContentRange(start, end, len(song))
	// c.Response().Header.SetContentLength(len(song))

	// c.Context().SetStatusCode(fiber.StatusPartialContent)

	fmt.Println("==> OKKK", len(song))
	return c.Send(song)
}

// func UpdateSong(c *fiber.Ctx) error {

// 	songId := c.Params("songId")
// 	pbSongId := &pb.SongId{Id: songId}

// 	pbSongMetadata := new(pb.SongMetadata)

// 	if err := c.BodyParser(pbSongMetadata); err != nil {
// 		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
// 			"errors": err.Error(),
// 		})
// 	}

// 	pbUpdatedSong := pb.UpdatedSong{
// 		Id:       pbSongId,
// 		Metadata: pbSongMetadata,
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	pbResp, err := songClient.UpdateSong(ctx, &pbUpdatedSong)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(
// 			responses.SongResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "error",
// 				Data:    &fiber.Map{"data": err.Error()},
// 			},
// 		)
// 	}

// 	if !pbResp.Success {
// 		return c.Status(http.StatusOK).JSON(
// 			responses.SongResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "unsuccess",
// 				Data:    &fiber.Map{},
// 			},
// 		)
// 	}

// 	return c.Status(http.StatusOK).JSON(
// 		responses.SongResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    &fiber.Map{},
// 		},
// 	)
// }

// func RemoveSongById(c *fiber.Ctx) error {

// 	songId := c.Params("songId")

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	pbResp, err := songClient.RemoveSongById(ctx, &pb.SongId{Id: songId})
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(
// 			responses.SongResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "error",
// 				Data:    &fiber.Map{"data": err.Error()},
// 			},
// 		)
// 	}

// 	if !pbResp.Success {
// 		return c.Status(http.StatusOK).JSON(
// 			responses.SongResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "unsuccess",
// 				Data:    &fiber.Map{},
// 			},
// 		)
// 	}

// 	return c.Status(http.StatusOK).JSON(
// 		responses.SongResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    &fiber.Map{},
// 		},
// 	)
// }

func SongFilter(c *fiber.Ctx) error {

	query := new(models.SongQuery)

	if err := c.BodyParser(query); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	fmt.Println("Query", query)

	queryString := convertQueryToString(*query)
	fmt.Println("QueryString", queryString)

	keyHash := trackerUtils.GetJsonMetadataKeyHash(queryString)
	songsList := net.Tracker.GetSongList(keyHash)

	// songsList := net.Tracker.GetSongList(queryString)
	fmt.Println("SongList", songsList)

	var songsResponse []models.SongsFilterResponse

	for _, song := range songsList {
		songResponse := convertStringToResponse(song)
		songsResponse = append(songsResponse, songResponse)
	}
	fmt.Println("SongResponse", songsResponse)

	return c.Status(http.StatusOK).JSON(
		responses.SongResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"songs": songsResponse},
		},
	)
}

func convertQueryToString(query models.SongQuery) string {
	jsonBytes, err := json.Marshal(query)
	if err != nil {
		fmt.Println("error:", err)
		return "{}"
	}
	jsonString := string(jsonBytes)
	return jsonString
}

func convertStringToResponse(queryString string) models.SongsFilterResponse {
	songData := []byte(queryString)
	var response models.SongsFilterResponse
	err := json.Unmarshal(songData, &response)
	if err != nil {
		fmt.Println("Error while Unmarshaling")
	}
	return response
}
