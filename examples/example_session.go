package main

import "github.com/dm20/people_index"

// 1) make sure to put the interface in $GOPATH : 
//       go get github.com/dm20/people_index
//
//    and to get the dependencies: 
//        go get github.com/Jeffail/gabs
//
// 2) go build example_session.go
// 3) go run example_session.go

func main() {
  people_index.Initialize() // open people.json and get existing file data, or create it if needed
  people_index.RunTests()   // add new person tags, retrieve them, delete them
  people_index.Exit().      // save changes that were made to data in people.json
}
