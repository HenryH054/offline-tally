package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tally struct {
	ID     string `json:"id"`
	Time   string `json:"time"`
	Adjust int    `json:"adjust"`
}

var tallies = []tally{
	{ID: "1", Time: "01-01-01-00:00", Adjust: 4},
	{ID: "2", Time: "01-02-01-00:00", Adjust: 6},
}

func main() {
	router := gin.Default()
	router.GET("/tallies", getTallies)
	router.POST("/postTallies", postTallies)

	router.Run("localhost:8080")
}

func getTallies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tallies)
}

func postTallies(c *gin.Context) {
	var newTally tally

	if err := c.BindJSON(&newTally); err != nil {
		err := errors.New("Unable to decode JSON")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	tallies = append(tallies, newTally)
	c.IndentedJSON(http.StatusCreated, newTally)
}
