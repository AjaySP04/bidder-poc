package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BidRequest struct {
	Id       string `json:"id"`
	Imp      []Imp  `json:"imp"`
	Site     Site   `json:"site"`
	Device   Device `json:"device"`
	User     User   `json:"user"`
	Test     int    `json:"test"`
	Auction  int    `json:"auction"`
	Tmax     int    `json:"tmax"`
	Bidfloor int    `json:"bidfloor"`
}

type Imp struct {
	Id       string  `json:"id"`
	Banner   Banner  `json:"banner"`
	Video    Video   `json:"video"`
	Native   Native  `json:"native"`
	Display  Display `json:"display"`
	Amp      Amp     `json:"amp"`
	Pmp      Pmp     `json:"pmp"`
	Ext      Ext     `json:"ext"`
	Bidfloor float32 `json:"bidfloor"`
}

type Banner struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Video struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Native struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Display struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Amp struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Pmp struct {
	PrivateAuction int    `json:"private_auction"`
	Deals          []Deal `json:"deals"`
}

type Deal struct {
	Id string `json:"id"`
}

type Ext struct {
	Floor float32 `json:"floor"`
}

type Site struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Device struct {
	Id     string `json:"id"`
	Ip     string `json:"ip"`
	Model  string `json:"model"`
	Os     string `json:"os"`
	Osver  string `json:"osver"`
	Geo    Geo    `json:"geo"`
	Make   string `json:"make"`
	Width  int    `json:"w"`
	Height int    `json:"h"`
}

type User struct {
	Id string `json:"id"`
}

type Geo struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type BidResponse struct {
	Id     string   `json:"id"`
	Bidder string   `json:"bidder"`
	Seat   string   `json:"seat"`
	Price  float32  `json:"price"`
	Adm    string   `json:"adm"`
	Ext    Response `json:"ext"`
}

type Response struct {
	Creative_id string `json:"creative_id"`
}

func main() {
	router := gin.Default()

	router.POST("/bid", handleBid)

	router.Run(":9090")
}

func handleBid(c *gin.Context) {
	var req BidRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// process the bid request and generate a bid response
	bidResp := BidResponse{
		Id:     req.Id,
		Bidder: "OpenRTB Bidder",
		Seat:   "123",
		Price:  req.Imp[0].Bidfloor + 0.1,
		Adm:    "<img src='https://'>'OpenRTB Bidder</img>'",
		Ext: Response{
			Creative_id: "456",
		},
	}

	// serialize the bid response to JSON
	respJSON, err := json.Marshal(bidResp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize response to JSON"})
		return
	}

	// send the bid response back to the client
	c.Data(http.StatusOK, "application/json", respJSON)
}
