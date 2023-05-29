package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

type header struct{
	UserId string
	Browser string
}

type personData struct{
	FirstName string `binding:"required,alpha,min=4,max=10"`
	LastName string `binding:"required,alpha,min=4,max=10"`
	MobileNumber string `binding:"required,mobile,min=11,max=11"`
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"result" : "Test",
	})
}
func (h *TestHandler) Users(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"result" : "Users",
	})
}
func (h *TestHandler) UsersById(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result" : "UsersById",
		"id": id,
	})
}
func (h *TestHandler) UserByUserName(c *gin.Context){
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result" : "UserByUserName",
		"username": username,
	})
}
func (h *TestHandler) Accounts(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result" : "Accounts",
		"id": id,
	})
}
func (h *TestHandler) AddUser(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"result" : "AddUser",
	})
}

func (h *TestHandler) HeaderBinder1(c *gin.Context){
	userId :=  c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"result" : "HeaderBinder1",
		"userId": userId,
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context){
	header := header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK, gin.H{
		"result" : "HeaderBinder2",
		"header": header,
	})
}

func (h *TestHandler) QueryBinder1(c *gin.Context){
	id :=  c.Query("id")
	name :=  c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result" : "QueryBinder1",
		"id": id,
		"name": name,
	})
}

func (h *TestHandler) QueryBinder2(c *gin.Context){
	ids:= c.QueryArray("id")

	c.JSON(http.StatusOK, gin.H{
		"result" : "QueryBinder2",
		"ids": ids,
	})
}

func (h *TestHandler) UriBinder(c *gin.Context){
	id :=  c.Param("id")
	name :=  c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"result" : "UriBinder",
		"id": id,
		"name": name,
	})
}

func (h *TestHandler) BodyBinder(c *gin.Context){
	person := personData{}
	err := c.ShouldBindJSON(&person)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result" : "BodyBinder",
			"error" : err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result" : "BodyBinder",
		"person": person,
	})
}

func (h *TestHandler) FormBinder(c *gin.Context){
	person := personData{}
	c.ShouldBind(&person)
	c.JSON(http.StatusOK, gin.H{
		"result" : "FormBinder",
		"person": person,
	})
}

func (h *TestHandler) FileBinder(c *gin.Context){
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"result" : "FormBinder",
		"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result" : "FormBinder",
		"fileName": file.Filename,
	})
}