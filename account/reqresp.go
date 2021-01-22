package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json: "password"`
	}
	CreateUserRequest struct {
		ok string `json: "ok"`
	}

	GetUserResponse struct {
		Email string `json: "email"`
	}
	GetUserRequest struct {
		Id string `json: "id"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decoderUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		Id: vars["id"],
	}
	return req, nil
}
