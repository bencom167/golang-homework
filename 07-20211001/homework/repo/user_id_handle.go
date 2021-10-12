package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
	Hàm trả lại thông tin user theo id
*/
func getUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	index := -1

	id, ok := query["id"]
	if ok {
		selectID, _ := strconv.Atoi(strings.Join(id, ""))
		index = findUserByID(selectID)
	}

	if index == -1 {
		fmt.Fprintln(w, "User is not found!")
		return
	}

	usersJson, err := json.Marshal(users[index])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJson)
}

/*
	Hàm cập nhật thông tin user theo id
*/
func updateUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	index := -1

	id, ok := query["id"]
	if ok {
		selectID, _ := strconv.Atoi(strings.Join(id, ""))
		index = findUserByID(selectID)
	}

	if index == -1 {
		fmt.Fprintln(w, "User is not found!")
		return
	}

	fullname, ok := query["fullname"]
	if ok {
		users[index].FullName = strings.Join(fullname, ",")
	}

	email, ok := query["email"]
	if ok {
		users[index].Email = strings.Join(email, ",")
	}

	phone, ok := query["phone"]
	if ok {
		users[index].Phone = strings.Join(phone, ",")
	}

	age, ok := query["age"]
	if ok {
		users[index].Age, _ = strconv.Atoi(strings.Join(age, ""))
	}

	sex, ok := query["sex"]
	if ok {
		users[index].Sex = strings.Join(sex, ",")
	}

	userJson, err := json.Marshal(users[index])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

/*
	Hàm xoá user theo id
*/
func deleteUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	index := -1

	id, ok := query["id"]
	if ok {
		selectID, _ := strconv.Atoi(strings.Join(id, ""))
		index = findUserByID(selectID)
		if index >= 0 {
			removeUserKeepOrder(index)
		}
	}

	userJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

/*
	Hàm xoá user, giữ thứ tự
*/
func removeUserKeepOrder(index int) {
	for i := index; i < len(users)-1; i++ {
		users[i] = users[i+1]
	}
	users = users[:len(users)-1]
}

/*
	Hàm tìm vị trí user bằng ID
*/
func findUserByID(id int) int {
	for i := 1; i < len(users); i++ {
		if id == users[i].Id {
			return i
		}
	}
	return -1
}
