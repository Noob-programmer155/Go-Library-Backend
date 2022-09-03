package configuration

import (
	"amrDev/libraryBackend.com/repository"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

var filePath, imagePath, imageEks, fileEks string

func init() {
	imagePath = os.Getenv("PATH_STORE_IMAGE")
	if imagePath != "" {
		log.Fatal("Path to store images has not found !!!")
	} else {
		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			err = os.MkdirAll(imagePath+"/users", os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			err = os.MkdirAll(imagePath+"/books", os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	filePath = os.Getenv("PATH_STORE_FILE")
	if filePath != "" {
		log.Fatal("Path to store files has not found !!!")
	} else {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			err = os.MkdirAll(filePath+"/temp", os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	imageEks = os.Getenv("IMAGE_EKS")
	if imageEks != "" {
		log.Fatal("images ekstension has not found !!!")
	}
	fileEks = os.Getenv("FILE_EKS")
	if fileEks != "" {
		log.Fatal("files ekstension has not found !!!")
	}

}

func SaveImageBook(h *http.Request, name string) (err error) {
	err = h.ParseMultipartForm(2000000)
	if err != nil {
		return
	}
	file, err := os.Create(imagePath + "/books/" + name + imageEks)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	data, _, err := h.FormFile("image")
	if err != nil {
		return
	}
	defer func() {
		err = data.Close()
		if err != nil {
			return
		}
	}()
	_, err = io.Copy(file, data)
	if err != nil {
		return
	}
	log.Println("Image Saved")
	return
}

func SaveImageUser(h *http.Request, name string) (err error) {
	err = h.ParseMultipartForm(2000000)
	if err != nil {
		return
	}
	file, err := os.Create(imagePath + "/users/" + name + imageEks)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	data, _, err := h.FormFile("image")
	if err != nil {
		return
	}
	defer func() {
		err = data.Close()
		if err != nil {
			return
		}
	}()
	_, err = io.Copy(file, data)
	if err != nil {
		return
	}
	log.Println("Image Saved")
	return
}

func SaveFileBook(h *http.Request, name string) (err error) {
	err = h.ParseMultipartForm(500000000)
	if err != nil {
		return
	}
	file, err := os.Create(filePath + "/" + name + imageEks)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()
	data, _, err := h.FormFile("file")
	if err != nil {
		return
	}
	defer func() {
		err = data.Close()
		if err != nil {
			return
		}
	}()
	_, err = io.Copy(file, data)
	if err != nil {
		return
	}
	log.Println("File Saved")
	return
}

func DeleteFile(subPath string, name string, isImage bool) (err error) {
	var path string
	if isImage {
		if name[0:4] == "http" {
			return
		} else {
			path = imagePath
		}
	} else {
		path = filePath
	}
	if subPath == "" {
		subPath = "/"
	}
	err = os.Remove(path + subPath + name)
	return
}

func GetReportFile(start time.Time, end time.Time, w http.ResponseWriter) (err error) {
	users, err := repository.FindUserBetween2Date(start, end)
	books, err := repository.FindBookBetween2Date(start, end)

	file := excelize.NewFile()
	file.SetSheetName("Sheet1", "Book Report")
	file.NewSheet("User Report")
	style, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{Bold: true},
	})

	file.SetCellStyle("Book Report", "A1", "L3", style)

	//set header book report sheet
	file.MergeCell("Book Report", "A1", "C2")
	file.SetCellValue("Book Report", "A1", "User")
	file.MergeCell("Book Report", "D1", "J1")
	file.SetCellValue("Book Report", "D1", "Book")
	file.MergeCell("Book Report", "K1", "K3")
	file.SetCellValue("Book Report", "K1", "Status")
	file.MergeCell("Book Report", "L1", "L3")
	file.SetCellValue("Book Report", "A1", "Date Report")
	file.SetCellValue("Book Report", "A3", "User ID")
	file.SetCellValue("Book Report", "B3", "Username")
	file.SetCellValue("Book Report", "C3", "User Email")
	file.MergeCell("Book Report", "D2", "D3")
	file.SetCellValue("Book Report", "D2", "Book ID")
	file.MergeCell("Book Report", "E2", "E3")
	file.SetCellValue("Book Report", "E2", "Book Title")
	file.MergeCell("Book Report", "F2", "H2")
	file.SetCellValue("Book Report", "F2", "Author")
	file.SetCellValue("Book Report", "F3", "Author ID")
	file.SetCellValue("Book Report", "G3", "Author Name")
	file.SetCellValue("Book Report", "H3", "Author Email")
	file.MergeCell("Book Report", "I2", "J2")
	file.SetCellValue("Book Report", "I2", "Publisher")
	file.SetCellValue("Book Report", "I3", "Publisher ID")
	file.SetCellValue("Book Report", "J3", "Publisher Name")

	for index, book := range books {
		file.SetCellValue("Book Report", fmt.Sprintf("A%d", index+4), book.User.Id)
		file.SetCellValue("Book Report", fmt.Sprintf("B%d", index+4), book.User.Name)
		file.SetCellValue("Book Report", fmt.Sprintf("C%d", index+4), book.User.Email)
		file.SetCellValue("Book Report", fmt.Sprintf("D%d", index+4), book.IdBook)
		file.SetCellValue("Book Report", fmt.Sprintf("E%d", index+4), book.TitleBook)
		file.SetCellValue("Book Report", fmt.Sprintf("F%d", index+4), book.Author.Id)
		file.SetCellValue("Book Report", fmt.Sprintf("G%d", index+4), book.Author.Name)
		file.SetCellValue("Book Report", fmt.Sprintf("H%d", index+4), book.Author.Email)
		file.SetCellValue("Book Report", fmt.Sprintf("I%d", index+4), book.Publisher.Id)
		file.SetCellValue("Book Report", fmt.Sprintf("J%d", index+4), book.Publisher.Name)
		file.SetCellValue("Book Report", fmt.Sprintf("K%d", index+4), book.StatusReport)
		file.SetCellValue("Book Report", fmt.Sprintf("L%d", index+4), book.DateReport)
	}

	// user sheet
	file.SetCellStyle("User Report", "A1", "H2", style)

	file.MergeCell("User Report", "A1", "C1")
	file.SetCellValue("User Report", "A1", "User")
	file.SetCellValue("User Report", "A2", "User ID")
	file.SetCellValue("User Report", "B2", "Username")
	file.SetCellValue("User Report", "C2", "User Email")
	file.MergeCell("User Report", "D1", "F1")
	file.SetCellValue("User Report", "D1", "Admin")
	file.SetCellValue("User Report", "D2", "Admin ID")
	file.SetCellValue("User Report", "E2", "Admin Name")
	file.SetCellValue("User Report", "F2", "Admin Email")
	file.MergeCell("User Report", "G1", "G2")
	file.SetCellValue("User Report", "G1", "Status")
	file.MergeCell("User Report", "H1", "H2")
	file.SetCellValue("User Report", "H1", "Date Report")

	for index, user := range users {
		file.SetCellValue("User Report", fmt.Sprintf("A%d", index+3), user.User.Id)
		file.SetCellValue("User Report", fmt.Sprintf("B%d", index+3), user.User.Name)
		file.SetCellValue("User Report", fmt.Sprintf("C%d", index+3), user.User.Email)
		file.SetCellValue("User Report", fmt.Sprintf("D%d", index+3), user.Admin.Id)
		file.SetCellValue("User Report", fmt.Sprintf("E%d", index+3), user.Admin.Name)
		file.SetCellValue("User Report", fmt.Sprintf("F%d", index+3), user.Admin.Email)
		file.SetCellValue("User Report", fmt.Sprintf("G%d", index+3), user.StatusReport)
		file.SetCellValue("User Report", fmt.Sprintf("H%d", index+3), user.DateReport)
	}
	uid := time.Now().Format("2006-01-02T15:04:05")
	if err := file.SaveAs(filePath + "/temp/" + uid + "-report.xlsx"); err != nil {
		log.Fatal(err)
	} else {
		w.Header().Add("Content-Description", "Download File")
		w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename='%s'", filePath+"/temp/"+uid+"-report.xlsx"))
		w.Header().Add("Content-Transfer-Encoding", "binary")
	}
	return
}
