package server

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

type Server struct {
	Port       int
	Addr       string
	HTTPServer *http.Server
}

func (e *Server) Start() {
	log.Println("Server started on port", e.Port)
	e.Logger.Fatal(e.Start(e.Addr))
	e.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func NewServer(port int, db *xorm.Engine) *Server {
	s := echo.New()
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())
	var s Server
	s.Port = port
	s.Addr = ":" + strconv.Itoa(port)

	s.Use(middleware.CORS())
	s.HTTPServer = &http.Server{
		Addr:         s.Addr,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	  }

	return &s
}