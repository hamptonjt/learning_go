package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "gopkg.in/goracle.v2"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Invalid credentials/instance supplied")
		return
	}
	username := os.Args[1]
	pass := os.Args[2]
	instance := os.Args[3]

	connStr := fmt.Sprintf("%s/%s@%s", username, pass, instance)

	db, err := sql.Open("goracle", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Get User input for a Banner ID to return their name:
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Enter a Banner ID: ")
	text, _ := reader.ReadString('\n')

	// Need the \r for windows.. *nix only needs to replace the \n
	text = strings.Replace(text, "\r\n", "", -1)

	query := fmt.Sprintf(`
		select spriden_first_name || ' ' || spriden_last_name 
		from spriden 
		where spriden_change_ind is null
		  and spriden_id = '%s'
	`, text)

	// Fetch a record from Banner.
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		rows.Scan(&name)
	}
	fmt.Printf("The person who matches the given ID: %s\n", name)

}
