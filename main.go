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

func getPlaneById(c *gin.Context) {
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
	router.GET("/planes/:id", getPlaneById)
	router.POST("/planes", postPlanes)

	router.Run("localhost:8080")
}

// create a put route

// create a delete route

// Prepare a resp to get all of the planes(items)

// postPlanes adds an plane from JSON received in the request body.
func postPlanes(c *gin.Context) {
	var newPlane plane

	// Call BindJSON to bind the received JSON to
	// newPlane.
	if err := c.BindJSON(&newPlane); err != nil {
		return
	}

	// Add the new plane to the slice.
	planes = append(planes, newPlane)
	c.IndentedJSON(http.StatusCreated, newPlane)
}

// getPlanesById locates the plane whose ID value matches the ID
// parameter sent by the client, then returns that plane as a respones.
func getPlaneByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of planes, looking for a plane whose ID value matches the parameter.
	for _, a := range planes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "plane not found"})
}
