package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/science-engineering-art/gotify/api/models"
	"github.com/science-engineering-art/gotify/api/net"
	"github.com/science-engineering-art/gotify/api/responses"

	"github.com/gofiber/fiber/v2"
)

func CreateSong(c *fiber.Ctx) error {
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

	fmt.Printf("Before Store().. len(data): %d\n", len(buffer))

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

	// Store metadata section
	jsonMap := make(map[string]string)
	jsonMap["title"] = c.FormValue("title")
	jsonMap["author"] = c.FormValue("author")
	jsonMap["album"] = c.FormValue("album")
	jsonMap["gender"] = c.FormValue("gender")

	jsonString, _ := json.Marshal(jsonMap)

	// Get datahash
	hash := sha1.Sum(buffer)
	songDataHash := base64.RawStdEncoding.EncodeToString(hash[:])

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

	rangeHeader := c.Get("Range")

	rangePattern := `bytes=(\d+)-(\d*)`
	re := regexp.MustCompile(rangePattern)
	matches := re.FindStringSubmatch(rangeHeader)

	var start int
	var endStr string

	if len(matches) == 3 {
		start, _ = strconv.Atoi(matches[1])
		endStr = matches[2]
	} else {
		fmt.Println("==> ERROR `Invalid or missing Range header`")
		return errors.New("invalid or missing range header")
	}

	contentLength := 4 // !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	end := contentLength - 1

	if endStr != "" {
		end, _ = strconv.Atoi(endStr)
	}

	song, err := net.Peer.GetValue(songId, int32(start), int32(end))
	if err != nil {
		fmt.Printf("==> ERROR %s\n", err)
		return nil
	}

	// Get the unique identifier from the request context
	requestId, ok := c.Context().UserValue("requestId").(uuid.UUID)
	if !ok {
		// Handle error if unique identifier cannot be obtained
		fmt.Println("==> ERROR `unique identifier cannot be obtained`")
		return fiber.NewError(fiber.StatusInternalServerError, "Unique identifier could not be obtained")
	}

	fileName := fmt.Sprintf("./tmp_%s.mp3", requestId.String())

	os.WriteFile(fileName, song, 0600)
	defer os.Remove(fileName)

	fmt.Println("==> OKKK")
	return c.SendFile(fileName, true)
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

	queryString := convertQueryToString(*query)
	songsList := net.Tracker.GetSongList(queryString)

	var songsResponse []models.SongsFilterResponse

	for _, song := range songsList {
		songResponse := convertStringToResponse(song)
		songsResponse = append(songsResponse, songResponse)
	}

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
