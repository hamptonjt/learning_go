package azgaorp

import (
	"fmt"
	"github.com/hamptonjt/asucommon"
	"database/sql"
	"bufio"
	"io/ioutil"
	"os"
	"time"
)

func getData() {

}

func main() {
	username = os.Args[1]
	passwd = os.Args[2]
	oneUp = os.Args[3]
	oraSID = os.Args[4]
	banHome = os.Args[5]
	homeDir = os.Args[6]

	fmt.Println("Started Execution at %s", time.Now().Format(time.RFC3339))

	// Now, process this application

	fmt.Println("Execution Completed at %s", time.Now().Format(time.RFC3339))
}