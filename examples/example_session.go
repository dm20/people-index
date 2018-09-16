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
  people_index.Initialize()
  people_index.RunTests()
}
