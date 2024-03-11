package handlers

import (
	"github.com/che-ict/DEV-DT-Microblog/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func FeedHandler(c echo.Context) error {
	posts := repositories.GetAllPosts()
	user, _ := c.Cookie("user")
	m := echo.Map{
		"Posts":             posts,
		"AuthenticatedUser": user.Value,
	}
	return c.Render(http.StatusOK, "feed", m)
}

func UserPageHandler(c echo.Context) error {
	username := c.Param("user")
	posts := repositories.GetPostsByUser(username)
	user, _ := c.Cookie("user")
	m := echo.Map{
		"Posts":             posts,
		"AuthenticatedUser": user.Value,
	}
	return c.Render(http.StatusOK, "user", m)
}

func CreatePostHandler(c echo.Context) error {
	user, err := c.Cookie("user")
	content := c.FormValue("content")
	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	repositories.CreatePost(content, user.Value)
	return c.Redirect(http.StatusFound, "/feed")
}

func DeletePostHandler(c echo.Context) error {
	postId, _ := strconv.Atoi(c.Param("id"))
	err := repositories.DeletePost(postId)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/feed")
}
