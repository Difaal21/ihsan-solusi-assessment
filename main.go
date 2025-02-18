package main

import (
	"difaal21/ihsan-solusi-assessment/config"
	"difaal21/ihsan-solusi-assessment/database/postgresql"
	"difaal21/ihsan-solusi-assessment/modules/users"
	"difaal21/ihsan-solusi-assessment/repositories"
	"difaal21/ihsan-solusi-assessment/responses"
	"difaal21/ihsan-solusi-assessment/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload" //for development
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

var cfg *config.Config
var response = responses.NewResponse()

func init() {
	cfg = config.Load()
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(cfg.Logger.Formatter)
	logger.SetReportCaller(true)

	validate := validator.New()

	postgresqlClient := postgresql.NewPostgreSQL(cfg.PostgreSQL.Driver, cfg.PostgreSQL.DSN, logger)
	pdb, err := postgresqlClient.Connect(cfg.PostgreSQL.MaxOpenConnections, cfg.PostgreSQL.MaxIdleConnections)
	if err != nil {
		logger.Fatal("Could not connect to PostgreSQL: ", err)
	}

	router := echo.New()
	router.GET("/ihsan-solusi-assessment", index)

	usersRepo := repositories.NewUserRepository(logger, pdb)
	usersUsecase := users.NewUseCase(logger, usersRepo)
	users.NewHTTPHandler(router, logger, validate, usersUsecase)

	handler := cors.New(cors.Options{
		AllowedOrigins:   cfg.Application.AllowedOrigins,
		AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	server := server.NewServer(logger, handler, cfg.Application.Port)
	server.Start()

	// When we run this program it will block waiting for a signal. By typing ctrl-C, we can send a SIGINT signal, causing the program to print interrupt and then exit.
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm

	// closing service for a gracefull shutdown.
	defer func() {
		// mysqlClient.Close(mdb)
		postgresqlClient.Close(pdb)
		server.Close()
	}()
}

var (
	healthCheckMessage  = "Application running properly"
	pageNotFoundMessage = "You're lost, double check the endpoint"
)

func index(c echo.Context) error {
	resp := response.Ok("").SetData(nil).SetMessage(healthCheckMessage).Send()
	return c.JSON(resp.Code, response)
}

func notFoundHandler(c echo.Context) error {
	resp := response.NotFound("").SetData(nil).SetMessage(pageNotFoundMessage).Send()
	return c.JSON(resp.Code, response)
}
