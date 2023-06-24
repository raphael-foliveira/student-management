package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/models"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
)

type TeachersController struct {
	repo repository.Repository[models.Teacher]
}

func NewTeachersController(repo repository.Repository[models.Teacher]) *TeachersController {
	return &TeachersController{repo: repo}
}

func (tc *TeachersController) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := tc.repo.Create(&teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, teacher)
}

func (tc *TeachersController) ListTeachers(c *gin.Context) {
	teachers, err := tc.repo.All()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, teachers)
}

func (tc *TeachersController) RetrieveTeacher(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	teacher, err := tc.repo.Find(uint(intId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if teacher.ID == 0 {
		c.JSON(404, gin.H{"error": "Teacher not found"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func (tc *TeachersController) DeleteTeacher(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = tc.repo.Delete(uint(intId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
