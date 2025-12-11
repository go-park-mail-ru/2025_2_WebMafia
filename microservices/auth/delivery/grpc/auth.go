package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"spotify/internal/middleware"
	pb "spotify/proto/auth"
)

func (h *Handler) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	const op = "grpc.ValidateToken"
	log := middleware.LoggerFromContext(ctx)

	claims, err := h.jwtManager.Validate(req.GetToken())
	if err != nil {
		log.Warnf("[%s]: token validation failed: %v", op, err)
		return &pb.ValidateTokenResponse{IsValid: false}, nil
	}

	return &pb.ValidateTokenResponse{
		IsValid:   true,
		UserId:    claims.UserID,
		SessionId: claims.SessionID,
	}, nil
}

func (h *Handler) CheckCSRF(ctx context.Context, req *pb.CheckCSRFRequest) (*pb.CheckCSRFResponse, error) {
	const op = "grpc.CheckCSRF"
	log := middleware.LoggerFromContext(ctx)

	isValid, err := h.csrfManager.Check(req.GetUserId(), req.GetSessionId(), req.GetCsrfToken())
	if err != nil {
		log.Warnf("[%s]: csrf check error: %v", op, err)
		return &pb.CheckCSRFResponse{IsValid: false}, nil
	}

	return &pb.CheckCSRFResponse{IsValid: isValid}, nil
}

func (h *Handler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	const op = "grpc.GetUsers"
	log := middleware.LoggerFromContext(ctx)

	if len(req.UserIds) == 0 {
		return &pb.GetUsersResponse{Users: []*pb.UserInfo{}}, nil
	}

	profiles, err := h.userService.GetUsersBatch(ctx, req.UserIds)

	if err != nil {
		log.Errorf("[%s]: service error: %v", op, err)
		return nil, status.Errorf(codes.Internal, "failed to fetch users: %v", err)
	}
	out := make([]*pb.UserInfo, 0, len(profiles))
	for _, p := range profiles {
		out = append(out, &pb.UserInfo{
			UserId:    p.ID,
			Login:     p.Login,
			AvatarUrl: p.AvatarURL,
		})
	}
	return &pb.GetUsersResponse{Users: out}, nil
}
