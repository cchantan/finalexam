package nroute

import (
	"finalexam/customerdatabase"
	"finalexam/middleware"
	"finalexam/todo"

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
	//r.GET("/customers/:id", s.Checkauthgetcustomer)
	return r
}
