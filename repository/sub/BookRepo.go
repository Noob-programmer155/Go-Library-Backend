package sub

import (
	"amrDev/libraryBackend.com/models"
	"database/sql"
	"fmt"
	"time"
)

func FindAllByDateReportBookBetween(db *sql.DB, start time.Time, end time.Time) (books []models.BookReport, err error) {
	query := fmt.Sprintf("select * from bookreport where dateReport between '%s' and '%s'",
		start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	res, _ := db.Query(query)
	defer res.Close()
	var id, idAuthor, idPublisher, idUser uint64
	var idBook, titleBook, nameAuthor, emailAuthor, nameUser, emailUser, namePublisher string
	var dateReport time.Time
	var statusReport int
	for res.Next() {
		err = res.Scan(&id, &dateReport, &emailUser, &emailAuthor,
			&idAuthor, &idBook, &idPublisher, &idUser,
			&nameAuthor, &namePublisher, &statusReport, &titleBook, &nameUser)
		if err == nil {
			books = append(books, models.BookReport{
				Id:        id,
				IdBook:    idBook,
				TitleBook: titleBook,
				Author: models.UserBookReport{
					Id:    idAuthor,
					Name:  nameAuthor,
					Email: emailAuthor,
				},
				Publisher: models.Publisher{
					Id:   idPublisher,
					Name: namePublisher,
				},
				User: models.UserBookReport{
					Id:    idUser,
					Name:  nameUser,
					Email: emailUser,
				},
				StatusReport: statusReport,
				DateReport:   dateReport,
			})
		}
	}
	return
}
