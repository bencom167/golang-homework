package repo

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

/*
	Hàm liệt kê users theo filter
*/
func listUsers(c *fiber.Ctx) error {
	var usersReturn []User
	var whereClause string
	var err error

	// Filter bằng fullname
	whereClause = getFilterByName(whereClause, c.Query("fullname"))

	// Filter bằng tuổi
	whereClause = getFilterByAge(whereClause, c.Query("age"))

	// Filter bằng giới tính
	whereClause = getFilterBySex(whereClause, c.Query("sex"))

	if whereClause != "" {
		err = DB.Model(&usersReturn).Where(whereClause).Select()
	} else {
		err = DB.Model(&usersReturn).Select()
	}

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(usersReturn)
}

/*
	Hàm tạo user
*/
func createUsers(c *fiber.Ctx) error {
	var err error
	var user User

	if err := c.QueryParser(&user); err != nil {
		return c.Status(400).JSON(err)
	}

	//Them user vao DB
	transaction, err := DB.Begin()
	if err != nil {
		return c.Status(400).JSON(err)
	}

	_, err = transaction.Model(&user).Insert()
	if !check_err(err, transaction) {
		return c.Status(400).JSON(err)
	}

	if err = transaction.Commit(); err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(user)
}

/*
	Lấy danh sách users theo danh sách fullname
*/
func getFilterByName(whereClause string, filterNames string) string {
	if filterNames == "" {
		return whereClause
	}

	str := fmt.Sprintf("fullname ILIKE '%%%v%%'", filterNames)

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
func getFilterByAge(whereClause string, filterAge string) string {
	if filterAge == "" {
		return whereClause
	}

	var str string
	age, ok := strconv.Atoi(filterAge)
	if ok == nil {
		str = fmt.Sprintf("age=%v", age)
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
func getFilterBySex(whereClause string, filterSex string) string {
	if filterSex == "" {
		return whereClause
	}

	str := fmt.Sprintf("sex='%v'", filterSex)
	if whereClause != "" {
		whereClause = fmt.Sprintf("%v AND (%v)", whereClause, str)
	} else {
		whereClause = fmt.Sprintf("(%v)", str)
	}

	return whereClause
}
