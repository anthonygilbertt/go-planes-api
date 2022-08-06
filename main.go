package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// plane represents data about a plane.
type plane struct {
	ID           string `json:"id"`
	PlaneName    string `json:"planename"`
	Manufacturer string `json:"manufacturer"`
	Year         int16  `json:"year"`
}

// getPlanes responds with the list of all planes as JSON.
func getPlanes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, planes)
}

// planes slice to seed plane data.
var planes = []plane{
	{ID: "1", Manufacturer: "Boeing", PlaneName: "777", Year: 2011},
	{ID: "2", Manufacturer: "Boeing", PlaneName: "747", Year: 1968},
	{ID: "3", Manufacturer: "Lockheed Martins", PlaneName: "Blackbird", Year: 1966},
}

// Assign the handler function to an endpoint path
// This sets up an association in which getPlanes handles requests to
// the /planes endpoint path
func main() {
	router := gin.Default()
	router.GET("/planes", getPlanes)

	router.Run("localhost:8080")
}

// create a put route

// create a delete route

// Prepare a resp to get all of the planes(items)
