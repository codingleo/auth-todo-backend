package api

import (
	"net/http"

	"github.com/codingleo/auth-todo-backend/database"
	"github.com/codingleo/auth-todo-backend/types"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	listenAddr string
}

type apiFunc func(c *gin.Context) error

func makeHTTPHandleFunc(f apiFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Start() error {
	router := gin.Default()
	router.GET("/users", makeHTTPHandleFunc(getUsers))
	router.POST("/users", makeHTTPHandleFunc(createUser))

	return router.Run(s.listenAddr)
}

func getUsers(c *gin.Context) error {
	var users []types.User
	err := database.Db.Select("id", "first_name", "last_name", "email").Find(&users).Error

	c.JSON(http.StatusOK, users)
	return err
}

func createUser(c *gin.Context) error {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		return err
	}

	if errs := user.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, errs)
		return nil
	}

	if err := database.Db.Create(&user).Error; err != nil {
		return err
	}

	user.Password = ""

	c.JSON(http.StatusOK, user)
	return nil
}
