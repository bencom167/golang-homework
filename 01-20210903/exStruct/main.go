package main

import "fmt"

type Person struct {
	Id       string
	FullName string
	Email    string
}

/*
func (p Person) String() string {
	return fmt.Sprintf("%s : %s : %s", p.Id, p.FullName, p.Email)
}
*/

func main() {
	type personRequest struct {
		FullName string
		Email    string
	}

	pRequest := personRequest{
		FullName: "Trinh Minh Cuong",
		Email:    "cuong@techmaster.vn",
	}

	person := Person{
		Id:       "ox-13",
		FullName: pRequest.FullName,
		Email:    pRequest.Email,
	}

	// Là dạng viết tắt của fmt.Println(person.String())
	// Hàm tự định nghĩa sẽ override hàm mặc định của Go
	fmt.Println(person)
}
