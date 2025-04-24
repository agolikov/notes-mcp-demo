# Application Entities

## Note Entity
- id (UUID, primary key)
- created_at (Timestamp)
- updated_at (Timestamp)
- deleted_at (Timestamp, nullable)
- title (String)
- content (Text)
- tags (Array of Strings)
- user_id (UUID, foreign key to User)
- attachments (Array of UUIDs, foreign keys to Attachment)

## Tag Entity
- id (UUID, primary key)
- name (String, unique)
- description (String, nullable)
- created_at (Timestamp)
- updated_at (Timestamp)
- deleted_at (Timestamp, nullable)