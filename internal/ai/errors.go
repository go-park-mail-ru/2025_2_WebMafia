package ai

import "errors"

var (
	ErrAIAuth        = errors.New("ai auth error")
	ErrAIRateLimit   = errors.New("ai rate limit")
	ErrAIUnavailable = errors.New("ai unavailable")
	ErrAINoChoices   = errors.New("ai no choices")
)
