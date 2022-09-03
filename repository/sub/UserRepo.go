package sub

import (
	"amrDev/libraryBackend.com/models"
	"database/sql"
	"fmt"
	"time"
)

func FindAllByDateReportUserBetween(db *sql.DB, start time.Time, end time.Time) (users []models.UserReport, err error) {
	query := fmt.Sprintf("select * from userreport where dateReport between '%s' and '%s'",
		start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	res, _ := db.Query(query)
	defer res.Close()
	var id, idAdmin, idUser uint64
	var nameAdmin, emailAdmin, nameUser, emailUser string
	var dateReport time.Time
	var statusReport int
	for res.Next() {
		err = res.Scan(&id, &emailAdmin, &nameAdmin, &dateReport,
			&emailUser, &idAdmin, &idUser, &statusReport,
			&nameUser)
		if err == nil {
			users = append(users, models.UserReport{
				Id: id,
				User: models.UserBookReport{
					Id:    idUser,
					Name:  nameUser,
					Email: emailUser,
				},
				Admin: models.UserBookReport{
					Id:    idAdmin,
					Name:  nameAdmin,
					Email: emailAdmin,
				},
				StatusReport: statusReport,
				DateReport:   dateReport,
			})
		}
	}
	return
}

func CheckByEmailUser(db *sql.DB, email string) (b bool, err error) {
	res, err := db.Query("select id from user where email = ?", email)
	defer func() {
		err = res.Close()
		if err != nil {
			b = false
			return
		}
	}()
	if err != nil {
		return false, err
	}

	if res.Next() {
		return true, nil
	}

	return false, nil
}

func FindByEmailUserValidation(db *sql.DB, email string) (data models.User, err error) {
	res, err := db.Query("select id,name,email from user where email = ?", email)
	defer func() {
		err = res.Close()
		if err != nil {
			return
		}
	}()
	if err != nil {
		return
	}

	if res.Next() {
		err = res.Scan(&data.Id, &data.Name, &data.Email)
		if err != nil {
			return
		}
	}
	return
}

func UpdateUser(user models.User, db *sql.DB) error {
	query := "update user set "
	if user.Name != "" {
		query += fmt.Sprintf("name = '%s',", user.Name)
	}
	if user.Password != "" {
		query += fmt.Sprintf("password = '%s',", user.Password)
	}
	if user.Email != "" {
		query += fmt.Sprintf("email = '%s',", user.Email)
	}
	if user.Role != -1 {
		query += fmt.Sprintf("role = %d,", user.Role)
	}
	if user.Provider != -1 {
		query += fmt.Sprintf("provider = %d,", user.Provider)
	}
	if user.ImageURL != "" {
		query += fmt.Sprintf("image_url = '%s',", user.ImageURL)
	}
	if user.ClientId != "" {
		query += fmt.Sprintf("clientId = '%s',", user.ClientId)
	}
	if query[len(query)-1:] == "," {
		query = query[:len(query)-1]
	}
	query += fmt.Sprintf(" where id = %d,", user.Id)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
