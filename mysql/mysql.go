package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Replace these connection parameters with your MySQL server details.
	dsn := "root:163766@tcp(yuhan.bestzyx.com:3306)/gorm"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening the database:", err)
		return
	}
	defer db.Close()
	tableName := "test"
	query := fmt.Sprintf("TRUNCATE TABLE %s", tableName)

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Error truncating the table:", err)
		return
	}

	fmt.Printf("Table %s truncated.\n", tableName)
}
