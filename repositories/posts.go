package repositories

import (
	"github.com/che-ict/DEV-DT-Microblog/models"
)

func GetAllPosts() []models.Post {
	var result []models.Post
	DB().Preload("User").Order("created_at desc").Find(&result)
	return result
}

func GetPostsByUser(username string) []models.Post {
	var user models.User
	var result []models.Post
	DB().Where("username = ?", username).First(&user)
	err := DB().Model(&user).Association("Posts").Find(&result)
	if err != nil {
		return nil
	}
	return result
}

func CreatePost(content, username string) {
	var user models.User
	DB().Where("username= ? ", username).First(&user)
	DB().Create(&models.Post{
		Content: content,
		UserId:  user.ID,
	})
}

func GetPost(id int) models.Post {
	var post models.Post
	DB().Preload("User").First(&post, id)
	return post
}

func DeletePost(id int) error {
	return DB().Delete(&models.Post{}, id).Error
}
