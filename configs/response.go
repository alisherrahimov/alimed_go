package configs

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Res struct {
	Data    interface{} `json:"data"`
	Time    time.Time   `json:"time"`
	Success bool        `json:"success"`
	Error   interface{} `json:"error"`
}

func Response(ctx *gin.Context, body Res, status int) {
	body.Time = time.Now()
	body.Success = true
	body.Error = nil
	ctx.JSON(status, body)
}
