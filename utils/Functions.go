package utils

import (
	"net/http"
)

func StatusText(code int) string {
	return http.StatusText(code)
}
