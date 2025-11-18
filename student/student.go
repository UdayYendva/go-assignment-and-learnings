package student

import "fmt"

type Student struct {
	Id int
}
type Interface interface {
	GetStudentId() int
	ChangeStudentIdWithoutPointers()
	ChangeStudentIdWithPointer()
}

func NewCreatedStudent(id int) *Student { // int page *int , page =&age and *page =42 means the value of age is changeto 42
	return &Student{Id: id}
}

func IgniteFunction(m Interface) {
	fmt.Println("The initial value of the student id is :", m.GetStudentId()) // here it will print the original in which the the constructor change the Student id while creating a new student

	m.ChangeStudentIdWithoutPointers()
	fmt.Println("The value without using pointers:", m.GetStudentId())

	m.ChangeStudentIdWithPointer()
	fmt.Println("The value with using pointers:", m.GetStudentId())

}

func (a *Student) GetStudentId() int {
	return a.Id
}

func (s *Student) ChangeStudentIdWithPointer() { // here it will change the value of Student (id) usimg pointer without creating a copy of it
	s.Id = 11
}

func (s Student) ChangeStudentIdWithoutPointers() { // here it will create a copy of the stored value Stduent Id instead of manipulationg the value
	s.Id = 10
}
