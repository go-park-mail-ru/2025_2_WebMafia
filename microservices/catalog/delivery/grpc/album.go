package grpc

import (
	"context"
	"spotify/internal/middleware"
	pb "spotify/proto/catalog"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetAlbumByID(ctx context.Context, req *pb.GetAlbumByIDRequest) (*pb.Album, error) {
	const op = "grpc.GetAlbumByID"
	log := middleware.LoggerFromContext(ctx)

	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Warnf("[%s]: invalid album ID format: %v", op, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid album ID format: %v", err)
	}

	album, err := h.service.GetAlbumByID(ctx, id)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.NotFound, "album not found: %v", err)
	}

	return AlbumToProto(album), nil
}

func (h *Handler) GetAllAlbums(ctx context.Context, req *pb.GetAllAlbumsRequest) (*pb.GetAllAlbumsResponse, error) {
	const op = "grpc.GetAllAlbums"
	log := middleware.LoggerFromContext(ctx)

	albums, err := h.service.GetAllAlbums(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get all albums: %v", err)
	}
	return &pb.GetAllAlbumsResponse{Albums: AlbumsToProto(albums)}, nil
}

func (h *Handler) GetAlbumsByArtistID(ctx context.Context, req *pb.GetAlbumsByArtistIDRequest) (*pb.GetAllAlbumsResponse, error) {
	const op = "grpc.GetAlbumsByArtistID"
	log := middleware.LoggerFromContext(ctx)

	id, err := uuid.Parse(req.GetArtistId())
	if err != nil {
		log.Warnf("[%s]: invalid artist ID format: %v", op, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid artist ID format: %v", err)
	}
	albums, err := h.service.GetAlbumsByArtistID(ctx, id, req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get albums by artist: %v", err)
	}
	return &pb.GetAllAlbumsResponse{Albums: AlbumsToProto(albums)}, nil
}

func (h *Handler) GetAlbumsByIDs(ctx context.Context, req *pb.GetAlbumsByIDsRequest) (*pb.GetAllAlbumsResponse, error) {
	const op = "grpc.GetAlbumsByIDs"
	log := middleware.LoggerFromContext(ctx)

	if len(req.GetIds()) == 0 {
		return &pb.GetAllAlbumsResponse{}, nil
	}

	ids := make([]uuid.UUID, 0, len(req.GetIds()))
	for _, idStr := range req.GetIds() {
		id, err := uuid.Parse(idStr)
		if err != nil {
			log.Warnf("[%s]: invalid album ID format: %v", op, err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid album ID format: %v", err)
		}
		ids = append(ids, id)
	}

	albums, err := h.service.GetAlbumsByIDs(ctx, ids)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get albums by IDs: %v", err)
	}

	return &pb.GetAllAlbumsResponse{Albums: AlbumsToProto(albums)}, nil
}

func (h *Handler) SearchAlbums(ctx context.Context, req *pb.SearchAlbumsRequest) (*pb.SearchAlbumsResponse, error) {
	const op = "grpc.SearchAlbums"
	log := middleware.LoggerFromContext(ctx)

	results, err := h.service.SearchAlbums(ctx, req.GetQuery(), req.GetLimit())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to search albums: %v", err)
	}

	protoResults := make([]*pb.AlbumSearchResult, 0, len(results))
	for i := range results {
		protoResults = append(protoResults, &pb.AlbumSearchResult{
			Album: AlbumToProto(&results[i].Album),
			Rank:  results[i].Rank,
		})
	}

	return &pb.SearchAlbumsResponse{Results: protoResults}, nil
}
