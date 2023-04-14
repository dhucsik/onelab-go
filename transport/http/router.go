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

	v1.POST("/record", s.handler.CreateRecord, s.m.Auth)
	v1.PUT("/record/:id", s.handler.UpdateRecord, s.m.Auth)
	v1.GET("/record", s.handler.ListRecords)
	v1.GET("/record/:id", s.handler.GetRecord)
	v1.DELETE("/record/:id", s.handler.DeleteRecord, s.m.Auth)
}
