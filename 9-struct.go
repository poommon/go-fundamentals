package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("================Main Struct=====================")
	main_struct()

	fmt.Println("================Pointer Struct=====================")
	pointer_struct()

	fmt.Println("================Json Struct=====================")
	json_struct()
}

// -----------------------------------
// 1. Capitalize fields so the JSON package can access them
type Employee_json struct {
	FirstName string `json:"name"`
	LastName  string `json:"lname"`
	Age       int    `json:"age"`
}

func json_struct() {
	json_string := `{
        "name":"Anucha",
        "lname":"Jan",
        "age":60
    }`
	emp := Employee_json{
		FirstName: "Suchat",
		LastName:  "jan",
		Age:       50,
	}
	jsonStr, _ := json.Marshal(emp)
	fmt.Printf("Struct Json : %s\n", jsonStr)

	var emp1 Employee_json // Create an instance of the struct

	// Unmarshal needs a byte slice and a pointer to the struct (&emp1)
	err := json.Unmarshal([]byte(json_string), &emp1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Print("Decode Json Struct : ")
	fmt.Printf("%+v\n", emp1) // %+v prints field names too
}

//-----------------------------------

type employee struct {
	name  string
	lname string
	age   int
}

func pointer_struct() {
	var emp3 = employee{"Nuth", "jan", 10}
	var ptr_emp *employee
	ptr_emp = &emp3

	fmt.Println(ptr_emp.name)
	fmt.Println(ptr_emp.lname)
	fmt.Println((*ptr_emp).age)
}
func main_struct() {
	var emp1 employee
	var emp2 employee

	// employee 1
	emp1.name = "Sittiporn"
	emp1.lname = "jang"
	emp1.age = 30

	// employee 2
	emp2.name = "Precha"
	emp2.lname = "Kuti"
	emp2.age = 20

	main_struct_print(emp1)
	main_struct_print(emp2)

}
func main_struct_print(std employee) {
	fmt.Printf("Name: %s, Lname: %s, Age: %d\n", std.name, std.lname, std.age)
}
