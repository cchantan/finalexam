package todo

import (
	"finalexam/customerdatabase"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Cust struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
type Custhandler struct{}

func (Custhandler) GetCustomersHandler(c *gin.Context) {
	db, err := customerdatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT  id, name, email, status FROM customers")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	custs := []Cust{}
	t := Cust{}
	i := 1
	for rows.Next() {
		err := rows.Scan(&t.ID, &t.Name, &t.Email, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}
		fmt.Println("Test for One row => ", i, "\nID     = ", t.ID, "\nName  = ", t.Name, "\nEmail  = ", t.Email, "\nStatus = ", t.Status)
		custs = append(custs, t)
		i++

	}
	fmt.Println(custs)
	c.JSON(200, custs)
	return
}

func (Custhandler) GetCustomersByIdHandler(c *gin.Context) {
	idinput := c.Param("id")
	db, err := customerdatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT  id, name, email, status FROM customers WHERE id=$1")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	row := stmt.QueryRow(idinput)
	t := Cust{}
	err = row.Scan(&t.ID, &t.Name, &t.Email, &t.Status)
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	fmt.Println("Test for One row => ", "\n ID     = ", t.ID, "\n Name  = ", t.Name, "\n Email  = ", t.Email, "\n Status = ", t.Status)
	fmt.Println(t)
	c.JSON(200, t)
	return
}

func (Custhandler) PostCustomersHandler(c *gin.Context) {
	t := Cust{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(t)

	db, err := customerdatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()
	name := t.Name
	email := t.Email
	status := t.Status
	query := `
		INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id
	`
	var id int
	row := db.QueryRow(query, name, email, status)
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't scan id", id)
	}
	fmt.Println("Insert success id", id)
	t.ID = id
	c.JSON(201, t)
}

func (Custhandler) DeleteCustomersByIdHandler(c *gin.Context) {
	custs := []Cust{}
	idinput := c.Param("id")

	fmt.Println(idinput)

	db, err := customerdatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	query := `
		DELETE FROM customers WHERE id =$1
	`
	db.QueryRow(query, idinput)
	fmt.Println(custs)
	c.JSON(200, gin.H{
		"message": "customer deleted",
	})
	return
}

func (Custhandler) PutCustomersByIdHandler(c *gin.Context) {
	idinput := c.Param("id")

	t := Cust{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(t)

	db, err := customerdatabase.GetDBConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	query := `
	UPDATE custs SET name = $2, email = $3, status = $4  WHERE id=$1
	`
	db.QueryRow(query, idinput, t.Name, t.Email, t.Status)

	t.ID, err = strconv.Atoi(idinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, t)
	return
}
