package main

import "github.com/dm20/people_index"

func main() {
  people_index.Initialize()   // open 'people.json' and get existing file data, or create it if needed
  people_index.RunTests()     // see RunTests() in people_index.go
  people_index.SaveAndExit()  // replace old json data with updated json data in 'people.json'
}

// 1) make sure to put the interface in $GOPATH : 
//       go get github.com/dm20/people_index
//
//    and to get the dependencies: 
//        go get github.com/Jeffail/gabs
//
// 2) go build example_session.go
// 3) go run example_session.go
