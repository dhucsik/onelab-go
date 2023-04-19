package http

import echoSwagger "github.com/swaggo/echo-swagger"

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	s.App.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	v1.POST("/sign-up", s.handler.SignUp)
	v1.POST("/sign-in", s.handler.SignIn, s.m.GenerateJWTMiddleware)
	v1.POST("/update-password", s.handler.UpdatePassword, s.m.Auth)
	v1.GET("/users-with-books", s.handler.GetUsersWithBooks)
	v1.GET("/users-with-books-for-month", s.handler.GetUsersWithBooksForMonth)

	v1.POST("/book", s.handler.CreateBook, s.m.Auth)
	v1.PUT("/book/:id", s.handler.UpdateBook, s.m.Auth)
	v1.GET("/book", s.handler.ListBooks)
	v1.GET("/book/:id", s.handler.GetBook)
	v1.DELETE("/book/:id", s.handler.DeleteBook, s.m.Auth)
	v1.GET("/book-users-income", s.handler.GetBooksUsersIncome)

	v1.POST("/rent", s.handler.CreateBookRent, s.m.Auth)
	v1.PUT("/rent/:id", s.handler.UpdateBookRent, s.m.Auth)
	v1.GET("/rent", s.handler.ListBookRents)
	v1.GET("/rent/:id", s.handler.GetBookRent)
	v1.DELETE("/rent/:id", s.handler.DeleteBookRent, s.m.Auth)
}
