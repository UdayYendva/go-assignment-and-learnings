package main

import (
	"fmt"

	"github.com/UdayYendva/go-assignment-and-learnings/student"
)

func main() {
	a := student.NewCreatedStudent(1)
	fmt.Println("Student with id is created ")
	student.IgniteFunction(a)

}
