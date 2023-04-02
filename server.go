package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Health Check api
func forceCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service is force=)!")
}

// Path Parameters
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// Query Parameters
func showUser(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// multipart/form-data
func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// Define the routes
	e.GET("/force", forceCheck)
	e.GET("/getuser/:id", getUser) // Path Parameters
	e.GET("/showuser", showUser)   // Query Parameters
	e.POST("/saveUser", saveUser)  // multipart/form-data
	// this line should be last
	e.Logger.Fatal(e.Start(":1323"))
}
