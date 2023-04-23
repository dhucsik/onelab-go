package http

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	s.App.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	v1.POST("/sign-up", s.handler.SignUp)
	v1.POST("/sign-in", s.handler.SignIn, s.m.GenerateJWTMiddleware)
	v1.POST("/update-password", s.handler.UpdatePassword, s.m.Auth)
	v1.GET("/users-with-books", s.handler.GetUsersWithBooks)
	v1.GET("/users-with-books-for-month", s.handler.GetUsersWithBooksForMonth)

	v1.POST("/book", func(c echo.Context) error {
		return s.handler.CreateBook(c, s.logger)
	}, s.m.Auth)
	v1.PUT("/book/:id", func(c echo.Context) error {
		return s.handler.UpdateBook(c, s.logger)
	}, s.m.Auth)
	v1.GET("/book", func(c echo.Context) error {
		return s.handler.ListBooks(c, s.logger)
	})
	v1.GET("/book/:id", func(c echo.Context) error {
		return s.handler.GetBook(c, s.logger)
	})
	v1.DELETE("/book/:id", func(c echo.Context) error {
		return s.handler.DeleteBook(c, s.logger)
	}, s.m.Auth)
	v1.GET("/book-users-income", func(c echo.Context) error {
		return s.handler.GetBooksUsersIncome(c, s.logger)
	})

	v1.POST("/rent", s.handler.CreateBookRent, s.m.Auth)
	v1.PUT("/rent/:id", s.handler.UpdateBookRent, s.m.Auth)
	v1.GET("/rent", s.handler.ListBookRents)
	v1.GET("/rent/:id", s.handler.GetBookRent)
	v1.DELETE("/rent/:id", s.handler.DeleteBookRent, s.m.Auth)
}
