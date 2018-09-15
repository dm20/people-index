# People Index
A simple Go service for creating and modifying key-value pairs in a persistent ledger.

### Features:
  1) Create and save JSON files locally if one does not yet exist
  2) Edit existing or new JSON files by adding 'people' tags  
     - A person tag can be any key and associated value
  3) Look up person information by key
  4) Delete person information associated with a given key  
     - Remove "hanging keys" and whitespace from the JSON file itself
  5) Add multiple values to one key
  6) Include the test() function in people_index.go to demo the functionality.
     - creates a file equivalent to example.json

### To run the service:
    go get go get github.com/Jeffail/gabs
    go build people_index.go
    go run people_index.go

Uses the <a href='https://github.com/Jeffail/gabs'>gabs</a> JSON library.    
