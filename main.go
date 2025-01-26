package main
import (
	"github.com/aljarida/formatore/dblogic"
	"fmt"
)

func main() {
	fmt.Println("Testing sqllogic.CheckConnection.")
	dblogic.CheckConnection()
	fmt.Println("Done checking connection.")
}
