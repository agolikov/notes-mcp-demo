# Journal Notes API Specification

## Base URL
`/api/v1`

## Endpoints

### Notes Management
- **GET /notes**  
  List all notes with optional filtering and sorting  
  Query Parameters:  
    - `start_date` (optional): Filter by creation date range start  
    - `end_date` (optional): Filter by creation date range end  
    - `tags` (optional): Filter by tags (comma-separated)  
    - `sort` (optional): Sort by field (date, title)  
    - `order` (optional): Sort order (asc, desc)

- **POST /notes**  
  Create a new note  
  Request Body:  
    - `title`: Note title  
    - `content`: Note content  
    - `tags`: Array of tags  
    - `attachments`: Array of attachment IDs (optional)

- **GET /notes/{id}**  
  Get a specific note by ID

- **PUT /notes/{id}**  
  Update an existing note  
  Request Body: Same as POST /notes

- **DELETE /notes/{id}**  
  Delete a note

### Tags Management
- **GET /tags**  
  List all available tags

## Response Structure
All endpoints return JSON responses with the following structure:
```json
{
  "data": {}, // Response data
  "meta": {
    "timestamp": "ISO-8601 datetime",
    "status": "success/error"
  }
}