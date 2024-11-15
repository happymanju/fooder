package cli

import (
	"flag"
	"log"
)

// CLI parses command line flags and executes app with arguments
func CLI(args []string) int {
	listFlag := flag.Bool("l", false, "list all records")
	newFlag := flag.Bool("n", false, "add a new food item")
	updateFlag := flag.Int("u", -1, "update an existing food record")
	deleteFlag := flag.Int("d", -1, "delete an existing record")

	flag.Parse()

	if *listFlag == true {
		app.List()
		return 0
	} else if *newFlag == true {
		err := app.New(args)
		if err != nil {
			return 0
		} else {
			log.Printf("error: %v", err)
			return 1
		}
	} else if *updateFlag != -1 {
		err := app.Update(*updateFlag)
		if err != nil {
			log.Printf("error: %v", err)
			return 1
		} else {
			return 0
		}
	} else if *deleteFlag != -1 {
		err := app.Delete(*deleteFlag)
		if err != nil {
			log.Printf("error: %v", err)
			return 1
		} else {
			return 0
		}
	}
	return 0
}
