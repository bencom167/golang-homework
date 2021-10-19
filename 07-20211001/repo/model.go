package repo

import (
	"math/rand"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Id       int
	FullName string
	Email    string
	Phone    string
	Age      int
	Sex      string
}

var (
	users    []User
	numUsers int = 10
)

func init() {
	users = []User{}
	for i := 0; i < numUsers; i++ {
		users = append(users, User{
			Id:       10000 + randNumber(89999), // Sinh ID: 10000 <= id <= 99998
			FullName: fake.Person().FirstName + " " + fake.Person().LastName,
			Email:    fake.Email(),
			Phone:    fake.Phone(),
			Age:      22 + randNumber(44), // Sinh tuổi: 22 <= age <= 65
			Sex:      fake.Person().Gender,
		})
	}
}

func randNumber(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}
