package rpc

import (
	"context"
	"strings"

	"github.com/science-engineering-art/spotify/src/peer/pb"
	"github.com/science-engineering-art/spotify/src/peer/services"
	"github.com/science-engineering-art/spotify/src/peer/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SongServer struct {
	pb.UnimplementedSongServiceServer
	songCollection *mongo.Collection
	songService    services.SongService
}

func NewGrpcSongServer(songCollection *mongo.Collection, songService services.SongService) (*SongServer, error) {

	songServer := &SongServer{
		songCollection: songCollection,
		songService:    songService,
	}

	return songServer, nil
}

func (songServer *SongServer) CreateSong(stream pb.SongService_CreateSongServer) error {

	buffer := []byte{}
	var init int32 = 0

	for {
		rawSong, err := stream.Recv()
		if rawSong == nil {
			break
		}

		if init == rawSong.Init {
			buffer = append(buffer, rawSong.Buffer...)
			init = rawSong.End
		} else {
			return err
		}

		if err != nil {
			return err
		}
	}

	err := songServer.songService.CreateSong(&buffer)
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			return status.Errorf(codes.AlreadyExists, err.Error())
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (songServer *SongServer) GetSongById(ctx context.Context, req *pb.SongId) (*pb.Song, error) {

	objID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	song, err := songServer.songService.GetSongById(&objID)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	songFileType := string(song.FileType)
	songFormat := string(song.Format)
	songYear := int32(song.Year)

	pbSong := &pb.Song{
		Id: &pb.SongId{
			Id: song.Id.Hex(),
		},
		RawSong: &pb.RawSong{
			Init:   0,
			End:    int32(1000024),
			Buffer: song.RawSong[:1000024],
		},
		Metadata: &pb.SongMetadata{
			Album:       &song.Album,
			AlbumArtist: &song.AlbumArtist,
			Artist:      &song.Artist,
			Comment:     &song.Comment,
			Composer:    &song.Composer,
			FileType:    &songFileType,
			Format:      &songFormat,
			Genre:       &song.Genre,
			Lyrics:      &song.Lyrics,
			Title:       &song.Title,
			Year:        &songYear,
		},
	}

	return pbSong, nil
}

func (songServer *SongServer) UpdateSong(ctx context.Context, req *pb.UpdatedSong) (*pb.Response, error) {

	objID, err := primitive.ObjectIDFromHex(req.Id.Id)
	if err != nil {
		return nil, err
	}

	updatedSong := utils.BuildQuery(req.Metadata)

	err = songServer.songService.UpdateSong(&objID, updatedSong)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &pb.Response{
		Success: true,
	}
	return resp, nil
}

func (songServer *SongServer) RemoveSongById(ctx context.Context, req *pb.SongId) (*pb.Response, error) {

	objID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	if err := songServer.songService.RemoveSongById(&objID); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.Response{
		Success: true,
	}

	return res, nil
}

func (songServer *SongServer) FilterSongs(req *pb.SongMetadata, stream pb.SongService_FilterSongsServer) error {

	query := utils.BuildQuery(req)

	songs, err := songServer.songService.FilterSongs(query)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, song := range songs {

		songFileType := string(song.FileType)
		songFormat := string(song.Format)
		songYear := int32(song.Year)

		pbSong := &pb.Song{
			Id: &pb.SongId{
				Id: song.Id.Hex(),
			},
			RawSong: &pb.RawSong{
				Init:   0,
				End:    int32(1024),
				Buffer: song.RawSong[:1024],
			},
			Metadata: &pb.SongMetadata{
				Album:       &song.Album,
				AlbumArtist: &song.AlbumArtist,
				Artist:      &song.Artist,
				Comment:     &song.Comment,
				Composer:    &song.Composer,
				FileType:    &songFileType,
				Format:      &songFormat,
				Genre:       &song.Genre,
				Lyrics:      &song.Lyrics,
				Title:       &song.Title,
				Year:        &songYear,
			},
		}

		stream.Send(pbSong)
	}

	return nil
}
