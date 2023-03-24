package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/models"
)

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(400, err)
		return
	}
	teacher.Create()
	c.JSON(201, teacher)
}

func ListTeachers(c *gin.Context) {
	teacher := models.Teacher{}
	teachers := teacher.All()
	c.JSON(200, teachers)
}

func RetrieveTeacher(c *gin.Context) {
	teacher := models.Teacher{}
	teacher.Find(c.Param("id"))
	if teacher.ID == 0 {
		c.JSON(404, gin.H{"error": "Teacher not found"})
		return
	}
	c.JSON(200, teacher)
}

func DeleteTeacher(c *gin.Context) {
	teacher := models.Teacher{}
	teacher.Delete(c.Param("id"))
	if teacher.ID == 0 {
		c.JSON(404, gin.H{"error": "Teacher not found"})
		return
	}
	c.JSON(204, nil)
}
