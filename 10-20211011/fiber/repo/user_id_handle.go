package repo

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

/*
	Hàm trả lại thông tin user theo id
*/
func getUser(c *fiber.Ctx) error {
	var usersReturn []User
	var err error

	// Filter bằng id
	filterId := c.Query("id")
	if filterId != "" {
		id, _ := strconv.Atoi(filterId)
		err = DB.Model(&usersReturn).Where("id = ?", id).Select()
	} else {
		err = DB.Model(&usersReturn).Select()
	}

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(usersReturn)
}

/*
	Hàm cập nhật thông tin user theo id
*/
func updateUser(c *fiber.Ctx) error {
	var err error
	var user User

	if err = c.QueryParser(&user); err != nil {
		return c.Status(400).JSON(err)
	}

	if user, err = matchUserData(user); err != nil {
		return c.Status(400).JSON(err)
	}

	//Update user vao DB
	transaction, err := DB.Begin()
	if err != nil {
		return c.Status(400).JSON(err)
	}

	_, err = transaction.Model(&user).Where("id = ?", user.Id).Update()
	if !check_err(err, transaction) {
		return c.Status(400).JSON(err)
	}

	if err = transaction.Commit(); err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(user)
}

/*
	Hàm xoá user theo id
*/
func deleteUser(c *fiber.Ctx) error {
	var err error
	var usersReturn []User

	//Update user vao DB
	transaction, err := DB.Begin()
	if err != nil {
		return c.Status(400).JSON(err)
	}

	// Filter bằng id
	filterId := c.Query("id")
	if filterId != "" {
		id, _ := strconv.Atoi(filterId)
		_, err = transaction.Model(&usersReturn).Where("id = ?", id).Delete()

		if !check_err(err, transaction) {
			return c.Status(400).JSON(err)
		}
	}

	if err = transaction.Commit(); err != nil {
		return c.Status(400).JSON(err)
	}

	err = DB.Model(&usersReturn).Select()
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(usersReturn)
}

func matchUserData(user User) (User, error) {
	var userOrigin []User

	err := DB.Model(&userOrigin).Where("id=?", user.Id).Select()
	if err != nil {
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
