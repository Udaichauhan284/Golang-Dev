package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udaichauhan/url_shortener/api/database"
)

func DeleteURL(c *gin.Context){
	shortID := c.Param("shortID");

	// r := database.CreateClient(0);
	// defer r.Close();

	err := database.Client.Del(database.Ctx, shortID).Err();

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Unable to Delete shortened Link",
		});
		return;
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Shortend URL Deleted Successfully",
	});
}