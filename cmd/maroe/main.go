package main

import (
	"log"
	"os"

	"github.com/tjamet/maroe/pkg/copy"

	"github.com/docopt/docopt-go"
)

const help = `Usage: importer --input=<Folder> --output=<Folder>

Options:
	--input=<Folder>, -i    The source folder (typically the path to the card) to be synchronised
	--output=<Folder>, -o   The destination folder where photos will be named uniquely
`

func main() {
	args, err := docopt.Parse(help, os.Args[1:], true, "0.0.0", false, true)
	if err != nil {
		log.Fatal(err)
	}
	err = copy.Copy(args["--output"].(string), args["--input"].(string))
	if err != nil {
		log.Fatalf("completed with errors")
	}
}
