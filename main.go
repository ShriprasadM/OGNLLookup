package main

import (
	"OGNLLookup/examples"
	"fmt"
)

func main() {

	student := examples.Student{
		Id: 2,
		Subjects: []examples.Subject{
			{Id: 1, Name: "Maths"},
			{Id: 2, Name: "Science"},
			{Id: 3, Name: "English"},
		},
	}

	subjectName, err := LookUpStudent("student.subjects.name", student, StudentIndexInfo{
		SubjectsIndex: 2,
	}, "")

	if err != nil {
		fmt.Println(err)
	}

	println(subjectName)
}
