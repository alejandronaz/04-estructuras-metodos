package main

import "fmt"

type Person struct {
	ID          string
	Name        string
	DateOfBirth string
}

type Employee struct {
	// ID string
	Position string
	Person   Person
}

func (emp Employee) PrintEmployee() {
	fmt.Println("Empleado --> ID: ", emp.Person.ID, " - Name: ", emp.Person.Name, " - Date of birth: ", emp.Person.DateOfBirth, " - Position: ", emp.Position)
}

func main() {
	p1 := Person{"1", "Juan", "01/01/1990"}
	emp1 := Employee{"Developer", p1}
	emp1.PrintEmployee()
}
