package handler

import "net/http"

type StaticURLProvider interface {
	GetStaticBaseURL(r *http.Request) string
}
