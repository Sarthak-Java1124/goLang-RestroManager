package controllers

import (
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/gin-gonic/gin"
)

var menuCollection = database.OpenCollection(*database.DBinstance(), "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
