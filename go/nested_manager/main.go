// Mastery (Logic & Architecture)
// 9. The Nested Manager
// Create a Department struct that contains a slice of Employee structs.
// Write a method on Department called TotalPayroll() that iterates
// through the employees and sums their salaries.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary int64
}

type Department struct {
	Name      string
	Employees []Employee
}

// TotalPayroll now returns an int64 so we can use the result elsewhere.
func (d Department) TotalPayroll() int64 {
	var total int64

	for _, e := range d.Employees {
		total += e.Salary
	}

	return total
}

func main() {
	// 1. Create some employees
	e1 := Employee{Name: "Alice", Salary: 95000}
	e2 := Employee{Name: "Bob", Salary: 82000}
	e3 := Employee{Name: "Charlie", Salary: 105000}

	// 2. Initialize the Department with the employees
	dept := Department{
		Name:      "Engineering",
		Employees: []Employee{e1, e2, e3},
	}

	// 3. Calculate and display the results
	payroll := dept.TotalPayroll()

	fmt.Printf("Department: %s\n", dept.Name)
	fmt.Printf("Total Payroll: $%d\n", payroll)
}
