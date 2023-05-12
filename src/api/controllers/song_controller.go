package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/science-engineering-art/spotify/src/api/models"
	"github.com/science-engineering-art/spotify/src/api/pb"
	"github.com/science-engineering-art/spotify/src/api/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address    = "0.0.0.0:8080"
	conn, _    = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	songClient = pb.NewSongServiceClient(conn)
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

	// songBytes := bytes.NewReader(buffer)
	// m, err := tag.ReadFrom(songBytes)
	// if err != nil {
	// 	return err
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // build the Song model with all its metadata
	// newSong := models.Song{
	// 	Album:       m.Album(),
	// 	AlbumArtist: m.AlbumArtist(),
	// 	Artist:      m.Artist(),
	// 	Comment:     m.Comment(),
	// 	Composer:    m.Composer(),
	// 	FileType:    m.FileType(),
	// 	Format:      m.Format(),
	// 	Genre:       m.Genre(),
	// 	Id:          objId,
	// 	Lyrics:      m.Lyrics(),
	// 	Title:       m.Title(),
	// 	Year:        m.Year(),
	// }

	// &pb.CreateSongRequest{
	// 	RawSong: base64.RawStdEncoding.EncodeToString(buffer),
	// 	}
	stream, err := songClient.CreateSong(ctx)
	if err != nil {
		return err
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	step := 1024
	for i := 0; i < len(buffer); i += step {
		init, end := i, min(i+step, len(buffer))

		err = stream.Send(&pb.RawSong{
			Init:   int32(init),
			End:    int32(end),
			Buffer: buffer[init:end],
		})

		if err != nil {
			return err
		}
	}

	return c.Status(
		http.StatusCreated).JSON(
		responses.SongResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"success": true},
		},
	)
}

func GetSongById(c *fiber.Ctx) error {
	// get the song ID
	songId := c.Params("songId")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	song, err := songClient.GetSongById(ctx, &pb.SongId{
		Id: songId,
	})
	if err != nil {
		return nil
	}

	// // find the song with the ID requested
	// err := songCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&song)
	// if err != nil {
	// 	return c.Status(
	// 		http.StatusInternalServerError).JSON(
	// 		responses.SongResponse{
	// 			Status:  http.StatusInternalServerError,
	// 			Message: "error",
	// 			Data:    &fiber.Map{"data": err.Error()},
	// 		},
	// 	)
	// }

	// songBytes, err := base64.RawStdEncoding.DecodeString(song.RawSong)
	// if err != nil {
	// 	return c.Status(
	// 		http.StatusInternalServerError).JSON(
	// 		responses.SongResponse{
	// 			Status:  http.StatusInternalServerError,
	// 			Message: "error",
	// 			Data:    &fiber.Map{"data": err.Error()},
	// 		},
	// 	)
	// }

	// Get the unique identifier from the request context
	requestId, ok := c.Context().UserValue("requestId").(uuid.UUID)
	if !ok {
		// Handle error if unique identifier cannot be obtained
		return fiber.NewError(fiber.StatusInternalServerError, "Unique identifier could not be obtained")
	}

	fileName := fmt.Sprintf("./tmp_%s.mp3", requestId.String())

	os.WriteFile(fileName, song.RawSong.Buffer, 0600)
	defer os.Remove(fileName)

	return c.SendFile(fileName, true)
}

// func EditSong(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	songId := c.Params("songId")
// 	var song models.Song
// 	defer cancel()

// 	objId, _ := primitive.ObjectIDFromHex(songId)

// 	//validate the request body
// 	if err := c.BodyParser(&song); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(
// 			responses.SongResponse{
// 				Status:  http.StatusBadRequest,
// 				Message: "error",
// 				Data:    &fiber.Map{"data": err.Error()},
// 			},
// 		)
// 	}

// 	// construct the updated song
// 	update := bson.M{
// 		"album":       song.Album,
// 		"albumartist": song.AlbumArtist,
// 		"artist":      song.Artist,
// 		"comment":     song.Comment,
// 		"composer":    song.Composer,
// 		"filetype":    song.FileType,
// 		"format":      song.Format,
// 		"genre":       song.Genre,
// 		"_id":         song.Id,
// 		"lyrics":      song.Lyrics,
// 		"rawSong":     song.RawSong,
// 		"title":       song.Title,
// 		"year":        song.Year,
// 	}

// 	// updated in the DB
// 	result, err := songCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(
// 			responses.SongResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "error",
// 				Data:    &fiber.Map{"data": err.Error()},
// 			},
// 		)
// 	}

// 	//get updated user details
// 	var updatedSong models.Song
// 	if result.MatchedCount == 1 {
// 		err := songCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedSong)

// 		if err != nil {
// 			return c.Status(http.StatusInternalServerError).JSON(
// 				responses.SongResponse{
// 					Status:  http.StatusInternalServerError,
// 					Message: "error",
// 					Data:    &fiber.Map{"data": err.Error()},
// 				},
// 			)
// 		}
// 	}

// 	return c.Status(http.StatusOK).JSON(
// 		responses.SongResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    &fiber.Map{"data": updatedSong},
// 		},
// 	)
// }

// func DeleteSong(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	songId := c.Params("songId")
// 	defer cancel()

// 	objId, _ := primitive.ObjectIDFromHex(songId)

// 	result, err := songCollection.DeleteOne(ctx, bson.M{"_id": objId})
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(
// 			responses.SongResponse{
// 				Status:  http.StatusInternalServerError,
// 				Message: "error",
// 				Data:    &fiber.Map{"data": err.Error()},
// 			},
// 		)
// 	}

// 	// check that as least there are one file with the specified ID
// 	if result.DeletedCount < 1 {
// 		return c.Status(http.StatusNotFound).JSON(
// 			responses.SongResponse{
// 				Status:  http.StatusNotFound,
// 				Message: "error",
// 				Data:    &fiber.Map{"data": "User with specified ID not found!"},
// 			},
// 		)
// 	}

// 	return c.Status(http.StatusOK).JSON(
// 		responses.SongResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    &fiber.Map{"data": "User successfully deleted!"},
// 		},
// 	)
// }

func FilterSongs(c *fiber.Ctx) error {

	query := new(pb.SongMetadata)

	if err := c.BodyParser(query); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := songClient.FilterSongs(ctx, query)
	if err != nil {
		return err
	}

	var songs []models.SongDTO

	for {
		song, err := stream.Recv()
		if err != nil && err != io.EOF {
			return err
		}
		if song == nil {
			break
		}

		objID, err := primitive.ObjectIDFromHex(song.Id.Id)
		if err != nil {
			return err
		}

		songs = append(songs, models.SongDTO{
			Artist: *song.Metadata.Artist,
			Id:     objID,
			Title:  *song.Metadata.Title,
			Year:   int(*song.Metadata.Year),
		})
	}

	// // get all docs of the DB
	// results, err := songCollection.Find(ctx, bson.M{})
	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(
	// 		responses.SongResponse{
	// 			Status:  http.StatusInternalServerError,
	// 			Message: "error",
	// 			Data:    &fiber.Map{"data": err.Error()},
	// 		},
	// 	)
	// }
	// defer results.Close(ctx)

	// // keep with the songs objects
	// for results.Next(ctx) {
	// 	var singleSong models.SongDTO
	// 	if err = results.Decode(&singleSong); err != nil {
	// 		return c.Status(http.StatusInternalServerError).JSON(
	// 			responses.SongResponse{
	// 				Status:  http.StatusInternalServerError,
	// 				Message: "error",
	// 				Data:    &fiber.Map{"data": err.Error()},
	// 			},
	// 		)
	// 	}
	// 	songs = append(songs, singleSong)
	// }

	return c.Status(http.StatusOK).JSON(
		responses.SongResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"songs": songs},
		},
	)
}

// func GreetPeer(c *fiber.Ctx) error {
// 	const (
// 		defaultName = "world!"
// 	)

// 	var (
// 		addr = flag.String("addr", "localhost:50051", "the address to connect to")
// 		name = flag.String("name", defaultName, "Name to greet")
// 	)
// 	flag.Parse()
// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	d := pb.NewGreeterClient(conn)

// 	// Contact the server and print out its response.
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := d.SayHello(ctx, &pb.HelloRequest{Name: *name})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}

// 	return c.Status(http.StatusOK).JSON(
// 		responses.SongResponse{
// 			Status:  http.StatusOK,
// 			Message: "success",
// 			Data:    &fiber.Map{"data": r.GetMessage()},
// 		},
// 	)
// }
