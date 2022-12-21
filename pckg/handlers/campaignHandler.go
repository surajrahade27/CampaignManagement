package handlers

import (
	"campaing_management/pckg/datastore"
	"campaing_management/pckg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetCampaings(c *gin.Context) {
	var Campaings []models.Campaing
	datastore.DB.Find(&Campaings)
	if Campaings == nil || len(Campaings) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, Campaings)
	}
}

func ReadJson(c *gin.Context) {

	var campaings []models.Campaing
	jsonFile, err := os.Open("./MOCK_DATA.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened campaings.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &campaings)
	// fmt.Println(campaing)

	for i, s := range campaings {
		datastore.DB.Create(&s)
		fmt.Println(i, s)
	}
	c.IndentedJSON(http.StatusOK, campaings)

}

func GetCampaing(c *gin.Context) {
	var campaing models.Campaing
	datastore.DB.First(&campaing, c.Param("Id"))
	c.IndentedJSON(http.StatusOK, campaing)
}
func DeleteCampaing(c *gin.Context) {
	var campaing models.Campaing

	datastore.DB.First(&campaing, c.Param("id"))
	if err := datastore.DB.Where("id = ?", c.Param("id")).First(&campaing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if campaing.ID != 0 {
		datastore.DB.Delete(&campaing)
		c.JSON(200, gin.H{"success": "campaing #" + c.Param("id") + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}

}
func CreateCampaing(c *gin.Context) {

	var input models.Campaing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Create book
	campaing := models.Campaing{ID: input.ID,
		Title:                 input.Title,
		Status:                input.Status, //after publish campaign status will be set as true
		Order_start_date:      input.Order_start_date,
		Order_end_date:        input.Order_end_date,
		Collection_start_date: input.Collection_start_date,
		Collection_end_date:   input.Collection_end_date,
		Swimlane_title:        input.Swimlane_title,
		Swimlane_desc:         input.Swimlane_desc,
		Swimlane_item_count:   input.Swimlane_item_count,
	}
	datastore.DB.Create(&campaing)
	c.JSON(http.StatusOK, campaing)
}
func UpdateCampaing(c *gin.Context) {

	var campaing models.Campaing
	datastore.DB.First(&campaing, c.Param("id"))
	if err := datastore.DB.Where("id = ?", c.Param("id")).First(&campaing).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Campaing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	datastore.DB.Model(&campaing).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": campaing})

}
