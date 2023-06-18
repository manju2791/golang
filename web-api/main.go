package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/courses", getCourses)
	router.GET("/courses/:course", getCourseByName)
	router.POST("/courses", addCourse)
	// router.PUT("/courses", updateCourse)
	router.DELETE("/courses/:course_id", removeCourse)

	router.Run("localhost:8080")

}
