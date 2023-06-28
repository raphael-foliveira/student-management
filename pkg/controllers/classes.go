package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/models"
	"github.com/raphael-foliveira/studentManagementSystem/pkg/repository"
)

type ClassesController struct {
	repo repository.Repository[models.Class]
}

func NewClassesController(repo repository.Repository[models.Class]) *ClassesController {
	return &ClassesController{repo: repo}
}

func (cc *ClassesController) List(c *gin.Context) {
	classes, err := cc.repo.All()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, classes)
}

func (cc *ClassesController) Retrieve(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	class, err := cc.repo.Find(uint(intId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if class.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	c.JSON(http.StatusOK, class)
}

func (cc *ClassesController) Create(c *gin.Context) {
	var newClass models.Class
	if err := c.BindJSON(&newClass); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := cc.repo.Create(&newClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newClass)
}

func (cc *ClassesController) Delete(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = cc.repo.Delete(uint(intId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
