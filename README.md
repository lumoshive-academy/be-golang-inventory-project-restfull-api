#  RESTful API request and response structure for each feature inventory system

## Success Response Format
```json
{
  "success": true,
  "message": "Data processed successfully.",
  "data": {
    // object or array depending on context
  }
}
```

##  Error Response Formats

###  Not Found – 404
```json
{
  "success": false,
  "message": "Data not found."
}
```

###  Unauthorized – 401
```json
{
  "success": false,
  "message": "Unauthorized. Please log in first."
}
```

### Forbidden – 403
```json
{
  "success": false,
  "message": "You do not have permission to access this feature."
}
```

### Server Error – 500
```json
{
  "success": false,
  "message": "An internal server error occurred."
}
```

### Validation Failed – 400
```json
{
  "success": false,
  "message": "Validation failed.",
  "errors": {
    "field": ["Invalid input."]
  }
}
```

##  Feature manage product
### Get All Product
```json
{
  "success": true,
  "message": "Products retrieved successfully.",
  "data": [
    {
      "id": 1,
      "name": "Logitech Mouse",
      "category_id": 1,
      "rack_id": 1,
      "warehouse_id": 1,
      "inventory_count": 100,
      "retail_price": 90000,
      "selling_price": 120000,
      "image": "mouse.jpg"
    },
    {
      "id": 2,
      "name": "Mechanical Keyboard",
      "category_id": 1,
      "rack_id": 2,
      "warehouse_id": 1,
      "inventory_count": 50,
      "retail_price": 250000,
      "selling_price": 300000,
      "image": "keyboard.jpg"
    },
    {
      "id": 3,
      "name": "USB Flash Drive 64GB",
      "category_id": 2,
      "rack_id": 3,
      "warehouse_id": 2,
      "inventory_count": 200,
      "retail_price": 80000,
      "selling_price": 95000,
      "image": "usb.jpg"
    }
  ]
}
```

### Create Product
```json
{
  "name": "Logitech Mouse",
  "category_id": 1,
  "rack_id": 1,
  "warehouse_id": 1,
  "inventory_count": 100,
  "retail_price": 90000,
  "selling_price": 120000,
  "image": "mouse.jpg"
}
```

##  Feature manage Product Categories
### Create Category
```json
{
  "name": "Electronics",
  "description": "All electronic devices and accessories"
}
```

##  Feature manage Product Racks
### Create rack
```json
{
  "name": "Rack A",
  "description": "Storage for electronics",
}
```

##  Feature manage Warehouses
### Create Warehouse
```json
{
  "name": "Main Warehouse",
  "description": "Primary storage facility"
}
```

##  Feature manage Users
### Create User
```json
{
  "name": "Raka",
  "email": "raka@email.com",
  "role": "admin",
  "password": "securePassword123",
  "status": "active"
}
```

##  Feature manage Sales
### Create Sales
```json
{
  "product_id": 3,
  "item_sold": 5,
  "total_bill": 500000,
  "date_sale": "2025-06-30"
}
```

##  Feature Product Report
### Response Get Report
```json
{
  "success": true,
  "message": "Report retrieved successfully.",
  "data": {
    "total_product": 150,
    "total_sale": 320,
    "total_revenue": 75000000
  }
}
```

### Request Get Report
GET /report/summary?start_date=2025-06-01&end_date=2025-06-30
