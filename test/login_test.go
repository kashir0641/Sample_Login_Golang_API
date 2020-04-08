package test

import (
	"testing"
	"login/models"
	"login/services"
	"fmt"
)

//Registration and login testing multiple parameters.
// So we took array of struct as request data.
//

func Test_Register(t *testing.T) {
	registerData := []models.Register{
		{"alice", "P@ss0rd", "alice@gmail.com", "#101 51 street"},
		{"", "P@ss0rd", "mike@gmail.com", "#101 51 street"},
		{"mike", "", "mike@gmail.com", "#101 51 street"},
		{"bob", "P@ss0rd", "", "#101 51 street"},
		{"mike", "P@ss0rd", "mike@gmail.com", ""},
	}

	for _, item := range registerData {
		response, err := services.Service.Register(services.Service{}, item);
		if err != nil {
			t.Errorf(fmt.Sprintf("Error message : %v  and request data : %v", response["message"], item))
		} else {
			t.Logf(fmt.Sprintf("Success message : %v  and request data : %v", response["message"], item))
		}
	}
}

func Test_Login(t *testing.T) {
	loginData := []models.Login{
		{"mike", "P@ss0rd"},
		{"", "P@ss0rd"},
		{"", ""},
		{"", "P@ss0rd"},
	}

	for _, item := range loginData {
		response, err := services.Service.LoginService(services.Service{}, item);
		if err != nil {
			t.Errorf(fmt.Sprintf("Error message : %v  and request data : %v", response["message"], item))
		} else {
			t.Logf(fmt.Sprintf("Success message : %v  and request data : %v", response["message"], item))
		}
	}
}
