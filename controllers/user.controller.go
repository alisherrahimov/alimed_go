package controllers

import (
	"alimed_go/configs"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Users struct {
	Id           int    `json:"id"`
	F_Name       string `json:"f_name"`
	L_Name       string `json:"l_name"`
	Phone_Number string `json:"phone_number"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
	Created_At   string `json:"created_at"`
	Updated_At   string `json:"updated_at"`
}

type UserInput struct {
	F_Name       string `json:"f_name"`
	L_Name       string `json:"l_name"`
	Phone_Number string `json:"phone_number"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
}

func Get(ctx *gin.Context, db *sql.DB) {

	var users []Users

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		configs.Response(ctx, configs.Res{
			Data:    err.Error(),
			Success: false,
		}, 500)
		return
	}

	defer rows.Close()

	var user Users
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.F_Name, &user.L_Name, &user.Phone_Number,
			&user.Gender, &user.Age, &user.Created_At, &user.Updated_At)

		if err != nil {
			configs.Response(ctx, configs.Res{
				Data:    err.Error(),
				Success: false,
			}, 500)
			return
		}
		users = append(users, user)
	}

	configs.Response(ctx, configs.Res{
		Data: users,
	}, 200)

}

func Post(ctx *gin.Context, db *sql.DB) {

	var user UserInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		configs.Response(ctx, configs.Res{
			Data: err.Error(),
		}, 400)
		return
	}

	result, err := db.Exec("INSERT INTO users (f_name, l_name, phone_number, gender,age) VALUES (?,?,?,?,?)", user.F_Name, user.L_Name, user.Phone_Number, user.Gender, user.Age)

	if err != nil {
		configs.Response(ctx, configs.Res{
			Data:    err.Error(),
			Success: false,
		}, 500)
		return
	}

	id, _ := result.LastInsertId()

	configs.Response(ctx, configs.Res{
		Data: id,
	}, 200)

}

func Delete(ctx *gin.Context, db *sql.DB) {

	id := ctx.Param("id")

	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		configs.Response(ctx, configs.Res{
			Data:    err.Error(),
			Success: false,
		}, 500)
		return
	}

	configs.Response(ctx, configs.Res{
		Data: "Data Deleted",
	}, 200)
}

func Put(ctx *gin.Context, db *sql.DB) {

	id := ctx.Param("id")

	var user UserInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		configs.Response(ctx, configs.Res{
			Data: err.Error(),
		}, 400)
		return
	}

	result, err := db.Exec("UPDATE users SET f_name = ?, l_name = ?, phone_number = ?, gender = ?, age = ? WHERE id = ?", user.F_Name, user.L_Name, user.Phone_Number, user.Gender, user.Age, id)

	if err != nil {
		configs.Response(ctx, configs.Res{
			Data:    err.Error(),
			Success: false,
		}, 500)
		return
	}

	configs.Response(ctx, configs.Res{
		Data: result,
	}, 200)

}

func GetById(ctx *gin.Context, db *sql.DB) {

	id := ctx.Param("id")

	var user Users

	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.F_Name, &user.L_Name, &user.Phone_Number, &user.Gender, &user.Age, &user.Created_At, &user.Updated_At)

	if err != nil {
		configs.Response(ctx, configs.Res{
			Data:    err.Error(),
			Success: false,
		}, 500)
		return
	}

	configs.Response(ctx, configs.Res{
		Data: user,
	}, 200)

}
