package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bidRequest struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var bidRequests = []bidRequest{
	{ID: "1", Item: "Item 1", Completed: false},
	{ID: "2", Item: "Item 2", Completed: false},
	{ID: "3", Item: "Item 3", Completed: false},
}

func getBidRequests(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, bidRequests)
}

func addBidRequest(context *gin.Context) {
	var newBidRequest bidRequest

	if err := context.BindJSON(&newBidRequest); err != nil {
		return
	}

	bidRequests = append(bidRequests, newBidRequest)
	context.IndentedJSON(http.StatusCreated, newBidRequest)
}

func getBidRequestById(id string) (*bidRequest, error) {
	for i, t := range bidRequests {
		if t.ID == id {
			return &bidRequests[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getBidRequest(context *gin.Context) {
	id := context.Param("id")
	bidRequest, err := getBidRequestById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "bid request not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, bidRequest)
}

func toggleBidRequestStatus(context *gin.Context) {
	id := context.Param("id")
	bidRequest, err := getBidRequestById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "bid request not found"})
		return
	}

	bidRequest.Completed = !bidRequest.Completed
	context.IndentedJSON(http.StatusAccepted, bidRequest)
}

func main() {
	router := gin.Default()
	router.GET("/bid_requests", getBidRequests)
	router.POST("/bid_requests", addBidRequest)
	router.GET("/bid_requests/:id", getBidRequest)
	router.PATCH("/bid_requests/:id", toggleBidRequestStatus)

	router.Run("localhost:9090")
}
