package examples

type Student struct {
	Id       int
	Subjects []Subject
}

type Subject struct {
	Id   int
	Name string
}
