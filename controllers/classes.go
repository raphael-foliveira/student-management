package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func ListClasses(c *gin.Context) {
	class := models.Class{}
	classes := class.All()
	c.JSON(http.StatusOK, classes)
}

func RetrieveClass(c *gin.Context) {
	class := models.Class{}
	class.Find(c.Param("id"))
	if class.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	c.JSON(http.StatusOK, class)
}

func CreateClass(c *gin.Context) {
	var newClass models.Class
	if err := c.BindJSON(&newClass); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newClass.Create()
	c.JSON(http.StatusCreated, newClass)
}

func DeleteClass(c *gin.Context) {
	class := models.Class{}
	class.Delete(c.Param("id"))
	if class.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
