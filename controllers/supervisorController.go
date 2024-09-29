package controllers

import (
	"foundry/dto"
	"foundry/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ScrapeInput(c *gin.Context) {
	var scrapeInput dto.ScrapeInput
	if err := c.ShouldBindJSON(&scrapeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	stockExisting, err := models.Find(scrapeInput.Name)
	if err == gorm.ErrRecordNotFound {
		stock := models.Stocks{
			Name:     scrapeInput.Name,
			Quantity: scrapeInput.Quantity - scrapeInput.BagWeight - scrapeInput.MiscellaneousWeight,
		}
		stock.Create()
	}

	if scrapeInput.Movement {
		scrapeInput.NetWeight = scrapeInput.Quantity - scrapeInput.BagWeight - scrapeInput.MiscellaneousWeight
		if stockExisting != nil {
			stockExisting.Quantity += scrapeInput.NetWeight
		}
	} else {
		scrapeInput.NetWeight = scrapeInput.Quantity + scrapeInput.BagWeight + scrapeInput.MiscellaneousWeight
		if stockExisting != nil {
			stockExisting.Quantity -= scrapeInput.Quantity
		}
	}

	stockExisting.Update()

	scrapeInputModel := models.Purchase{
		Name:                scrapeInput.Name,
		SupplierName:        scrapeInput.SupplierName,
		Movement:            scrapeInput.Movement,
		Date:                scrapeInput.Date,
		GrossWeight:         scrapeInput.Quantity,
		BagWeight:           scrapeInput.BagWeight,
		MiscellaneousWeight: scrapeInput.MiscellaneousWeight,
		NetWeight:           scrapeInput.NetWeight,
	}
	scrapeInputModel.Create()
	c.JSON(http.StatusOK, gin.H{"success": "Data Entered"})
}
