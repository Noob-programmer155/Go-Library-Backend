package repository

import (
	"amrDev/libraryBackend.com/models"
	"amrDev/libraryBackend.com/repository/sub"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var db *sql.DB

func InitDB(server *http.Server) {
	usename, password, host, port, database := os.Getenv("USERNAME_DB"), os.Getenv("PASSWORD_DB"), os.Getenv("HOST_DB"),
		os.Getenv("PORT_DB"), os.Getenv("DATABASE_DB")
	dbdata, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", usename, password, host, port, database))
	if err != nil {
		log.Fatal("Cannot Connect to Database")
	}
	db = dbdata

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("Cannot Close Database")
		} else {
			log.Println("Database Closed !!!")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	<-stop
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Cannot shutdown server")
	} else {
		log.Println("Server is Shutdown !!!")
	}
}

func CheckByEmail(email string) (bool, error) {
	return sub.CheckByEmailUser(db, email)
}

func FindUserBetween2Date(start time.Time, end time.Time) ([]models.UserReport, error) {
	return sub.FindAllByDateReportUserBetween(db, start, end)
}

func FindIdNameEmailByEmail(email string) (models.User, error) {
	return sub.FindByEmailUserValidation(db, email)
}

func Update(user models.User) error {
	return sub.UpdateUser(user, db)
}

func FindBookBetween2Date(start time.Time, end time.Time) ([]models.BookReport, error) {
	return sub.FindAllByDateReportBookBetween(db, start, end)
}
