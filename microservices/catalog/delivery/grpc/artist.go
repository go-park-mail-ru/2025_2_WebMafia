package grpc

import (
	"context"
	"spotify/internal/middleware"

	pb "spotify/proto/catalog"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetArtistByID(ctx context.Context, req *pb.GetArtistByIDRequest) (*pb.Artist, error) {
	const op = "grpc.GetArtistByID"
	log := middleware.LoggerFromContext(ctx)

	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Warnf("[%s]: invalid artist ID format: %v", op, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid artist ID format: %v", err)
	}

	artist, err := h.service.GetArtistByID(ctx, id)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.NotFound, "artist not found: %v", err)
	}

	return ArtistToProto(artist), nil
}

func (h *Handler) GetAllArtists(ctx context.Context, req *pb.GetAllArtistsRequest) (*pb.GetAllArtistsResponse, error) {
	const op = "grpc.GetAllArtists"
	log := middleware.LoggerFromContext(ctx)

	artists, err := h.service.GetAllArtists(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get all artists: %v", err)
	}
	return &pb.GetAllArtistsResponse{Artists: ArtistsToProto(artists)}, nil
}

func (h *Handler) GetArtistsByIDs(ctx context.Context, req *pb.GetArtistsByIDsRequest) (*pb.GetArtistsByIDsResponse, error) {
	const op = "grpc.GetArtistsByIDs"
	log := middleware.LoggerFromContext(ctx)

	if len(req.GetIds()) == 0 {
		return &pb.GetArtistsByIDsResponse{}, nil
	}

	ids := make([]uuid.UUID, 0, len(req.GetIds()))
	for _, idStr := range req.GetIds() {
		id, err := uuid.Parse(idStr)
		if err != nil {
			log.Warnf("[%s]: invalid artist ID format: %v", op, err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid artist ID format: %v", err)
		}
		ids = append(ids, id)
	}

	artists, err := h.service.GetArtistsByIDs(ctx, ids)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get artists by IDs: %v", err)
	}

	return &pb.GetArtistsByIDsResponse{Artists: ArtistsToProto(artists)}, nil
}

func (h *Handler) SearchArtists(ctx context.Context, req *pb.SearchArtistsRequest) (*pb.SearchArtistsResponse, error) {
	const op = "grpc.SearchArtists"
	log := middleware.LoggerFromContext(ctx)

	results, err := h.service.SearchArtists(ctx, req.GetQuery(), req.GetLimit())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to search artists: %v", err)
	}

	protoResults := make([]*pb.ArtistSearchResult, 0, len(results))
	for i := range results {
		protoResults = append(protoResults, &pb.ArtistSearchResult{
			Artist: ArtistToProto(&results[i].Artist),
			Rank:   results[i].Rank,
		})
	}

	return &pb.SearchArtistsResponse{Results: protoResults}, nil
}
