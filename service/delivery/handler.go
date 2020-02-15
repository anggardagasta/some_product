package delivery

import (
	"encoding/json"
	"errors"
	"github.com/anggardagasta/mini_wallet/models"
	"github.com/anggardagasta/mini_wallet/service/repository/response"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func authenticateJWT(r *http.Request) (result models.UserJWT, err error) {
	authorizationHeader := r.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return result, errors.New("missing token")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNATURE")), nil
	})
	if err != nil {
		return result, err
	}

	_, ok := claims["id"]
	if !ok {
		return result, errors.New("invalid token")
	}

	_, ok = claims["username"]
	if !ok {
		return result, errors.New("invalid token")
	}

	for key, val := range claims {
		if key == "id" {
			result.ID = int64(val.(float64))
		}
		if key == "username" {
			result.Username = val.(string)
		}
	}

	return result, nil
}

func (hd handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var request models.FormRegister

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}

	data, err := hd.usersUseCase.RegisterUser(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}
	response.ResultWithData(w, data, response.MessageSucceed, http.StatusOK)
}

func (hd handler) Auth(w http.ResponseWriter, r *http.Request) {
	var request models.FormAuth

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}

	data, err := hd.usersUseCase.Auth(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}
	response.ResultWithData(w, data, response.MessageSucceed, http.StatusOK)
}

func (hd handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	user, err := authenticateJWT(r)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}

	data, err := hd.usersUseCase.GetProfile(user.ID)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}
	if data.ID == 0 {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, errors.New("user not found"))
		return
	}
	response.ResultWithData(w, data, response.MessageSucceed, http.StatusOK)
}

func (hd handler) UpdatePicture(w http.ResponseWriter, r *http.Request) {
	user, err := authenticateJWT(r)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}

	var request models.FormUpdateProfile

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}

	_, err = govalidator.ValidateStruct(request)
	if err != nil {
		response.ResultError(w, http.StatusBadRequest, response.MessageBadRequest, err)
		return
	}

	err = hd.usersUseCase.UpdateProfile(user.ID, request)
	if err != nil {
		response.ResultError(w, http.StatusInternalServerError, response.MessageInternalError, err)
		return
	}
	response.Result(w, response.MessageSucceed, http.StatusOK)
}
