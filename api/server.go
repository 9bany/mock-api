package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	route *gin.Engine
	data  []byte
}

func NewServer(data []byte) *Server {
	r := gin.Default()
	server := &Server{
		route: r,
		data:  data,
	}

	server.setUp(r)
	return server
}

func (s *Server) setUp(route *gin.Engine) {
	route.GET("/mock", s.MockHandle)
}

func (s *Server) MockHandle(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var dataResp map[string]interface{}
	err := json.Unmarshal(s.data, &dataResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, dataResp)
}

func (s *Server) Run() {
	s.route.Run("0.0.0.0:8080")
}
