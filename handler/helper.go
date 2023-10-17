package handler

import (
	"net/mail"
)

func isZip(mime string) bool {
	return mime == "application/zip"
}

func allowToArchive(mime string) bool{
	return allowedToSave[mime]
}

func allowToSend(mime string) bool{
	return allowedToSend[mime]
}

func checkEmail(email string) bool{
	_, err := mail.ParseAddress(email)
    return err == nil
}

var (allowedToSave map[string]bool = map[string]bool{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	"application/xml": true,
	"image/jpeg" : true,
	"image/png": true,
	};
	
	allowedToSend map[string]bool = map[string]bool{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	"application/pdf": true,
	}																										
)

