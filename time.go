package kumex

import (
	"net/http"
)

// ServerTime returns the API server time.
func (as *ApiService) ServerTime() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/timestamp", nil)
	return as.Call(req)
}
