# People Index
A simple Go service for creating and modifying key-value pairs in a persistent JSON ledger.  
#### The service is demonstrated using people's names as values and integer IDs as keys

### Features:
  1) Create and save JSON files locally if one does not yet exist
  2) Edit existing or new JSON files by adding 'people' tags  
     - A person tag can be any key and associated value
     - Add multiple values to one key if desired
  3) Look up person information by key
  4) Delete person information associated with a given key  
     - Remove "hanging keys" and whitespace from the JSON file itself
  5) Include the test() function in people_index.go to demo functionality
     - Creates a file equivalent to example.json called people.json

### To run the example session*
    go get github.com/dm20/people_index    
    go build example_session.go
    go run example_session.go
    
example session code:
    
    package main

    import "github.com/dm20/people_index"

    func main() {
      people_index.Initialize()
      people_index.RunTests()
    }

*Note dependencies: <a href='https://github.com/Jeffail/gabs'>gabs</a>  

    go get github.com/Jeffail/gabs
