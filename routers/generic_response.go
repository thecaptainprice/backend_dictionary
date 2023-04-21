package routers

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/status"
)

type GenericResponse struct {
	Success  bool        `json:"success"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
}

func (r *GenericResponse) HandleError(err error) bool {
	if err == nil {
		// no error happened
		return false
	}
	// r.Error = err.Error()
	e, _ := status.FromError(err)

	r.Messages = append(r.Messages, e.Message())
	// r.Messages = append(r.Messages, err.Error())
	r.Success = false
	// error happened
	return true
}

// ────────────────────────────────────────────────────────────────────────────────
func (r *GenericResponse) ResponseJson(w http.ResponseWriter, status int) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		return
	}
	w.WriteHeader(status)
}
