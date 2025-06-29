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