package services

import (
	"login/models"
	"login/validaton"
	"login/repository"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"login/common"
	"log"
)

type Service struct{}

func (r Service) Register(registrationData models.Register) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})

	//If any register values missed it through errors.
	if resp, err := validaton.RegistrationValidation(registrationData); err != nil {
		//log.Fatal(err)
		return resp, err
	}
	/*	Connect database and store registration data
	*/
	resp, err := repository.RegistrationRepository(registrationData)

	return resp, err
}

func (r Service) LoginService(loginData models.Login) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	if resp, err := validaton.LoginValidation(loginData); err != nil {
		return resp, err
	}
	//Connect database and store login data
	userResp, err := repository.LoginRepository(loginData)

	if err != nil {
		resp["statusCode"] = 500
		resp["message"] = "User name or password wrong."
		return resp, err
	}
	resp["statusCode"] = 500
	resp["User Details"] = userResp
	resp["message"] = "User login successfully"
	return resp, err
}
func (r Service) CreateJWTTokenService(email string) (*http.Cookie, string, error) {
	//Expiry time is set 1 hrs after 1 hrs delete token from cookies
	// Expiry time is expressed as unix milliseconds
	expirationTime := time.Now().Add(60 * time.Minute)
	c := &http.Cookie{}

	// JWT claims contains username and expiry time
	claims := &models.Claims{
		Username: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Declaring token with SHA256 algorithm and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Creating the JWT string
	tokenString, err := token.SignedString(common.JWT_KEY)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		log.Fatal(err)
		return c, tokenString, err
	}

	c = &http.Cookie{
		Name:    "access_token",
		Value:   tokenString,
		Expires: expirationTime,
	}

	return c, tokenString, nil
}

func (r Service) LogoutService(w http.ResponseWriter, loginUserData models.Login) (map[string]interface{}, error) {

	var resp = make(map[string]interface{})

	if resp, err := validaton.LoginValidation(loginUserData); err != nil {
		return resp, err
	}

	//TODO
	//First check user exist in data base and session exist in the cookies.


	/*
	Logout service delete access tokes from browser cookies. And
	Update expiration time to -1.
	*/

	c := http.Cookie{
		Name:   "access_token",
		MaxAge: -1}
	http.SetCookie(w, &c)

	resp["statusCode"] = 200
	resp["message"] = "You have been successfully logged out."

	return resp, nil
}

func (r Service) UpdateService(updateData models.Register) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	if resp, err := validaton.RegistrationValidation(updateData); err != nil {
		return resp, err
	}
	//Connect database and store login data
	userResp, err := repository.UpdateRepository(updateData)

	if err != nil {
		resp["statusCode"] = 500
		resp["message"] = "User data not updated. Please insert proper data."
		return resp, err
	}
	resp["statusCode"] = 200
	resp["message"] = userResp
	return resp, err
}