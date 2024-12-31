package controller

import (
	"a1_bms_project/model"
	"a1_bms_project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
	return &BlogController{BlogService: blogService}
}

func (controller *BlogController) CreateBlog(c *gin.Context) {
	var blog model.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBlog, err := controller.BlogService.CreateBlog(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdBlog)
}

func (controller *BlogController) GetBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	blog, err := controller.BlogService.GetBlog(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (controller *BlogController) GetAllBlogs(c *gin.Context) {
	blogs, err := controller.BlogService.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (controller *BlogController) UpdateBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog model.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.ID = id
	updatedBlog, err := controller.BlogService.UpdateBlog(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

func (controller *BlogController) DeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.BlogService.DeleteBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
