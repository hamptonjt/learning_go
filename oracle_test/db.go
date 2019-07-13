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
	id, _ := reader.ReadString('\n')

	// Need the \r for windows.. *nix only needs to replace the \n
	id = strings.Replace(id, "\r\n", "", -1)

	query := `
		select spriden_first_name || ' ' || spriden_last_name 
		from spriden 
		where spriden_change_ind is null
		  and spriden_id = :id
	`

	// Fetch a record from Banner.
	row := db.QueryRow(query, id)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}

	var name string
	row.Scan(&name)

	if name != "" {
		fmt.Printf("The person who matches the given ID: %s\n", name)
	} else {
		fmt.Printf("No person found for the given id: %s", id)
	}

}
