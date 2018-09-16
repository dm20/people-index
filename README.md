# People Index
A simple Go service for creating and modifying key-value pairs in a persistent JSON ledger.  
The service is demonstrated using people's names as values and integer IDs as keys

## To run an example session:
    go get github.com/dm20/people_index    

```go
package main

import "github.com/dm20/people_index"

func main() {
  people_index.Initialize() // open 'people.themjson' and get existing file data, or create it if needed
  people_index.RunTests()   // add new person tags, retrieve them, delete them. See RunTests() in people_index.go
  people_index.Exit().      // save changes that were made to data in people.json
}
```
generates:
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
       
         
           
##### *Note dependencies: <a href='https://github.com/Jeffail/gabs'>gabs</a>  
    go get github.com/Jeffail/gabs

