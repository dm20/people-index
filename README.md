## A simple Go service for creating and modifying key-value pairs in a persistent JSON file.    

## Features:
 - Create and save JSON files locally if one does not yet exist
 - Edit existing or new JSON files by adding 'people' tags  
     - A person tag can be any key and associated value
     - Add multiple values to one key if desired
 - Look up person information by key
 - Delete person information associated with a given key  
     - Remove "hanging keys" and whitespace from the JSON file itself
 - Include the test() function in people_index.go to demo functionality
     - Creates a file equivalent to example.json called people.json
     
## To run an example session:  
1) get people_index  
```shell
    go get github.com/dm20/people_index
```    
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
3) get external dependencies (one time only) and run  
```shell
    go get ./examples/example_session.go
    go run example_session.go
```

which creates a JSON file named 'people.json':
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
[<img src='https://static.allcloud.com/assets/images/blog/golang.png' width = '90' height = '50'>](https://golang.org)[<img src='https://upload.wikimedia.org/wikipedia/commons/thumb/c/c9/JSON_vector_logo.svg/1000px-JSON_vector_logo.svg.png' width = '45' height ='45'>](https://www.json.org)
