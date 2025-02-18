package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
    "context"
    "github.com/jackc/pgx/v5"
)

var host string;
var p string;
var user string;
var pass string;
var database string;

func libPqGetSchema(port int) {
    conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, database)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT "datname" FROM "pg_database" WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if err != nil {
            fmt.Println("Scan", err)
            return
        }
        fmt.Println("*", name)
    }
    defer rows.Close()

    // Get all tables from __current__ database
    query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name`
    rows, err = db.Query(query)
    if err != nil {
        fmt.Println("Query", err)
        return
    }

    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if err != nil {
            fmt.Println("Scan", err)
            return
        }
        fmt.Println("+T", name)
    }
    defer rows.Close()
}

// exercise 2: Rewrite the getSchema.go utility so that it works with the jackc/pgx package
func jackcPgxGetSchema(port int) {
    connUrlString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pass, host, port, database)
    conn, err := pgx.Connect(context.Background(), connUrlString)
    if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

    rows, err := conn.Query(context.Background(), `SELECT "datname" FROM "pg_database" WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if err != nil {
            fmt.Println("Scan", err)
            return
        }
        fmt.Println("*", name)
    }
    defer rows.Close()

    // Get all tables from __current__ database
    query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name`
    rows, err = conn.Query(context.Background(), query)
    if err != nil {
        fmt.Println("Query", err)
        return
    }

    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if err != nil {
            fmt.Println("Scan", err)
            return
        }
        fmt.Println("+T", name)
    }
    defer rows.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname port username password db")
		return
	}

	host = arguments[1]
	p = arguments[2]
	user = arguments[3]
	pass = arguments[4]
	database = arguments[5]

	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Not a valid port number:", err)
		return
	}
    fmt.Println("libpq:")
	libPqGetSchema(port)
    fmt.Println("\npgx:")
    jackcPgxGetSchema(port)
}
