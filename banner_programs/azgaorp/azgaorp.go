package main

import (
	"database/sql"
	"fmt"

	asucommon "github.com/hamptonjt/learning_go/banner_programs/asucommon"

	// "bufio"
	// "io/ioutil"
	"encoding/csv"
	"os"
	"time"
)

func getData(conn *sql.DB, homeDir, oneUp string) {
	csvFile, err := os.Create(fmt.Sprintf("%s/azgaorp_%s.csv", homeDir, oneUp))
	if err != nil {
		fmt.Println("Error creating csv file")
		fmt.Println(err)
	}

	lisFile, err := os.Create(fmt.Sprintf("%s/azgaorp_%s.lis", homeDir, oneUp))
	if err != nil {
		fmt.Println("Error creating lis file")
		fmt.Println(err)
	}

	csvWriter := csv.NewWriter(csvFile)
	lisWriter := csv.NewWriter(lisFile)
	header := []string{"ID", "FirstName", "LastName", "MI", "Email", "Role"}
	csvWriter.Write(header)
	lisWriter.Write(header)
	fmt.Println(header)

	rows, err := conn.Query(`
		select spriden_id, spriden_first_name, spriden_last_name, spriden_mi, goremal_email_address, twgrrole_role
           from spriden, goremal, twgrrole 
          where spriden_pidm = goremal_pidm 
            and spriden_pidm = twgrrole_pidm 
            and twgrrole_role in ('MOVESMANAGER', 'DEVELOPMENTOFFICER', 'ADVANCEMENTQUERY') 
            and goremal_preferred_ind = 'Y'
            and goremal_emal_code = 'ASUE'
            and goremal_status_ind = 'A'
            and spriden_change_ind is null
	`)
	if err != nil {
		fmt.Println("Error fetching data")
		fmt.Println(err)
	}

	for rows.Next() {
		var id, fname, lname, mi, email, role string
		// var values []string
		// rows.Scan(&values)
		rows.Scan(&id, &fname, &lname, &mi, &email, &role)
		values := []string {id, fname, lname, mi, email, role}
		csvWriter.Write(values)
		csvWriter.Flush()
		lisWriter.Write(values)
		lisWriter.Flush()
		fmt.Println(values)
	}

}

func main() {
	username := os.Args[1]
	passwd := os.Args[2]
	oneUp := os.Args[3]
	oraSID := os.Args[4]
	// banHome := os.Args[5]
	homeDir := os.Args[6]

	fmt.Printf("Started Execution at %s\n", time.Now().Format(time.RFC3339))

	// Now, process this application
	conn := asucommon.OpenConnection(oraSID, username, passwd)
	if asucommon.CheckRole(conn, "AZGAORP") {
		getData(conn, homeDir, oneUp)
	}

	fmt.Printf("Execution Completed at %s\n", time.Now().Format(time.RFC3339))
}
