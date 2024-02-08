/* You're tasked with building a student information system 
in GoLang for a school. Each student record needs to store 
details such as student ID, name, age, and grade. Define 
variables to store the information of a single student and 
assign values accordingly. Pay attention to selecting appropriate
data types to represent each piece of information.*/

package main

import(
	"fmt"
)


// Define a struct to represent a student
type Student struct {
    ID    int
    Name  string
    Age   int
    Grade string
}

func main() {
    // Create a new student record
	var student Student
    fmt.Println("Enter student information:")

    // Take input for student ID
    fmt.Print("ID: ")
    fmt.Scanln(&student.ID)

    // Take input for student name
    fmt.Print("Name: ")
    fmt.Scanln(&student.Name)

    // Take input for student age
    fmt.Print("Age: ")
    fmt.Scanln(&student.Age)

    // Take input for student grade
    fmt.Print("Grade: ")
    fmt.Scanln(&student.Grade)

    // Print out the student information
    fmt.Println("\nStudent Information:")
    fmt.Println("ID:", student.ID)
    fmt.Println("Name:", student.Name)
    fmt.Println("Age:", student.Age)
    fmt.Println("Grade:", student.Grade)
}
