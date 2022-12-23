package controllers

import (
	"log"
	"net/http"

	"online-transportation/config"
	"online-transportation/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	defer db.Close()

	var response models.UsersResponse

	query := "SELECT id, email, phone_number, firstname, lastname, gender, province, city, address, registration_date, is_verified FROM user"

	id := r.URL.Query()["id"]

	if id != nil {
		query += " WHERE id = " + id[0]
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		SendResponse(w, response.Status, response)
		return
	}

	var user models.User
	var users []models.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email, &user.PhoneNumber, &user.Firstname, &user.Lastname, &user.Gender, &user.Province, &user.City, &user.Address, &user.RegistrationDate, &user.IsVerified); err != nil {
			log.Println(err.Error())
		} else {
			users = append(users, user)
		}
	}

	if len(users) != 0 {
		response.Status = 200
		response.Message = "Success get user data"
		response.Data = users
	} else {
		response.Status = 200
		response.Message = "User data not found"
	}

	SendResponse(w, response.Status, response)

}
