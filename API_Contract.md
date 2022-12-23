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
**GET /getCampaigns**
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

**GET /campaigns/:id**
----
  Returns the specified campaign.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:** 
* **Code:** 200  
  **Content:**  `{ <campaign_object> }` 
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Campaign doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**GET /campaigns/:id/orders**
----
  Returns all Orders associated with the specified campaign.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**  
* **Code:** 200  
  **Content:**  
```
{
  orders: [
           {<order_object>},
           {<order_object>},
           {<order_object>}
         ]
}
```
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Campaign doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**POST /campaigns**
----
  Creates a new Campaign and returns the new object.
* **URL Params**  
  None
* **Headers**  
  Content-Type: application/json  
* **Data Params**  
```
  {
    campaignname: string,
    email: string
  }
```
* **Success Response:**  
* **Code:** 200  
  **Content:**  `{ <campaign_object> }` 

**PATCH /campaigns/:id**
----
  Updates fields on the specified campaign and returns the updated object.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
```
  {
  	campaignname: string,
    email: string
  }
```
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:** 
* **Code:** 200  
  **Content:**  `{ <campaign_object> }`  
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Campaign doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**DELETE /campaigns/:id**
----
  Deletes the specified campaign.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:** 
  * **Code:** 204 
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Campaign doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

#Products
* Product object
```
{
  id: integer
  name: string
  cost: float(2)
  available_quantity: integer
  created_at: datetime(iso 8601)
  updated_at: datetime(iso 8601)
}
```
**GET /products**
----
  Returns all products in the system.
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
  products: [
           {<product_object>},
           {<product_object>},
           {<product_object>}
         ]
}
``` 

**GET /products/:id**
----
  Returns the specified product.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**  
* **Code:** 200  
  **Content:**  `{ <product_object> }` 
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Product doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**GET /products/:id/orders**
----
  Returns all Orders associated with the specified product.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:** 
* **Code:** 200  
  **Content:**  
```
{
  orders: [
           {<order_object>},
           {<order_object>},
           {<order_object>}
         ]
}
``` 
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Product doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**POST /products**
----
  Creates a new Product and returns the new object.
* **URL Params**  
  None
* **Data Params**  
```
  {
    name: string
    cost: float(2)
    available_quantity: integer
  }
```
* **Headers**  
  Content-Type: application/json  
* **Success Response:**  
* **Code:** 200  
  **Content:**  `{ <product_object> }` 

**PATCH /products/:id**
----
  Updates fields on the specified product and returns the updated object.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
```
  {
  	name: string
    cost: float(2)
    available_quantity: integer
  }
```
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:** 
* **Code:** 200  
  **Content:**  `{ <product_object> }`  
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Product doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**DELETE /products/:id**
----
  Deletes the specified product.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**  
  * **Code:** 204
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Product doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

#Orders
* Order object
```
{
  id: integer
  campaign_id: <campaign_id>
  total: float(2)
  products: [
              { 
                product: <product_id>,
                quantity: integer 
              },
              { 
                product: <product_id>,
                quantity: integer 
              },
              { 
                product: <product_id>,
                quantity: integer 
              },
            ]
  created_at: datetime(iso 8601)
  updated_at: datetime(iso 8601)
}
```
**GET /orders**
----
  Returns all campaigns in the system.
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
  orders: [
           {<order_object>},
           {<order_object>},
           {<order_object>}
         ]
}
``` 

**GET /orders/:id**
----
  Returns the specified order.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**  
* **Code:** 200  
  **Content:**  `{ <order_object> }` 
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Order doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**GET /orders/:id/products**
----
  Returns all Products associated with the specified order.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**  
* **Code:** 200  
  **Content:**  
```
{
  products: [
           {<product_object>},
           {<product_object>},
           {<product_object>}
         ]
}
```
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Order doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**GET /orders/:id/campaign**
----
  Returns all Campaigns associated with the specified order.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:** `{ <campaign_object> }`  
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Order doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**POST /orders**
----
  Creates a new Order and returns the new object.
* **URL Params**  
  None
* **Data Params**  
```
  {
  	campaign_id: <campaign_id>
  	product: <product_id>,
  	quantity: integer 
  }
```
* **Headers**  
  Content-Type: application/json  
* **Success Response:**  
* **Code:** 200  
  **Content:**  `{ <order_object> }` 

**PATCH /orders/:id**
----
  Updates fields on the specified order and returns the updated object.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
```
  {
  	product: <product_id>,
  	quantity: integer 
  }
```
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:**  
* **Code:** 200  
  **Content:**  `{ <order_object> }` 
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Order doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`

**DELETE /orders/:id**
----
  Deletes the specified order.
* **URL Params**  
  *Required:* `id=[integer]`
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<OAuth Token>`
* **Success Response:** 
  * **Code:** 204 
* **Error Response:**  
  * **Code:** 404  
  **Content:** `{ error : "Order doesn't exist" }`  
  OR  
  * **Code:** 401  
  **Content:** `{ error : error : "You are unauthorized to make this request." }`