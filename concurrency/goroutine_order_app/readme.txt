Order App Functionality
* View all products and their amount of stock
* Place a new order if it is valid and we have enough stock
* View an existing order

To achieve the above, we need four Order App EndPoints
* GET/
- The get root end Points simpy get the hardcoded response message
* GET/products
- Show us all the products information
* POST/orders
- Create new orders
* GET/orders/{orderId}
- Retrieve the existing order by referencing the orderId in the URL
