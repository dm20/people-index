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
// 2) call runTests() from main 

package main

import (
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  "github.com/Jeffail/gabs"
)

var existing_json *gabs.Container
var num_deletions int

func main() {
  initialize()
  // call runTests() here if desired
  addNewPerson("a key","a value")
  exit()
}

// create a new JSON object for each person that is created
func addNewPerson(key string, value string) {
  existing_json.Set(value,"people",key)
}

// add another value to the given key (creates an array)
func addChildToKey(key string, value string) {
  newChild := gabs.New()
  newChild.Set(value,"people",key)
  existing_json.Merge(newChild);
}

// list the associated value (single person or array) for a given key
func listValueForKey(key string) string {
	return existing_json.Path("people." + key).String()
}

// delete a person from the JSON file (key can still be used)
// see next function also
func deletePerson(key string) {
  num_deletions++
  existing_json.Set("nil","people",key)
}

// delete any lines in the JSON file marked as deleted
// called whenever a session ends
func clearHangingKeys(input string) string {
  if (num_deletions == 0) { return input }
  lines := strings.Split(string(input), "\n")
  newLines := make([]string, len(lines) - num_deletions)
  j := 0
  for _, line := range lines {
    if (!strings.Contains(line, "nil")) { 
      newLines[j] = line
      j++
    }
  num_deletions = 0 // reset if more than one call per session
  return strings.Join(newLines, "\n")
}

// Update and close the file. 
func exit() {
  str := existing_json.StringIndent("", "  ")
  newFile := clearHangingKeys(str) // Any entries deleted in this session are removed
  f, _ := os.Create("./people.json")
  f.Sync()
  f.WriteString(newFile);
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
