package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	simbirhealth "github.com/ursulgwopp/simbir-health"
	"github.com/ursulgwopp/simbir-health/configs"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/repository"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/service"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/transport"
)

// @title Hospital Microservice
// @version 1.0

// @host localhost:8082
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(configs.Config{
		Host:     viper.GetString("hospitalMicroserviceDb.host"),
		Port:     viper.GetString("hospitalMicroserviceDb.port"),
		Username: viper.GetString("hospitalMicroserviceDb.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("hospitalMicroserviceDb.dbname"),
		SSLMode:  viper.GetString("hospitalMicroserviceDb.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewPostgresRepository(db)
	service := service.NewService(repo)
	transport := transport.NewTransport(service)

	srv := &simbirhealth.Server{}
	go func() {
		if err := srv.Run(viper.GetString("hospitalMicroservicePort"), transport.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
