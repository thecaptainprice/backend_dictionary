package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thecaptainprice/dictionary-app/backend/request"
	"google.golang.org/grpc/status"
)

type Handler func(r *request.GenericRequest) (interface{}, error)

func DefaultHttpHandler(handler Handler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		var res GenericResponse
		genReq, err := request.NewGenericHTTPRequestFromHTTPRequest(req)
		if err != nil {
			fmt.Println(err)
			respondError(writer, code2Http[status.Code(err)], err.Error())
			return
		}

		res.Data, err = handler(genReq)
		if err != nil {
			fmt.Println(err)
			res.HandleError(err)
			res.Success = false
			res.ResponseJson(writer, code2Http[status.Code(err)])
			return
		}
		res.Success = true
		res.ResponseJson(writer, http.StatusOK)
	}
}

// respondError- Returns an error as a response
func respondError(w http.ResponseWriter, status int, msg string) {
	errBody := GenericResponse{
		Messages: []string{msg},
	}
	respondJSON(w, status, errBody)
}

/*
	respondJSON- Returns the response in JSON format to the response

writer | Adds content type and charset to the header map | Sends an
HTTP response header with the provided status code.
*/
func respondJSON(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(status)
}
