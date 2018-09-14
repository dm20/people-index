// to be used in "people_index.go" to test the functionalities of the service
// run from main

func runTests() {
  addNewPerson("1","Oliver")
  addChildToKey("1","A similar person to Oliver, Bill")
  addNewPerson("5","Eric")
  addChildToKey("5","A similar person to Eric, Jane")

  addNewPerson("2","John")
  addNewPerson("3","Amy")
  addNewPerson("4","Reece")
  addNewPerson("6","Allice")
  addNewPerson("7","Bob")
  addNewPerson("8","Joanne")
  addNewPerson("9","Dani")

  deletePerson("9")
  deletePerson("8")
  deletePerson("7")

  name1 := listValueForKey("1")
  name2 := listValueForKey("2")
  name3 := listValueForKey("3")
  name4 := listValueForKey("4")
  name5 := listValueForKey("5")
  name6 := listValueForKey("6")
  fmt.Println(name1)
  fmt.Println(name2)
  fmt.Println(name3)
  fmt.Println(name4)
  fmt.Println(name5)
  fmt.Println(name6)
}
