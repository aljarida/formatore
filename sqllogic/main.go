package sqllogic

import(
	"database/sql"
	"fmt"
	"log"
	"github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERISON()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}
}
