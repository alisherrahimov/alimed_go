package services

import (
	"alimed_go/pkg/configs"
	"alimed_go/pkg/helper"
	"alimed_go/pkg/models"
	"database/sql"
)

func Login(req models.LoginRequest) bool {
	code := helper.GenerateOTP()
	isSend := helper.SendSMS(req.PhoneNumber, code)
	if isSend {
		println("send")
		configs.SetItem("phone_number", code)
		return true
	} else {
		return false
	}
}

func VerifyOTP(db *sql.DB, req models.VerifyOTPRequest) (bool, error) {
	code, found := configs.GetItem("phone_number")
	if !found {
		println("not found")
		return false, nil
	}
	println("found", code.(int), req.OTP)
	if code.(int) == req.OTP {
		rows := db.QueryRow("SELECT phone_number FROM users WHERE phone_number = $1", req.PhoneNumber)
		err := rows.Scan(&req.PhoneNumber)
		if err != nil {

			res, err := db.Exec("INSERT INTO users (phone_number) VALUES ($1)", req.PhoneNumber)
			if err != nil {
				return false, err
			}

			_, err = res.RowsAffected()
			if err != nil {
				return false, err
			}
		}

		return true, nil
	} else {
		return false, nil
	}
}

func GetUserByPhoneNumber(db *sql.DB, phoneNumber string) (models.User, error) {
	var user models.User
	rows := db.QueryRow("SELECT * FROM users WHERE phone_number = $1", phoneNumber)
	err := rows.Scan(&user.Id, &user.FName, &user.LName, &user.PhoneNumber, &user.Gender, &user.Age, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, nil
	}
	return user, nil
}

func UpdateSomeDataByPhoneNumber(db *sql.DB, req models.AskSomeDataRequest) (models.User, error) {
	var user models.User
	rows := db.QueryRow("UPDATE users SET f_name = $1, l_name = $2, gender = $3, age = $4 WHERE phone_number = $5", req.FName, req.LName, req.Gender, req.Age, req.PhoneNumber)
	rows.Scan(&user.Id, &user.FName, &user.LName, &user.PhoneNumber, &user.Gender, &user.Age, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return user, nil
}
