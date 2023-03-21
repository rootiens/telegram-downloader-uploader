package utils

import (
	"errors"
	"github.com/cavaliergopher/grab/v3"
	"log"
	"os"
)

func CreateDownloadDir() {
	path := "files"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println("\033[31m", err, "\033[0m")
		}
	}
}

func DeleteFileFromDisk(filepath string) {
	err := os.Remove(filepath)
	if err != nil {
		log.Println("\033[31m", err, "\033[0m")
	}
}

func DownloadFileFromURL(RequestedURL string) (message, method, filename string) {
	row := GetDB().QueryRow("select * from files where link = ?", RequestedURL)

	var id int64
	var link string
	var fileid string
	var expire_at string
	row.Scan(&id, &link, &fileid, &expire_at)

	if id == 0 {
		resp, err := grab.Get("files", RequestedURL)
		if err != nil {
			log.Println("\033[31m", err, "\033[0m")
			message = "error"
			method = ""
			filename = ""
		} else {
			message = "فایل دانلود شد و به زودی برای شما ارسال میشود."
			method = "upload"
			filename = resp.Filename
		}
	} else {
		message = "فایل دانلود شد و به زودی برای شما ارسال میشود."
		method = "forward"
		filename = fileid
	}
	return message, method, filename
}
