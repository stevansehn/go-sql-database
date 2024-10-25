# Go SQL Database

A Go program that connects to a MySQL database, performs CRUD operations on a `users` table, and handles SQL query results. This application allows you to insert new users, retrieve user information, and count the total number of users in the database.

## Features

- Connects to a MySQL database.
- Inserts new users with a username and password.
- Retrieves the last inserted user and the total count of users.
- Displays user information in a structured format.
- Ensures secure password handling (consider hashing in future updates).

## Prerequisites

- Go (version 1.16 or later)
- MySQL server
- Go MySQL driver: `github.com/go-sql-driver/mysql`
- Environment variable loader: `github.com/joho/godotenv`

## Installation

1. **Clone the repository:**
   ```bash
   git clone git@github.com:stevansehn/go-sql-database.git
   cd go-sql-database
   ```

2. **Install dependencies:**
   Make sure you have Go installed, then run:
   ```bash
   go mod tidy
   ```

3. **Set up your MySQL database:**
   - Create a new database (e.g., `go_sql_database`).
   - Update the `.env` file with your database credentials:
     ```
     DB_USER=your_db_user
     DB_PASSWORD=your_db_password
     DB_HOST=localhost
     DB_PORT=3306
     DB_NAME=go_sql_database
     DB_PARSE_TIME=true
     ```

4. **Create the `users` table:**
   The application will automatically create the `users` table if it does not exist when you run the program.

## Usage

1. **Run the application:**
   ```bash
   go run main.go
   ```

2. **Expected Output:**
   Upon running the application, you should see output similar to:
   ```
   Connected to database: go_sql_database
   Inserted user ID: 1
   Total users: 1
   Last inserted user: 1 johndoe secret 2024-10-25 14:53:59 +0000 UTC
   ```

3. **Database Queries:**
   You can query the `users` table directly using MySQL commands:
   ```sql
   SELECT * FROM users;
   ```

## Future Improvements

- Implement password hashing for secure storage.
- Add user authentication and session management.
- Enhance error handling and logging.
- Create a RESTful API for user management.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
