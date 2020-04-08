package controllers

import (
	"net/http"
	"login/services"
	"encoding/json"
	"login/models"
	"log"
)

//Controller ...
type Controller struct {
	Repository services.Service
}

//Registration Controller
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Calling Registration with register modal..
	resp, _ := services.Service.Register(services.Service{}, register)

	respons, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respons)
	return
}

// Login Controller
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var respnse = models.RegisterResponse{}
	message := make(map[string]interface{})

	var loginData models.Login

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		log.Fatal(err)
		return
	}
	respnce, err := services.Service.LoginService(services.Service{}, loginData)

	if err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	//Creating token for User and storing all those details in the form of cookies.....
	cookies, tokenString, err := services.Service.CreateJWTTokenService(services.Service{}, loginData.Email)

	// Finally, we set the client cookie for "access_token" as the JWT we just generated
	// we also set an expiry time which is the same as the access_token itself
	http.SetCookie(w, cookies)

	message["message"] = "success"
	message["accessToken"] = tokenString

	respnse.Body = message
	respnse.StatusCode = 200

	resp, err := json.Marshal(respnse)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	var loginUserData models.Login
	if err := json.NewDecoder(r.Body).Decode(&loginUserData); err != nil {
		log.Fatal(err)
		return
	}

	respnce, err := services.Service.LogoutService(services.Service{}, w, loginUserData)

	if (err != nil) {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(respnce)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp)
	return
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	var updateData models.Register

	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		log.Fatal(err)
		return
	}
	respnce, err := services.Service.UpdateService(services.Service{}, updateData)

	if err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}
	resp, err := json.Marshal(respnce)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return
}
