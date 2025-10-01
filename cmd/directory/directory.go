package main

import (
	"fmt"

	"github.com/mazimkhan/golang_example/pkg/directory"
)

func main() {
	var employees = map[int]directory.Employee{}
	employees = directory.AddEmployee(employees, directory.Employee{ID: 1, Name: "John Doe", Phone: "123-456-7890"})
	employees = directory.AddEmployee(employees, directory.Employee{ID: 2, Name: "Jane Smith", Phone: "456-654-3210"})
	employees = directory.AddEmployee(employees, directory.Employee{ID: 3, Name: "Alice Johnson", Phone: "123-555-5555"})
	employees = directory.AddEmployee(employees, directory.Employee{ID: 4, Name: "Bob Brown", Phone: "456-444-4444"})
	employees = directory.AddEmployee(employees, directory.Employee{ID: 5, Name: "Charlie Davis", Phone: "456-333-3333"})
	println("--------------------------------------------")
	println("Employees added:")
	for _, e := range employees {
		fmt.Printf("ID=%d, Name=%s, Phone=%s\n", e.ID, e.Name, e.Phone)
	}

	var result = directory.SearchEmployees(employees, func(e directory.Employee) bool {
		return e.Name[0] == 'J'
	})
	println("--------------------------------------------")
	println("Employees whose names start with 'J':")
	for _, e := range result {
		fmt.Printf("ID=%d, Name=%s, Phone=%s\n", e.ID, e.Name, e.Phone)
	}
	result = directory.SearchEmployees(employees, func(e directory.Employee) bool {
		return e.Phone[:3] == "123"
	})
	println("--------------------------------------------")
	println("Employees whose phone numbers start with '123':")
	for _, e := range result {
		fmt.Printf("ID=%d, Name=%s, Phone=%s\n", e.ID, e.Name, e.Phone)
	}
	println("--------------------------------------------")
	println("Remove employees whose phone numbers start with '123':")
	for _, e := range result {
		employees = directory.RemoveEmployee(employees, e.ID)
	}
	println("--------------------------------------------")
	println("Employees after removal:")
	for _, e := range employees {
		fmt.Printf("ID=%d, Name=%s, Phone=%s\n", e.ID, e.Name, e.Phone)
	}
}
