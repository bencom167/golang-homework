package repo

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func getUserFromUrlValue(query url.Values) User {
	var user User

	id, ok := query["id"]
	if ok {
		user.Id, _ = strconv.Atoi(strings.Join(id, ""))
	}

	fullname, ok := query["fullname"]
	if ok {
		user.Fullname = strings.Join(fullname, ",")
	}

	email, ok := query["email"]
	if ok {
		user.Email = strings.Join(email, ",")
	}

	phone, ok := query["phone"]
	if ok {
		user.Phone = strings.Join(phone, ",")
	}

	age, ok := query["age"]
	if ok {
		user.Age, _ = strconv.Atoi(strings.Join(age, ""))
	}

	sex, ok := query["sex"]
	if ok {
		user.Sex = strings.Join(sex, ",")
	}

	return user
}

func logError(w http.ResponseWriter, err error) {
	fmt.Fprintln(w, "500 Internal Server!")
	log.Printf("%v", err)
}
