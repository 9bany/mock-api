package listen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	route *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	server := &Server{
		route: r,
	}

	server.setUp(r)
	return server
}

func (s *Server) setUp(route *gin.Engine) {
	route.POST("/push", s.MushHandle)
}

func (s *Server) MushHandle(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var dataResp map[string]interface{}
	err = json.Unmarshal(jsonData, &dataResp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(dataResp)
	c.JSON(http.StatusOK, dataResp)
}

func (s *Server) Run() {
	s.route.Run("0.0.0.0:8080")
}
