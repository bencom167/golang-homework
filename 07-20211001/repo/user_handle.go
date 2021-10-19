package repo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

/*
	Hàm liệt kê users theo filter
*/
func listUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var usersReturn []User

	// Filter bằng fullname
	filterNames, ok := query["fullname"]
	if ok {
		usersReturn = filterByName(users, filterNames)
	} else {
		usersReturn = append(usersReturn, users...)
	}

	// Filter bằng tuổi
	age, ok := query["age"]
	if ok {
		filterAge, _ := strconv.Atoi(strings.Join(age, ""))
		usersReturn = filterByAge(usersReturn, filterAge)
	}

	// Filter bằng giới tính
	filterSex, ok := query["sex"]
	if ok {
		usersReturn = filterBySex(usersReturn, filterSex)
	}

	usersJson, err := json.Marshal(usersReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJson)
}

/*
	Hàm tạo user
*/
func createUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var user User

	id, ok := query["id"]
	if ok {
		user.Id, _ = strconv.Atoi(strings.Join(id, ""))
	}

	fullname, ok := query["fullname"]
	if ok {
		user.FullName = strings.Join(fullname, ",")
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

	//Them user vao danh sach
	users = append(users, user)

	userJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

/*
	Lấy danh sách users theo danh sách fullname
*/
func filterByName(users []User, filterNames []string) (result []User) {
	for _, user := range users {
		isFilter := false
		for _, fullname := range filterNames {
			if user.FullName == fullname {
				isFilter = true
				break
			}
		}
		if isFilter {
			result = append(result, user)
		}
	}
	return result
}

/*
	Lấy danh sách users theo tuổi
*/
func filterByAge(users []User, filterAge int) (result []User) {
	for _, user := range users {
		if user.Age == filterAge {
			result = append(result, user)
		}
	}
	return result
}

/*
	Lấy danh sách users theo giới tính
*/
func filterBySex(users []User, filterSex []string) (result []User) {
	for _, user := range users {
		isFilter := false
		for _, sex := range filterSex {
			if user.Sex == sex {
				isFilter = true
				break
			}
		}
		if isFilter {
			result = append(result, user)
		}
	}
	return result
}
