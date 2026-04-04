package controllers

import (
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/gin-gonic/gin"
)

var tableCollection = database.OpenCollection(*database.DBinstance(), "table")

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
