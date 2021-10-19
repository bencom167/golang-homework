package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
	Hàm liệt kê users theo filter
*/
func listUsers(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	var usersReturn []User
	var whereClause string
	var err error

	// Filter bằng fullname
	filterNames, ok := query["fullname"]
	if ok {
		whereClause = getFilterByName(whereClause, filterNames)
	}

	// Filter bằng tuổi
	filterAge, ok := query["age"]
	if ok {
		whereClause = getFilterByAge(whereClause, filterAge)
	}

	// Filter bằng giới tính
	filterSex, ok := query["sex"]
	if ok {
		whereClause = getFilterBySex(whereClause, filterSex)
	}

	if whereClause != "" {
		err = DB.Model(&usersReturn).Where(whereClause).Select()
	} else {
		err = DB.Model(&usersReturn).Select()
	}

	if err != nil {
		return err
	}

	usersJson, err := json.Marshal(usersReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJson)

	return nil
}

/*
	Hàm tạo user
*/
func createUsers(w http.ResponseWriter, r *http.Request) error {
	user := getUserFromUrlValue(r.URL.Query())

	//Them user vao DB
	var err error
	transaction, err := DB.Begin()
	if err != nil {
		logError(w, err)
		return err
	}

	_, err = transaction.Model(&user).Insert()
	if !check_err(err, transaction) {
		logError(w, err)
		return err
	}

	if err = transaction.Commit(); err != nil {
		logError(w, err)
		return err
	}

	// In thong tin user vua tao
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
	Lấy danh sách users theo danh sách fullname
*/
func getFilterByName(whereClause string, filterNames []string) string {
	if len(filterNames) == 0 {
		return whereClause
	}

	str := fmt.Sprintf("fullname ILIKE '%%%v%%'", filterNames[0])
	for i := 1; i < len(filterNames); i++ {
		str = fmt.Sprintf("%v OR fullname ILIKE '%%%v%%'", str, filterNames[i])
	}
	if whereClause != "" {
		whereClause = fmt.Sprintf("%v AND (%v)", whereClause, str)
	} else {
		whereClause = fmt.Sprintf("(%v)", str)
	}

	return whereClause
}

/*
	Lấy danh sách users theo tuổi
*/
func getFilterByAge(whereClause string, filterAge []string) string {
	if len(filterAge) == 0 {
		return whereClause
	}

	var str string
	age, ok := strconv.Atoi(filterAge[0])
	if ok == nil {
		str = fmt.Sprintf("age=%v", age)
	}

	for i := 1; i < len(filterAge); i++ {
		age, ok := strconv.Atoi(filterAge[i])
		if ok == nil {
			str = fmt.Sprintf("%v OR age=%v", str, age)
		}
	}

	if whereClause != "" {
		whereClause = fmt.Sprintf("%v AND (%v)", whereClause, str)
	} else {
		whereClause = fmt.Sprintf("(%v)", str)
	}

	return whereClause
}

/*
	Lấy danh sách users theo giới tính
*/
func getFilterBySex(whereClause string, filterSex []string) string {
	if len(filterSex) == 0 {
		return whereClause
	}

	str := fmt.Sprintf("sex='%v'", filterSex[0])
	for i := 1; i < len(filterSex); i++ {
		str = fmt.Sprintf("%v OR sex='%v'", str, filterSex[i])
	}
	if whereClause != "" {
		whereClause = fmt.Sprintf("%v AND (%v)", whereClause, str)
	} else {
		whereClause = fmt.Sprintf("(%v)", str)
	}

	return whereClause
}
