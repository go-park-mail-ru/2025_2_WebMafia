package response

var (
	ErrBadRequest     = ErrorResponse{Error: "bad request"}
	ErrUnauthorized   = ErrorResponse{Error: "unauthorized"}
	ErrNotFound       = ErrorResponse{Error: "not found"}
	ErrConflict       = ErrorResponse{Error: "resource conflict"}
	ErrInternalServer = ErrorResponse{Error: "internal server error"}
	ErrForbidden      = ErrorResponse{Error: "forbidden"}
)
