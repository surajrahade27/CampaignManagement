package models

import "gorm.io/gorm"

// var DB *gorm.DB
// var err error
// const DNS = "root:Password@1@tcp(localhost:3306)/FairPrise?parseTime=true"

type Campaing struct {
	gorm.Model

	ID                    int    `gorm:"primaryKey" json:"ID"`
	Title                 string `json:"title"`
	Status                string `json:"status"`
	Order_start_date      string `json:"order_start_date"`
	Order_end_date        string `json:"order_end_date"`
	Collection_start_date string `json:"collection_start_date"`
	Collection_end_date   string `json:"collection_end_date"`
	Type                  string `json:"type"`
	Swimlane_title        string `json:"swimlane_title"`
	Swimlane_desc         string `json:"swimlane_desc"`
	Onboarding_title      string `json:"onboarding_title"`
	Onboarding_desc       string `json:"onboarding_desc"`
	Onboarding_image      string `json:"onboarding_image"`
	Swimlane_item_count   int    `json:"swimlane_item_count"`
}

// type Product struct {
// 	gorm.Model

// 	ID           int     `gorm:"primaryKey" json:"id"`
// 	Descripstion string  `json:"descripstion"`
// 	Sku_number   uint32  `json:"Sku_number"`
// 	Name         string  `json:"Name"`
// 	Category     string  `json:"Category"`
// 	Unit_price   float32 `json:"Unit_price"`
// 	Lead_time    int     `json:"Lead_time"`
// }
// type Store struct {
// 	gorm.Model

// 	ID           int    `gorm:"primaryKey" json:"id"`
// 	Name         string `json:"Name"`
// 	Type         string `json:"Type"`
// 	Location     string `json:"Location"`
// 	Is_pre_order bool   `json:"Is_pre_order"`
// }

// func InitialMigration() {
// 	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Cannot connect to DB")
// 	}
// 	var models = []interface{}{&Campaing{}, &OnBoard{}}

// 	DB.AutoMigrate(models...)
// 	// DB.AutoMigrate(&Campaing{})
// 	// DB.AutoMigrate(&OnBoard{})
// }

// func GetCampaings(c *gin.Context) {
// 	var Campaings []Campaing
// 	DB.Find(&Campaings)
// 	if Campaings == nil || len(Campaings) == 0 {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.IndentedJSON(http.StatusOK, Campaings)
// 	}
// }

// func ReadJson(c *gin.Context) {

// 	var campaings []Campaing
// 	jsonFile, err := os.Open("./MOCK_DATA.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Successfully Opened campaings.json")
// 	defer jsonFile.Close()
// 	byteValue, _ := io.ReadAll(jsonFile)

// 	json.Unmarshal(byteValue, &campaings)
// 	// fmt.Println(campaing)

// 	for i, s := range campaings {
// 		DB.Create(&s)
// 		fmt.Println(i, s)
// 	}
// 	c.IndentedJSON(http.StatusOK, campaings)

// }

// func GetCampaing(c *gin.Context) {
// 	var campaing Campaing
// 	DB.First(&campaing, c.Param("Id"))
// 	c.IndentedJSON(http.StatusOK, campaing)
// }
// func DeleteCampaing(c *gin.Context) {
// 	var campaing Campaing

// 	DB.First(&campaing, c.Param("id"))
// 	if err := DB.Where("id = ?", c.Param("id")).First(&campaing).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	if campaing.ID != 0 {
// 		DB.Delete(&campaing)
// 		c.JSON(200, gin.H{"success": "campaing #" + c.Param("id") + " deleted"})
// 	} else {
// 		c.JSON(404, gin.H{"error": "User not found"})
// 	}

// }
// func CreateCampaing(c *gin.Context) {

// 	var input Campaing
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	// Create book
// 	campaing := Campaing{ID: input.ID, Name: input.Name, Title: input.Title, Status: input.Status, StartDate: input.StartDate, EndDate: input.EndDate, OrderDate: input.OrderDate, CollectionDate: input.CollectionDate}
// 	DB.Create(&campaing)
// 	c.JSON(http.StatusOK, campaing)
// }
// func UpdateCampaing(c *gin.Context) {

// 	var campaing Campaing
// 	DB.First(&campaing, c.Param("id"))
// 	if err := DB.Where("id = ?", c.Param("id")).First(&campaing).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input Campaing
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	DB.Model(&campaing).Updates(input)

// 	c.JSON(http.StatusOK, gin.H{"data": campaing})

// }
