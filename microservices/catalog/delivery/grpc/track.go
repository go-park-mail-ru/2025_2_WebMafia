package grpc

import (
	"context"
	"spotify/internal/middleware"
	pb "spotify/proto/catalog"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetTrackByID(ctx context.Context, req *pb.GetTrackByIDRequest) (*pb.Track, error) {
	const op = "grpc.GetTrackByID"
	log := middleware.LoggerFromContext(ctx)

	id, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Warnf("[%s]: invalid track ID format: %v", op, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid track ID format: %v", err)
	}

	track, err := h.service.GetTrackByID(ctx, id)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.NotFound, "track not found: %v", err)
	}

	return TrackToProto(track), nil
}

func (h *Handler) GetAllTracks(ctx context.Context, req *pb.GetAllTracksRequest) (*pb.GetAllTracksResponse, error) {
	const op = "grpc.GetAllTracks"
	log := middleware.LoggerFromContext(ctx)

	tracks, err := h.service.GetAllTracks(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get all tracks: %v", err)
	}
	return &pb.GetAllTracksResponse{Tracks: TracksToProto(tracks)}, nil
}

func (h *Handler) GetTracksByArtistID(ctx context.Context, req *pb.GetTracksByArtistIDRequest) (*pb.GetAllTracksResponse, error) {
	const op = "grpc.GetTracksByArtistID"
	log := middleware.LoggerFromContext(ctx)

	id, err := uuid.Parse(req.GetArtistId())
	if err != nil {
		log.Warnf("[%s]: invalid artist ID format: %v", op, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid artist ID format: %v", err)
	}
	tracks, err := h.service.GetTracksByArtistID(ctx, id, req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get tracks by artist: %v", err)
	}
	return &pb.GetAllTracksResponse{Tracks: TracksToProto(tracks)}, nil
}

func (h *Handler) GetTracksByAlbumID(ctx context.Context, req *pb.GetTracksByAlbumIDRequest) (*pb.GetAllTracksResponse, error) {
	const op = "grpc.GetTracksByAlbumID"
	log := middleware.LoggerFromContext(ctx)
	id, err := uuid.Parse(req.GetAlbumId())
	if err != nil {
		log.Warnf("[%s]: invalid album ID format: %v", op, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid album ID format: %v", err)
	}
	tracks, err := h.service.GetTracksByAlbumID(ctx, id, req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get tracks by album: %v", err)
	}
	return &pb.GetAllTracksResponse{Tracks: TracksToProto(tracks)}, nil
}

func (h *Handler) GetTracksByGenreID(ctx context.Context, req *pb.GetTracksByGenreIDRequest) (*pb.GetAllTracksResponse, error) {
	const op = "grpc.GetTracksByGenreID"
	log := middleware.LoggerFromContext(ctx)
	id, err := uuid.Parse(req.GetGenreId())
	if err != nil {
		log.Warnf("[%s]: invalid genre ID format: %v", op, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid genre ID format: %v", err)
	}
	tracks, err := h.service.GetTracksByGenreID(ctx, id, req.GetLimit(), req.GetOffset())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get tracks by genre: %v", err)
	}
	return &pb.GetAllTracksResponse{Tracks: TracksToProto(tracks)}, nil
}

func (h *Handler) SearchTracks(ctx context.Context, req *pb.SearchTracksRequest) (*pb.SearchTracksResponse, error) {
	const op = "grpc.SearchTracks"
	log := middleware.LoggerFromContext(ctx)

	results, err := h.service.SearchTracks(ctx, req.GetQuery(), req.GetLimit())
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to search tracks: %v", err)
	}
	protoResults := make([]*pb.TrackSearchResult, 0, len(results))
	for i := range results {
		protoResults = append(protoResults, &pb.TrackSearchResult{
			Track: TrackToProto(&results[i].Track),
			Rank:  results[i].Rank,
		})
	}

	return &pb.SearchTracksResponse{Results: protoResults}, nil
}

func (h *Handler) RegisterPlay(ctx context.Context, req *pb.RegisterPlayRequest) (*pb.RegisterPlayResponse, error) {
	const op = "grpc.RegisterPlay"
	log := middleware.LoggerFromContext(ctx)

	id, err := uuid.Parse(req.GetTrackId())
	if err != nil {
		log.Warnf("[%s]: invalid track ID format: %v", op, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid track ID format: %v", err)
	}
	err = h.service.RegisterPlay(ctx, id)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to register play: %v", err)
	}
	return &pb.RegisterPlayResponse{}, nil
}

func (h *Handler) GetTracksByIDs(ctx context.Context, req *pb.GetTracksByIDsRequest) (*pb.GetTracksByIDsResponse, error) {
	const op = "grpc.GetTracksByIDs"
	log := middleware.LoggerFromContext(ctx)

	if len(req.GetIds()) == 0 {
		return &pb.GetTracksByIDsResponse{Tracks: []*pb.Track{}}, nil
	}

	uuids := make([]uuid.UUID, 0, len(req.Ids))
	for _, id := range req.GetIds() {
		parsed, err := uuid.Parse(id)
		if err != nil {
			log.Warnf("[%s]: invalid ID: %v", op, err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid track id: %s", id)
		}
		uuids = append(uuids, parsed)
	}

	tracks, err := h.service.GetTracksByIDs(ctx, uuids)
	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to get tracks: %v", err)
	}

	return &pb.GetTracksByIDsResponse{
		Tracks: TracksToProto(tracks),
	}, nil
}
