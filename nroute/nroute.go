package nroute

import (
	"github.com/cchantan/finalexam/customerdatabase"
	"github.com/cchantan/finalexam/middleware"
	"github.com/cchantan/finalexam/todo"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Nroute() *gin.Engine {
	s := todo.Custhandler{}
	r := gin.Default()

	// Middleware use for every functions
	r.Use(middleware.Authmiddleware)
	//
	err := customerdatabase.Createtable()
	if err != nil {
		return nil
	}
	r.GET("/customers/", s.GetCustomersHandler)
	r.GET("/customers/:id", s.GetCustomersByIdHandler)
	r.POST("/customers/", s.PostCustomersHandler)
	r.DELETE("/customers/:id", s.DeleteCustomersByIdHandler)
	r.PUT("/customers/:id", s.PutCustomersByIdHandler)
	return r
}
