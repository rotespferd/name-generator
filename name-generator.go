package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

// define CLI flags
//TODO: add default path for default file
var filepath = flag.String("filepath", "./names.json", "The path to the file with the names")
var removeName = flag.Bool("remove", true, "if true the name will be removed from list")

// Names is an array of names
type Names struct {
	Description string
	Names       []Name
}

// Name is a name of the tool
type Name struct {
	Name    string
	Removed bool
}

func main() {
	flag.Parse()

	names := loadFile(*filepath)

	log.Println("Loaded names with description ", names.Description)
}

func loadFile(filepath string) Names {
	log.Println("Load files from ", filepath)

	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Println("file could not be opened")
		log.Fatal(err)
	}

	var names Names

	err = json.Unmarshal(file, &names)
	if err != nil {
		log.Fatal(err)
	}

	return names
}
