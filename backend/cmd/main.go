package main

import (
	"context"
	"github.com/foxleren/auth-service/backend/internal/handler/rest"
	"github.com/foxleren/auth-service/backend/internal/models"
	"github.com/foxleren/auth-service/backend/internal/repository"
	"github.com/foxleren/auth-service/backend/internal/service"
	"github.com/foxleren/auth-service/backend/pkg/smtpService"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/siruspen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func main() {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if err := initConfig(); err != nil {
		logrus.Fatalf("Caught error while initializing config: ", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Caught error while loading .env file: ", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Caught error while creating database: ", err.Error())
	}

	emailService := smtpService.NewSmtpProvider(&smtpService.SMTPConfig{
		SenderEmail:    os.Getenv("SMTP_EMAIL"),
		SenderPassword: os.Getenv("SMTP_PASSWORD"),
		Host:           viper.GetString("smtp.host"),
		Port:           viper.GetString("smtp.port"),
	})

	repos := repository.NewRepository(db, emailService)
	services := service.NewService(repos)
	handlers := rest.NewHandler(services)

	srv := new(models.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error while running http server: ", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Server shutting down")

	err = srv.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("Error while shutting down http server: ", err.Error())
	}
}
