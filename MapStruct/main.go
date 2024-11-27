package main

import (
	"fmt"
	"reflect"
)

//	type Student struct {
//		number  int
//		name    string
//		gender  bool
//		subject []string
//	}
type People struct {
	name string
	age  int
}
type Student struct {
	People
	classOfStudent string
	number         int `max:"100"` //tag
}

func main() {
	studentJake := Student{
		People: People{
			name: "Loc",
			age:  12,
		},
		classOfStudent: "12A2",
		number:         99,
	}
	t := reflect.TypeOf(studentJake)
	field, _ := t.FieldByName("number") // Use double quotes for field name
	// fmt.Println("Student Struct:", studentJake)
	fmt.Printf("Field: %v \n", field)
	fmt.Println("Tag:", field.Tag.Get("max"))
	// Student := struct{ name string }{name: "Loc"}
	// copyStudent := &Student
	// copyStudent.name = "Jake"
	// fmt.Println(Student)
	// fmt.Println(copyStudent)
	//::::::::::::STRUCT:::
	// studentYUH := Student{}
	// studentJake := Student{
	// 	number: 1,
	// 	name:   "Huu Loc",
	// 	gender: true,
	// 	subject: []string{
	// 		"Math",
	// 		"English",
	// 		"Computer",
	// 	},
	// }
	// // fmt.Printf()
	// // fmt.Printf()
	// fmt.Println(studentJake)
	// fmt.Println(studentJake.name)

	// Student := struct{ name string }{name: "Loc"}
	// copyStudent := &Student
	// copyStudent.name = "Jake"
	// fmt.Println(Student)
	// fmt.Println(copyStudent)

	//
	// ::::::::::::::MAP::::::::::::::::::::
	// studentNameAgeMap := make(map[string]int)
	// // fmt.Println(studentNameAgeMap)
	// studentNameAgeMap = map[string]int{
	// 	"Jake":  19,
	// 	"Tom":   35,
	// 	"Jerry": 40,
	// } // not like order init
	// // m := map[[3]int]int{} // can not with slice for key
	// // fmt.Println(m)
	// // fmt.Println(studentNameAgeMap)
	// // studentNameAgeMap["Micheal"] = 25
	// // studentNameAgeMap["Jake"] = 1999

	// copyMap := studentNameAgeMap
	// fmt.Println(studentNameAgeMap)
	// copyMap["Jake"] = 1999
	// fmt.Println(studentNameAgeMap)
	// fmt.Println(copyMap)
	// delete(studentNameAgeMap, "Micheal")
	// _, isExist := studentNameAgeMap["Micheal"]
	// fmt.Println(isExist)
	// fmt.Println(len(studentNameAgeMap))
	// fmt.Println(studentNameAgeMap)
	// fmt.Println("ageMike", ageOfMike)
	// fmt.Println(studentNameAgeMap["Jake"])
	// fmt.Println(studentNameAgeMap["Tom"])
}
