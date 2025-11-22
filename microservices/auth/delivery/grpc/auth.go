package grpc

import (
	"context"
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
