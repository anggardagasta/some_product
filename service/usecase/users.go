package usecase

import (
	"encoding/base64"
	"errors"
	"github.com/anggardagasta/some_product/models"
	"github.com/anggardagasta/some_product/service"
	"github.com/anggardagasta/some_product/service/repository/constant"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type serviceUsersUsecase struct {
	serviceUsersRepo service.IServiceUsersRepository
}

func NewServiceUsersUsecase(serviceUserRepo service.IServiceUsersRepository) service.IServiceUsersUseCase {
	return serviceUsersUsecase{serviceUsersRepo: serviceUserRepo}
}

func (uc serviceUsersUsecase) GeneratingToken(userID int64, username string) (result string, err error) {
	claims := models.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    constant.ApplicationName,
			ExpiresAt: time.Now().Add(constant.LoginExpirationDuration).Unix(),
		},
		ID:       userID,
		Username: username,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (uc serviceUsersUsecase) RegisterUser(form models.FormRegister) (result models.AuthResult, err error) {
	checkUser, err := uc.serviceUsersRepo.GetUserByUsername(form.Username)
	if err != nil {
		return result, err
	}
	if checkUser.ID.Int64 > int64(0) {
		return result, errors.New("username is already exist")
	}

	form.Password = base64.StdEncoding.EncodeToString([]byte(form.Password))
	userID, err := uc.serviceUsersRepo.InsertUser(form)
	if err != nil {
		return result, err
	}

	token, err := uc.GeneratingToken(userID, form.Username)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (uc serviceUsersUsecase) Auth(form models.FormAuth) (result models.AuthResult, err error) {
	user, err := uc.serviceUsersRepo.GetUserByUsername(form.Username)
	if err != nil {
		return result, err
	}
	if user.ID.Int64 == 0 {
		return result, errors.New("user not found")
	}

	userPassword, err := base64.StdEncoding.DecodeString(user.Password.String)
	if err != nil {
		return result, err
	}
	if form.Password == string(userPassword) {
		token, err := uc.GeneratingToken(user.ID.Int64, user.Username.String)
		if err != nil {
			return result, err
		}

		result.Token = token
	} else {
		return result, errors.New("invalid password")
	}
	return result, nil
}

func (uc serviceUsersUsecase) GetProfile(id int64) (result models.GetProfileResult, err error) {
	user, err := uc.serviceUsersRepo.GetUserByID(id)
	if err != nil {
		return result, err
	}

	result.ID = user.ID.Int64
	result.UserName = user.Username.String
	result.FullName = user.FullName.String
	result.Picture = user.Picture.String

	return result, nil
}

func (uc serviceUsersUsecase) UpdateProfile(id int64, form models.FormUpdateProfile) (err error) {
	//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(form.Picture))
	//m, _, err := image.Decode(reader)
	//if err != nil {
	//	return err
	//}
	//_ := m.Bounds()
	//
	//pngFilename := strconv.Itoa(int(id)) + `.png`
	//f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	//if err != nil {
	//	return err
	//}
	//
	//err = png.Encode(f, m)
	//if err != nil {
	//	return err
	//}

	form.Picture = "1.jpg"
	err = uc.serviceUsersRepo.UpdateUser(id, form)
	if err != nil {
		return err
	}

	return nil
}
