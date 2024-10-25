package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

type User struct {
    ID        int
    Username  string
    Password  string
    CreatedAt time.Time
}

// Helper function for error handling
func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    // Load the .env file
    err := godotenv.Load()
    checkErr(err)

    // Retrieve environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    parseTime := os.Getenv("DB_PARSE_TIME")

    // Construct the connection string
    dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=%s", dbUser, dbPassword, dbHost, dbPort, dbName, parseTime)

    // Connect to the database
    db, err := sql.Open("mysql", dsn)
    checkErr(err)
    defer db.Close()

    // Check if the connection is alive
    checkErr(db.Ping())

    // Create a new table if it doesn't exist
    query := `
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT,
            username TEXT NOT NULL,
            password TEXT NOT NULL,
            created_at DATETIME,
            PRIMARY KEY (id)
        );`
	result, err := 	db.Exec(query)
    checkErr(err)

    // Insert a new user
    username := "johndoe"
    password := "secret"
    createdAt := time.Now()

    result, err = db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
    checkErr(err)

    id, err := result.LastInsertId()
    checkErr(err)
    fmt.Println("Inserted user ID:", id)

    // Fetch the last inserted user
    var lastUser User
    err = db.QueryRow("SELECT id, username, password, created_at FROM users ORDER BY id DESC LIMIT 1").Scan(&lastUser.ID, &lastUser.Username, &lastUser.Password, &lastUser.CreatedAt)
    checkErr(err)

    fmt.Printf("Last inserted user: %d %s %s %s\n", lastUser.ID, lastUser.Username, lastUser.Password, lastUser.CreatedAt)

    // Get total users count
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
    checkErr(err)

    fmt.Printf("Total users: %d\n", count)

    // Verify the contents of the users table
    rows, err := db.Query("SELECT id, username, password, created_at FROM users")
    checkErr(err)
    defer rows.Close()

    for rows.Next() {
        var u User
        checkErr(rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt))
        fmt.Printf("User: %d %s %s %s\n", u.ID, u.Username, u.Password, u.CreatedAt)
    }
    checkErr(rows.Err())
}
