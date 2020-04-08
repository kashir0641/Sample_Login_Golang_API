package repository

import (
	"login/models"
	"gopkg.in/mgo.v2"
	"login/common"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

/*
Inside Repository perform database operations
All services database logic's are written in repository files.
*/

func RegistrationRepository(registrationData models.Register) (map[string]interface{}, error) {
	var user models.Register
	var resp = make(map[string]interface{})
	session, err := mgo.Dial(common.MongoUrl)
	common.Check(err)
	defer session.Close()
	collection := session.DB(common.Database).C(common.UserCollecion)

	if err := collection.Find(bson.M{"email": registrationData.Email}).One(&user); err != nil {
		if err := collection.Insert(registrationData); err != nil {
			panic(err)
		}
		resp["statusCode"] = 201
		resp["message"] = "New user created"

		return resp, err
	} else {
		resp["statusCode"] = 200
		resp["message"] = "User already exist."
		return resp, nil

	}
	return resp, nil
}

func LoginRepository(loginData models.Login) (models.Register, error) {
	var userData models.Register
	session, err := mgo.Dial(common.MongoUrl)
	common.Check(err)
	defer session.Close()
	collection := session.DB(common.Database).C(common.UserCollecion)
	if err := collection.Find(bson.M{"email": loginData.Email, "password": loginData.Password}).One(&userData); err != nil {
		return userData, err
	}
	return userData, nil
}


func UpdateRepository(updateData models.Register) (string, error) {
	var responseMessage string
	session, err := mgo.Dial(common.MongoUrl)
	common.Check(err)
	defer session.Close()

	filter := bson.M{"email": updateData.Email}
	change := bson.M{
		"$set": bson.M{
			"username": updateData.UserName,
			"password": updateData.Password,
			"email":    updateData.Email,
			"address":  updateData.Address,
		},
	}

	collection := session.DB(common.Database).C(common.UserCollecion)

	err = collection.Update(filter, change)
	if err != nil {
		fmt.Println(err)
		return responseMessage, err
	}
	message := "User name : " + updateData.UserName + " is updated successfully with email address ( " + updateData.Email + " )"
	responseMessage = message
	return responseMessage, nil
}