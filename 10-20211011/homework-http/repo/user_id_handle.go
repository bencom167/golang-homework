package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
	Hàm trả lại thông tin user theo id
*/
func getUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var usersReturn []User

	id, ok := r.URL.Query()["id"]
	if ok {
		selectID, _ := strconv.Atoi(strings.Join(id, ""))
		err = DB.Model(&usersReturn).Where("id=?", selectID).Select()
	}

	if err != nil {
		logError(w, err)
	}
	if len(usersReturn) == 0 {
		fmt.Fprintln(w, "User is not found!")
		return
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
	Hàm cập nhật thông tin user theo id
*/
func updateUser(w http.ResponseWriter, r *http.Request) error {
	var err error

	user, err := matchUserInfor(w, getUserFromUrlValue(r.URL.Query()))
	if err != nil {
		logError(w, err)
		return err
	}

	//Update user vao DB
	transaction, err := DB.Begin()
	if err != nil {
		logError(w, err)
		return err
	}

	_, err = transaction.Model(&user).Where("id=?", user.Id).Update()
	if !check_err(err, transaction) {
		logError(w, err)
		return err
	}

	if err = transaction.Commit(); err != nil {
		logError(w, err)
		return err
	}

	// In thong tin user
	userJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)

	return nil
}

/*
	Hàm xoá user theo id
*/
func deleteUser(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	var usersReturn []User
	var err error

	id, ok := query["id"]
	if ok {
		selectID, _ := strconv.Atoi(strings.Join(id, ""))
		//Xoa user trong DB
		transaction, err := DB.Begin()
		if err != nil {
			logError(w, err)
			return err
		}

		_, err = transaction.Model(&usersReturn).Where("id=?", selectID).Delete()
		if !check_err(err, transaction) {
			logError(w, err)
			return err
		}

		if err = transaction.Commit(); err != nil {
			logError(w, err)
			return err
		}
	}

	err = DB.Model(&usersReturn).Select()
	if err != nil {
		logError(w, err)
		return err
	}

	userJson, err := json.Marshal(usersReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)

	return nil
}

func matchUserInfor(w http.ResponseWriter, user User) (User, error) {
	var userOrigin []User

	err := DB.Model(&userOrigin).Where("id=?", user.Id).Select()
	if err != nil {
		logError(w, err)
		return user, err
	}

	if len(userOrigin) == 0 {
		return user, errors.New("User is not found")
	}

	if user.Fullname == "" {
		user.Fullname = userOrigin[0].Fullname
	}
	if user.Email == "" {
		user.Email = userOrigin[0].Email
	}
	if user.Phone == "" {
		user.Phone = userOrigin[0].Phone
	}
	if user.Age == 0 {
		user.Age = userOrigin[0].Age
	}
	if user.Sex == "" {
		user.Sex = userOrigin[0].Sex
	}

	return user, nil
}
