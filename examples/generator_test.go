package examples

import (
	"OGNLLookup/ognl"
	"testing"
)

func TestGenerateStudentOGNL(t *testing.T) {
	ognl.GenerateSource(Student{}, "student", "../student_ognl.go", "main")
}
