package controllers

// func AddStudentToClass(c *gin.Context) {
// 	classId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, "invalid id")
// 		return
// 	}
// 	studentId, err := strconv.Atoi(c.Param("studentId"))
// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, "invalid id")
// 		return
// 	}
// 	studentClass := models.StudentClasses{ClassID: uint(classId), StudentID: uint(studentId)}
// 	db.Db.Create(&studentClass)
// 	c.JSON(http.StatusOK, studentClass)
// }

// func RemoveStudentFromClass(c *gin.Context) {

// }
