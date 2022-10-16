package controllers

import (
	"assesment3/webserver/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetWeathers(ctx *gin.Context) {
	r := rand.NewSource(time.Now().UnixNano())
	ctx.Header("content-type", "application/json")
	data := models.Status{
		Water: rand.New(r).Intn(100),
		Wind:  rand.New(r).Intn(100),
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": data,
	})
}
