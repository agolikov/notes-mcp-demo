Journal Notes Application API
===================
This is a journal notes application that allows users to create, organize, and manage their personal notes and entries.

Requirements

Note Data Structure:
- Note must contain the following fields:
  - title (Note title)
  - content (Main note content)
  - tags (List of tags for categorization)
  - created_at (Note creation date)
  - updated_at (Last modification date)

Tags:
- System must support the following note tags:
  - Personal
  - Work
  - Travel
  - Health
  - Ideas
  - Recipes
  - Memories
  - Goals
  - Finance
  - Education

Note Management API:
- [ ] As a user, I want a dedicated endpoint (/api/v1/notes) to create and manage notes.
- [ ] As a user, I want the note API to return structured data with metadata about each note.

Note Processing and Organization:
- [ ] As a user, I want to create new notes with rich text formatting.
- [ ] As a user, I want to review and edit my notes before saving.
- [ ] As a user, I want to see a preview of my note with formatting.

Note Management:
- [ ] As a user, I want to save my notes into the system, so I can keep track of my thoughts.
- [ ] As a user, I want to view all my notes in a list, so I can see my journal history.
- [ ] As a user, I want to filter notes by:
  - Date range
  - Tags

  So that I can organize my notes effectively.
- [ ] As a user, I want to sort notes by:
  - Date (ascending/descending)
  - Title (alphabetical)
  - Tags

Data Management:
- [ ] As a user, I want to edit any field of a saved note.
- [ ] As a user, I want to categorize my notes using predefined tags.
- [ ] As a user, I want to see a summary of my notes by category and time period.
- [ ] As a user, I want to access a list of all available tags through an API endpoint.