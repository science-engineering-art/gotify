package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/science-engineering-art/gotify/peer/persistence"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"

	"github.com/science-engineering-art/gotify/api/responses"

	"github.com/gofiber/fiber/v2"
)

var (
	db   = persistence.NewMongoDb("admin", "songs")
	peer = kademlia.NewFullNode("0.0.0.0", 8080, 32140, db, false)
)

func CreateSong(c *fiber.Ctx) error {
	// get file from the multipart-form
	fileForm, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename := fileForm.Filename

	// create a temporal file with the received file
	err = c.SaveFile(fileForm, fmt.Sprintf("./%s", filename))
	if err != nil {
		return err
	}
	// then remove it
	defer os.Remove(filename)

	// open the temporal file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	// when it finish, close it
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)

	// keep in a buffer the file information
	file.Read(buffer)

	hash := sha1.Sum(buffer)
	key := base64.RawStdEncoding.EncodeToString(hash[:])

	fmt.Printf("Song ID Created: %s\n", key)
	peer.StoreValue(key, buffer)

	return c.Status(
		http.StatusCreated).JSON(
		responses.SongResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"success": true},
		},
	)
}

// func GetSongById(c *fiber.Ctx) error {
// 	// get the song ID
// 	songId := c.Params("songId")

// 	rangeHeader := c.Get("Range")

// 	rangePattern := `bytes=(\d+)-(\d*)`
// 	re := regexp.MustCompile(rangePattern)
// 	matches := re.FindStringSubmatch(rangeHeader)

// 	var start int
// 	var endStr string

// 	if len(matches) == 3 {
// 		start, _ = strconv.Atoi(matches[1])
// 		endStr = matches[2]
// 	} else {
// 		return errors.New("Invalid or missing Range header")
// 	}

// 	contentLength := 4 // !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// 	end := contentLength - 1

// 	if endStr != "" {
// 		end, _ = strconv.Atoi(endStr)
// 	}

// 	song, err := songClient.GetSongById(ctx, &pb.SongId{
// 		Id: songId,
// 	})

// 	if err != nil {
// 		return nil
// 	}

// 	// Get the unique identifier from the request context
// 	requestId, ok := c.Context().UserValue("requestId").(uuid.UUID)
// 	if !ok {
// 		// Handle error if unique identifier cannot be obtained
// 		return fiber.NewError(fiber.StatusInternalServerError, "Unique identifier could not be obtained")
// 	}

// 	fileName := fmt.Sprintf("./tmp_%s.mp3", requestId.String())

// 	os.WriteFile(fileName, song.RawSong.Buffer, 0600)
// 	defer os.Remove(fileName)

// 	return c.SendFile(fileName, true)
// }

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

// func FilterSongs(c *fiber.Ctx) error {

// 	query := new(pb.SongMetadata)

// 	if err := c.BodyParser(query); err != nil {
// 		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
// 			"errors": err.Error(),
// 		})
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	stream, err := songClient.FilterSongs(ctx, query)
// 	if err != nil {
// 		return err
// 	}

// 	var songs []models.SongDTO

// 	for {
// 		song, err := stream.Recv()
// 		if err != nil && err != io.EOF {
// 			return err
// 		}
// 		if song == nil {
// 			break
// 		}

// 		objID, err := primitive.ObjectIDFromHex(song.Id.Id)
// 		if err != nil {
// 			return err
// 		}

// 		songs = append(songs, models.SongDTO{
// 			Artist: *song.Metadata.Artist,
// 			Id:     objID,
// 			Title:  *song.Metadata.Title,
// 			Year:   int(*song.Metadata.Year),
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(
// 		responses.SongResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    &fiber.Map{"songs": songs},
// 		},
// 	)
// }
