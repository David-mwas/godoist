# Godoist

Godoist is a to-do application built with **Go** and **MongoDB**. It serves as a simple task management tool where users can add, view, update, and delete tasks via an API. The application is deployed on **Vercel**, providing a scalable and efficient solution for handling to-do tasks.

## Live Demo

- **Go Server (API)**: [https://godoist.vercel.app/](https://godoist.vercel.app/)
- **Frontend (Vite Client)**: [Client URL on Vercel or your desired link]

## Repository

- **GitHub**: [https://github.com/David-mwas/godoist/](https://github.com/David-mwas/godoist/)

## Folder Structure

```plaintext
godoist/
├── api/                        # API logic and handlers
│   ├── index.go                # Main entry point for the application (Serverless function)
│   └── handler/                # Folder containing individual handler functions
├── client/                     # Frontend build folder for Vite client (static files served)
├── .env                        # Environment variables configuration
├── go.mod                      # Go module file
├── go.sum                      # Go checksum file
└── README.md                   # Project documentation (this file)
|___ vercel.json                #  "rewrites": [{ "source": "(.*)", "destination": "api/index.go" }]

```

## API Endpoints

### 1. Get All Todos

- **Method**: `GET`
- **Endpoint**: `/api/todo`
- **Response**:
  ```json
  [
    {
      "_id": "unique_id",
      "completed": false,
      "body": "Sample Task"
    }
  ]
  ```

### 2. Add a New Todo

- **Method**: `POST`
- **Endpoint**: `/api/todo`
- **Request Body**:
  ```json
  {
    "body": "New Task"
  }
  ```
- **Response**:
  ```json
  {
    "_id": "unique_id",
    "completed": false,
    "body": "New Task"
  }
  ```

### 3. Update a Todo

- **Method**: `PATCH`
- **Endpoint**: `/api/todo/:id`
- **Request**: The `id` parameter represents the task ID to update.
- **Response**:
  ```json
  {
    "success": true
  }
  ```

### 4. Delete a Todo

- **Method**: `DELETE`
- **Endpoint**: `/api/todo/:id`
- **Request**: The `id` parameter represents the task ID to delete.
- **Response**:
  ```json
  {
    "success": true
  }
  ```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/David-mwas/godoist.git
   cd godoist
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your environment variables in the `.env` file:

   - `MONGODB_URI`: Your MongoDB connection string.
   - `MONGODB_DB`: The name of the MongoDB database.

4. Run the server locally:

   ```bash
   go run api/index.go
   ```

5. The server will be available at `http://localhost:3000`.

## Deployment on Vercel

1. Create a Vercel account and log in.
2. Connect your GitHub repository to Vercel.
3. In your Vercel project, specify the following settings for deployment:
   - **Build Command**: `go build -o api/index.go`
   - **Output Directory**: `api/`
4. Deploy your project.

Once deployed, Vercel will automatically provide you with a URL like `https://godoist.vercel.app/` for your live server.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
