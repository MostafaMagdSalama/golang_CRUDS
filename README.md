# User Management API

This is a simple RESTful API for managing users, built using Go, Gorilla Mux, and PostgreSQL.

## Prerequisites

- Go 1.16+
- PostgreSQL
- Git

## Setup

1. Clone the repository:

   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up your environment variables:

   ```sh
   export DATABASE_URL=postgres://<user>:<password>@<host>:<port>/<dbname>?sslmode=disable
   ```

4. Run the application:

   ```sh
   go run main.go
   ```

## API Endpoints

### Get All Users

- **URL**: `/users`
- **Method**: `GET`
- **Description**: Retrieves a list of all users.

**Request**:

    ```sh
    curl -X GET http://localhost:8000/users
    ```

**Response**:

    ```json
    [
        {
            "id": "1",
            "name": "John Doe",
            "email": "john.doe@example.com"
        }
    ]
    ```

### Get User by ID

- **URL**: `/user/{id}`
- **Method**: `GET`
- **Description**: Retrieves a user by their ID.

**Request**:

    ```sh
    curl -X GET http://localhost:8000/user/1
    ```

**Response**:

    ```json
    {
        "id": "1",
        "name": "John Doe",
        "email": "john.doe@example.com"
    }
    ```

### Create User

- **URL**: `/user`
- **Method**: `POST`
- **Description**: Creates a new user.

**Request**:

    ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"name": "Jane Doe", "email": "jane.doe@example.com"}' http://localhost:8000/user
    ```

**Response**:

    ```json
    {
        "id": "2",
        "name": "Jane Doe",
        "email": "jane.doe@example.com"
    }
    ```

### Delete User by ID

- **URL**: `/user/{id}`
- **Method**: `DELETE`
- **Description**: Deletes a user by their ID.

**Request**:

    ```sh
    curl -X DELETE http://localhost:8000/user/1
    ```

**Response**:

    ```json
    "User Deleted"
    ```

### Update User

- **URL**: `/user/{id}`
- **Method**: `PUT`
- **Description**: Updates an existing user.

**Request**:

    ```sh
    curl -X PUT -H "Content-Type: application/json" -d '{"name": "John Smith", "email": "john.smith@example.com"}' http://localhost:8000/user/1
    ```

**Response**:

    ```json
    {
        "id": "1",
        "name": "John Smith",
        "email": "john.smith@example.com"
    }
    ```

## Code Structure

- **main.go**: The main file containing the entry point and router setup.
- **User struct**: Defines the user model.
- **Database Functions**: Functions to interact with the PostgreSQL database.
- **Handler Functions**: Functions to handle HTTP requests and responses.

## Middleware

- **jsonContentTypeMiddleware**: Ensures the Content-Type of all responses is set to `application/json`.

## Error Handling

Errors are logged using the `log.Fatal` function, which terminates the application if an error occurs.

## Future Improvements

- Add proper error handling and return appropriate HTTP status codes.
- Implement authentication and authorization.
- Add request validation.
- Write unit tests.

## License

This project is licensed under the MIT License.
