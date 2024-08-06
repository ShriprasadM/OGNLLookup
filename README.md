# OGNLLookup

# _Retrieve value from your Go Objects based on OGNL expression_
**__Influenced from [OGNL](https://commons.apache.org/dormant/commons-ognl/index.html)__**

# Example
1. ``` git clone --single-branch --branch examples_readme git@github.com:ShriprasadM/OGNLLookup.git ```
2. ``` cd OGNLLookup ```
> [!NOTE]
> Phase I - Generate OGNL Source
3. ``` make ognl_generate ```
4. This will generate `student_ognl.go` next to `main.go`  
5. Open `student_ognl.go`. Add following statements  `package` line in `student_ognl.go`
```go
import (
	"OGNLLookup/examples"
	"fmt"
)
```
> [!NOTE]
> Phase II - Determine value using OGNL Lookup

6. Now OGNL Library is ready for demo
8. Open main.go and see all compilation errors are resolved
9. ``` make demo ```
10. This will show `English` as output. see main.go for more details on example

# Example Code
```go
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
```
