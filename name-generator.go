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
var shouldNameBeRemoved = flag.Bool("remove", true, "if true the name will be removed from list")

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

	nextName, names := nextName(names)

	json, _ := json.Marshal(names)

	ioutil.WriteFile(*filepath, json, 0777)

	log.Println("==================================================")
	log.Println("The next name is:")
	log.Println(nextName)
	log.Println("==================================================")

	log.Println(names)
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

func nextName(names Names) (string, Names) {
	for index, name := range names.Names {
		if !name.Removed {
			if *shouldNameBeRemoved {
				names.Names[index].Removed = true
			} else {
				log.Println("The name will not be removed!")
			}
			return name.Name, names
		}
	}
	return "", names
}
