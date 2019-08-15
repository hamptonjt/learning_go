package asucommon

import (
	"fmt"
	"database/sql"
	"log"
	// importing goracle - not by name
	_ "gopkg.in/goracle.v2"
)

// JobSubParam stucture hold the parameter number and value stored in GJBPRUN
type JobSubParam struct {
	ParmNum string
	ParmVal string
}

// OpenConnection returns a connection to the database using the passed in credentials
func OpenConnection(oraSid, username, passwd string) *sql.DB {
	connectionStr := ""

	if username[:1] == "[" {
		connectionStr = fmt.Sprintf("%s@//%s", username, passwd[1:])
	} else {
		connectionStr = fmt.Sprintf("%s/%s@%s", username, passwd, oraSid)
	}

	conn, err := sql.Open("goracle", connectionStr)
	if err != nil {
		fmt.Println(err)
	}
	return conn
}

// CheckRole validates the user has access to the given Banner Object/Job
func CheckRole(conn *sql.DB, obj string) bool {
	secure := false
	role := ""

	_, err := conn.Exec("begin :res := get_banner_role(:obj, :version, :seed1, :seed3); end;", sql.Out{Dest: &role}, obj, "", "12345678", "87651234")
	if err != nil {
		fmt.Println("Error executing 'get_banner_role' function")
		fmt.Println(err)
		return secure
	}

	if role == "INSECURED" {
		secure = false
	} else {
		secure = true
	}

	if secure {
		_, err := conn.Exec("begin dbms_session.set_role(:role); end;", role)
		if err != nil {
			fmt.Println("Error setting role for job")
			fmt.Println(err)
		}
	}

	return secure
}

// GetJobSubParams returns an array of job Number/Value pairs
func GetJobSubParams (conn *sql.DB, jobName, oneUp string) []JobSubParam {
	var params []JobSubParam

	rows, err := conn.Query(`
		select distinct gjbprun_number, gjbprun_value
		from gjbprun
		where gjbprun_job = :job
		  and gjbprun_one_up_no = :oneup
		order by gjbprun_number
	`, jobName, oneUp)
	if err != nil {
		fmt.Println("Error fetching job submission parameters")
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var num, val string
		err := rows.Scan(&num, &val)
		if err != nil {
			log.Fatal(err)
		}
		params = append(params, JobSubParam{ParmNum: num, ParmVal: val})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return params
}