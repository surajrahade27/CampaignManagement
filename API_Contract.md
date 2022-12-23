#Campaigns
* Campaign object
```
{

	ID                  : int
	Title               : string
	Status              : string
	Order_start_date      : datetime(iso 8601)
	Order_end_date        : datetime(iso 8601)
	Collection_start_date : datetime(iso 8601)
	Collection_end_date   : datetime(iso 8601)
	Type                : string
	Swimlane_title      : string
	Swimlane_desc       : string
	Onboarding_title    : string
	Onboarding_desc     : string
	Onboarding_image    : blob
	Swimlane_item_count : int    
}
```
**GET /campaigns/getCampaingList**
----
  Returns all Campaigns in the system.
* **URL Params**  
  None
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Success Response:**  
* **Code:** 200  
  **Content:**  
```
{
  campaigns: [
           {<campaign_object>},
           {<campaign_object>},
           {<campaign_object>}
            ]
}
```

**GET /campaigns/getCampaingDetailsById/:id**
----
  Returns the specified campaign.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
   <!-- Authorization: Bearer `<OAuth Token>` -->
* **Success Response:** 
* **Code:** 200  
  **Content:**  `{ <campaign_object> }` 
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Campaign doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**POST /campaigns/createCampaign**
----
  Creates a new Campaign and returns the new object.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**  
```
  {
	Title               : string
	Status              : string
	Order_start_date      : datetime(iso 8601)
	Order_end_date        : datetime(iso 8601)
	Collection_start_date : datetime(iso 8601)
	Collection_end_date   : datetime(iso 8601)
	Type                : string
	Swimlane_title      : string
	Swimlane_desc       : string
	Onboarding_title    : string
	Onboarding_desc     : string
	Onboarding_image    : blob
	Swimlane_item_count : int  
  }
```
* **Success Response:**  
* **Code:** 200  
  **Content:**  `{ <campaign_object> }` 
