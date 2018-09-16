# People Index
## A simple Go service for creating and modifying key-value pairs in a persistent JSON ledger.    

The service is demonstrated using people's names as values and integer IDs as keys

## Features:
  1) Create and save JSON files locally if one does not yet exist
  2) Edit existing or new JSON files by adding 'people' tags  
     - A person tag can be any key and associated value
     - Add multiple values to one key if desired
  3) Look up person information by key
  4) Delete person information associated with a given key  
     - Remove "hanging keys" and whitespace from the JSON file itself
  5) Include the test() function in people_index.go to demo functionality
     - Creates a file equivalent to example.json called people.json

## To run an example session:
1) get people_index

    go get github.com/dm20/people_index
    
2) create 'example_session.go':
```go
package main

import "github.com/dm20/people_index"

func main() {
  people_index.Initialize()   // open 'people.json' and get existing file data, or create it if needed
  people_index.RunTests()     // see RunTests() in people_index.go
  people_index.SaveAndExit()  // replace old json data with updated json data in 'people.json'
}
```
3) get external dependencies and run

    go get ./examples/example_session.go
    go run example_session.go

which creates a file that looks like this:
```json
{
  "people": {
    "1": [
      "Oliver",
      "A similar person to Oliver, Bill"
    ],
    "2": "John",
    "4": "Reece",
    "5": [
      "Eric",
      "A similar person to Eric, Jane"
    ],
    "6": "Allice",
  }
}
```
