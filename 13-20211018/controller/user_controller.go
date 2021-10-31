package controller

import (
	"gin-postgres/model"
	"gin-postgres/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Lấy danh sách user
*/
func GetAllUser(c *gin.Context) {
	users, err := repo.GetAllUser()
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, users)
}

/*
Chi tiết thông tin user
*/
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := repo.GetUserById(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

/*
Tạo user
*/
func CreateUser(c *gin.Context) {
	req := new(model.CreateUser)

	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	user, err := repo.CreateUser(req)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

/*
Xóa user
*/
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := repo.DeleteUser(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, "Xóa user thành công")
}

/*
Cập nhật thông tin user
*/
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	req := new(model.CreateUser)
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	user, err := repo.UpdateUser(id, req)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, user)
}
