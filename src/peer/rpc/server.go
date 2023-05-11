package rpc

import (
	"context"
	"strings"

	"github.com/dhowden/tag"
	"github.com/science-engineering-art/spotify/src/peer/models"
	"github.com/science-engineering-art/spotify/src/peer/pb"
	"github.com/science-engineering-art/spotify/src/peer/services"

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
			buffer = append(buffer, rawSong.GetRawSong()...)
			init = rawSong.End
		} else {
			return err
		}

		if err != nil {
			return err
		}
	}

	err := songServer.songService.CreateSong(buffer)
	if err != nil {

		if strings.Contains(err.Error(), "title already exists") {
			return status.Errorf(codes.AlreadyExists, err.Error())
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (songServer *SongServer) UpdateSong(ctx context.Context, req *pb.UpdateSongRequest) (*pb.Response, error) {

	songId := req.GetId()

	objID, err := primitive.ObjectIDFromHex(songId)
	if err != nil {
		return nil, err
	}

	song := &models.UpdateSongRequest{
		Album:       req.GetAlbum(),
		AlbumArtist: req.GetAlbumArtist(),
		Artist:      req.GetArtist(),
		Comment:     req.GetComment(),
		Composer:    req.GetComposer(),
		FileType:    tag.FileType(req.GetFileType()),
		Format:      tag.Format(req.GetFormat()),
		Genre:       req.GetGenre(),
		Id:          objID,
		Lyrics:      req.GetLyrics(),
		Title:       req.GetTitle(),
		Year:        int(req.GetYear()),
	}

	err = songServer.songService.UpdateSong(song)

	if err != nil {
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

func (songServer *SongServer) GetSongById(ctx context.Context, req *pb.SongIdRequest) (*pb.Song, error) {

	songId := req.GetId()

	song, err := songServer.songService.GetSongById(songId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.Song{
		Album:       song.Album,
		AlbumArtist: song.AlbumArtist,
		Artist:      song.Artist,
		Comment:     song.Comment,
		Composer:    song.Composer,
		FileType:    string(song.FileType),
		Format:      string(song.Format),
		Genre:       song.Genre,
		Id:          songId,
		Lyrics:      song.Lyrics,
		Title:       song.Title,
		Year:        int32(song.Year),
	}
	return res, nil
}

func (songServer *SongServer) GetSongs(req *pb.Request, stream pb.SongService_GetSongsServer) error {

	songs, err := songServer.songService.GetSongs()
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, song := range songs {
		stream.Send(&pb.Song{
			Album:       song.Album,
			AlbumArtist: song.AlbumArtist,
			Artist:      song.Artist,
			Comment:     song.Comment,
			Composer:    song.Composer,
			FileType:    string(song.FileType),
			Format:      string(song.Format),
			Genre:       song.Genre,
			Id:          song.Id.String(),
			Lyrics:      song.Lyrics,
			Title:       song.Title,
			Year:        int32(song.Year),
		})
	}

	return nil
}

func (songServer *SongServer) RemoveSongById(ctx context.Context, req *pb.SongIdRequest) (*pb.Response, error) {
	songId := req.GetId()

	if err := songServer.songService.DeleteSong(songId); err != nil {
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
