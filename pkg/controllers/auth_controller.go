package controllers

import (
	"alimed_go/pkg/configs"
	"alimed_go/pkg/helper"
	"alimed_go/pkg/models"
	"alimed_go/pkg/services"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, db *sql.DB) {
	// code here
	var loginReq models.LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		configs.Response(ctx, configs.Res{
			Data:    nil,
			Error:   err,
			Success: false,
		}, http.StatusBadRequest)
	}

	isSend := services.Login(loginReq)
	if isSend {
		configs.Response(ctx, configs.Res{
			Data:    "OTP has been sent to your phone number",
			Error:   nil,
			Success: true,
		}, http.StatusOK)
	} else {
		configs.Response(ctx, configs.Res{
			Data:    nil,
			Error:   "Failed to send OTP",
			Success: false,
		}, http.StatusBadRequest)
	}
}

func LoginByEmail(ctx *gin.Context, db *sql.DB) {

}

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
func VerifyOTP(ctx *gin.Context, db *sql.DB) {
	// code here

	var verifyReq models.VerifyOTPRequest

	if err := ctx.ShouldBindJSON(&verifyReq); err != nil {
		configs.Response(ctx, configs.Res{
			Data:    nil,
			Error:   err,
			Success: false,
		}, http.StatusBadRequest)
	}

	isValid, err := services.VerifyOTP(db, verifyReq)
	if err != nil {
		println("error")
		configs.Response(ctx, configs.Res{
			Data:    nil,
			Error:   err,
			Success: false,
		}, http.StatusBadRequest)
	}

	if isValid {

		configs.Response(ctx, configs.Res{
			Data:    "OTP is valid",
			Error:   nil,
			Success: true,
		}, http.StatusOK)
	} else {
		configs.Response(ctx, configs.Res{
			Data:    nil,
			Error:   "OTP is invalid",
			Success: false,
		}, http.StatusBadRequest)
	}
}

func AskSomeData(ctx *gin.Context, db *sql.DB) {

	var askReq models.AskSomeDataRequest

	if err := ctx.ShouldBindJSON(&askReq); err != nil {
		configs.Response(ctx, configs.Res{
			Data:    nil,
			Error:   err,
			Success: false,
		}, http.StatusBadRequest)
	}

	user, err := services.UpdateSomeDataByPhoneNumber(db, askReq)
	if err != nil {
		configs.Response(ctx, configs.Res{
			Data:    nil,
			Error:   err,
			Success: false,
		}, http.StatusBadRequest)
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		configs.Response(ctx, configs.Res{
			Data:    nil,
			Error:   err,
			Success: false,
		}, http.StatusBadRequest)
	}

	configs.Response(ctx, configs.Res{
		Data:    token,
		Error:   nil,
		Success: true,
	}, http.StatusOK)
}

func DeleteSomeData(ctx *gin.Context, db *sql.DB) {
	configs.Response(ctx, configs.Res{
		Data:    "Data has been deleted",
		Error:   nil,
		Success: true,
	}, http.StatusOK)
}
