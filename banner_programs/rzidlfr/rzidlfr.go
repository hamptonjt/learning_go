package main

import (
	"database/sql"
	"fmt"

	asucommon "github.com/hamptonjt/learning_go/banner_programs/asucommon"

	// "encoding/csv"
	"os"
	"time"
)

func runProcess(conn *sql.DB, homeDir, oneUp string, params []asucommon.JobSubParam) {
	aidy, term := "", ""
	ttlUndergrad, ttlParent, ttlGrad, grandTotal := 0.00, 0.00, 0.00, 0.00
	lisFile, err := os.Create(fmt.Sprintf("%s/rzidlfr_%s.lis", homeDir, oneUp))
	if err != nil {
		fmt.Println("Error creating lis file")
		fmt.Println(err)
	}

	for i := 0; i < len(params); i++ {
		param := params[i]
		if param.ParmNum == "01" {
			aidy = param.ParmVal
		}
		if param.ParmNum == "02" {
			term = param.ParmVal
		}
	}

	_, err = conn.Exec("begin baninst1.rzidlfr_calc(:1, :2, :3, :4, :5, :6); end;", aidy, term,
		sql.Out{Dest: &ttlUndergrad}, sql.Out{Dest: &ttlParent}, sql.Out{Dest: &ttlGrad}, sql.Out{Dest: &grandTotal})

	if err != nil {
		fmt.Println("Error calling rzidlfr_calc")
		fmt.Println(err)
	}

	fmt.Println("")
	fmt.Printf("Total Undergraduate: %.2f\n", ttlUndergrad)
	fmt.Printf("Total Graduate:      %.2f\n", ttlGrad)
	fmt.Printf("Total Parent:        %.2f\n", ttlParent)
	fmt.Println("")
	fmt.Printf("Grand Total:         %.2f\n", grandTotal)

	lisFile.WriteString(fmt.Sprintf("Total Undergraduate: %.2f\n", ttlUndergrad))
	lisFile.WriteString(fmt.Sprintf("Total Graduate:      %.2f\n", ttlGrad))
	lisFile.WriteString(fmt.Sprintf("Total Parent:        %.2f\n", ttlParent))
	lisFile.WriteString("")
	lisFile.WriteString(fmt.Sprintf("Grand Total:         %.2f\n", grandTotal))

}

func main() {
	// jobName := os.Args[0]
	username := os.Args[1]
	passwd := os.Args[2]
	oneUp := os.Args[3]
	oraSID := os.Args[4]
	// banHome := os.Args[5]
	homeDir := os.Args[6]

	fmt.Printf("Started Execution at %s\n", time.Now().Format(time.RFC3339))

	// Now, process this application
	conn := asucommon.OpenConnection(oraSID, username, passwd)
	if asucommon.CheckRole(conn, "RZIDLFR") {
		params := asucommon.GetJobSubParams(conn, "RZIDLFR", oneUp)

		runProcess(conn, homeDir, oneUp, params)
	}

	fmt.Printf("Execution Completed at %s\n", time.Now().Format(time.RFC3339))
}
