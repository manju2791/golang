package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type course struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Professor string  `json:"professor"`
	Price     float64 `json:"price"`
}

var courses = []course{
	{ID: "1", Title: "GoLang", Professor: "Jay Ney", Price: 56.99},
	{ID: "2", Title: "Python", Professor: "Joe Nair", Price: 17.99},
	{ID: "3", Title: "Java", Professor: "Jane Nick", Price: 39.99},
}

func getCourses(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, courses)
}

func getCourseByName(c *gin.Context) {
	course_name := c.Param("course")
	for _, course := range courses {
		if course.Title == course_name {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Course not found"})
}

func addCourse(c *gin.Context) {

	var newCourse course
	if err := c.BindJSON(&newCourse); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request found. Returning."})
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Course created"})
}

func removeCourse(c *gin.Context) {

	course_id := c.Param("course_id")

	for idx, course := range courses {
		if course.ID == course_id {
			courses = append(courses[:idx], courses[idx+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Course deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Course not found"})
}
