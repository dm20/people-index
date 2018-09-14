/*
* AUTHOR: Daniel McGrath
* DATE: 9/13/18
* DESCRIPTION:
* This service tracks "people" information by creating a JSON file called 
* "people.json" in the location that the service is stored, if one does not already exist, 
* and can add new people as well as list people info, and delete people info.
*/

// To run the service
// 1) go get github.com/Jeffail/gabs
// 2) go build people_index.go
// 3) go run people_index.go

// To test
// 1) paste the function test.go to the bottom of this file
// 2) call test() from main 

package main

import (
  "fmt"
	"os"
  "io/ioutil"
	"github.com/Jeffail/gabs"
)

var existing_json *gabs.Container

func main() {

  initialize()

  // call test() here if desired

  addNewPerson("a key","a value")

  exit()
}

// create a new JSON object for each person that is created
func addNewPerson(key string, value string) {
  existing_json.Set(value,"people",key)
}

// add another value to the given key (creates an array)
func addChildToKey(key string, value string) {
  newCan := gabs.New()
  newCan.Set(value,"people",key)
  existing_json.Merge(newCan);
}

// delete a person from the JSON file (key can still be used)
func deletePerson(key string) {
  existing_json.Set("","people",key)
}

// list the associated value (single person or array) for a given key
func listValueForKey(key string) string {
	return existing_json.Path("people." + key).String()
}

// Update and close the file
func exit() {
  str := existing_json.StringIndent("", "  ")
  f, _ := os.Create("./people.json")
  f.Sync()
  f.WriteString(str);
  defer f.Close()
}

// If no file exists, create it, otherwise open it in writable format
func initialize() {
  f, err := os.OpenFile("./people.json",os.O_APPEND | os.O_WRONLY, 0600)
  if (err != nil) {
    fmt.Println("\n\nCreating 'people.json'...\n\n")
    f, _ = os.Create("./people.json")
    jsonObj, _ := gabs.ParseJSON([]byte(`{"people":{}}`))
    newFile := jsonObj.StringIndent("", "  ")
    f.WriteString(newFile)
    fmt.Println("\n\nDone.\n\n")
  }

  existing_data, _ := ioutil.ReadFile("./people.json")
  existing_json, _ = gabs.ParseJSON(existing_data)
}
