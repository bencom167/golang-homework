package controller

import (
	"gin-postgres/model"
	"gin-postgres/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Lấy danh sách post của user
*/
func GetPostsOfUser(c *gin.Context) {
	userId := c.Param("id")
	posts, err := repo.GetPostsOfUser(userId)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, posts)
}

/*
Lấy thông tin của post
*/
func GetPostDetail(c *gin.Context) {
	userId := c.Param("id")
	postId := c.Param("postId")

	post, err := repo.GetPostDetail(userId, postId)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, post)
}

/*
Tạo post
*/
func CreatePost(c *gin.Context) {
	userId := c.Param("id")

	req := new(model.CreatePost)
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	post, err := repo.CreatePost(userId, req)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, post)
}

/*
Xóa post
*/
func DeletePost(c *gin.Context) {
	userId := c.Param("id")
	postId := c.Param("postId")

	err := repo.DeletePost(userId, postId)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, "Xóa post thành công")
}

/*
Update post
*/
func UpdatePost(c *gin.Context) {
	userId := c.Param("id")
	postId := c.Param("postId")

	req := new(model.CreatePost)
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	post, err := repo.UpdatePost(userId, postId, req)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, post)
}
