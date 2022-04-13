package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.POST("/deck/create", create)
	r.GET("/deck/open", open)
	r.POST("/deck/draw", draw)
	r.Run()
}

func create(c *gin.Context) {
	partial, _ := c.GetQueryArray("cards")
	_, shuffle := c.GetQuery("shuffle")
	baseDeck := CreateDeck(shuffle, partial)
	c.JSON(http.StatusOK, gin.H{
		"id":        baseDeck.Id,
		"shuffled":  baseDeck.Shuffled,
		"remaining": len(baseDeck.Cards),
	})
}

func open(c *gin.Context) {
	deckId := c.Query("id")
	deck, err := DeckExists(deckId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"id":        deck.Id,
			"shuffled":  deck.Shuffled,
			"remaining": len(deck.Cards),
			"cards":     deck.Cards,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

}

func draw(c *gin.Context) {
	deckId := c.Query("id")
	amount, _ := strconv.Atoi(c.Query("amount"))
	deck, err := DeckExists(deckId)
	if err == nil {
		draw := Draw(deck, amount)
		c.JSON(http.StatusOK, gin.H{
			"cards": draw,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

}
