# Todo List API


## MVP Features

### 1. **Create a Task**
   - **Endpoint:** `POST /tasks`
   - Receives task title, description, due date, etc.
   - Saves the data into Postgres or SQLite.
   - Returns the details of the newly created task.

### 2. **Get All Tasks**
   - **Endpoint:** `GET /tasks`
   - Retrieves a list of all pending tasks.
   - Optionally support pagination and sorting (e.g., by creation time or due date).
   - Use Redis for caching to reduce frequent database queries.

### 3. **Get a Single Task**
   - **Endpoint:** `GET /tasks/:id`
   - Fetches a specific task by its ID and returns its details.
   - Consider caching the result in Redis to speed up retrieval.

### 4. **Update a Task**
   - **Endpoint:** `PUT /tasks/:id`
   - Updates a task by ID (e.g., title, description, status).
   - After updating in the database, also update Redis cache.

### 5. **Delete a Task**
   - **Endpoint:** `DELETE /tasks/:id`
   - Deletes a task by ID, removes it from the database, and clears the Redis cache.

### 6. **Mark Task as Complete**
   - **Endpoint:** `PATCH /tasks/:id/complete`
   - Changes the task status to "complete."

## Database Schema

**Tasks Table** (Postgres or SQLite)

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMP,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
