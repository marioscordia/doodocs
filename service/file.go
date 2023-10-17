package service

import (
	"archive/zip"
	"bytes"
	"doodocs/model"
	"mime/multipart"
	"net/smtp"
	"os"
)

type File interface {
	AcceptZip(*multipart.FileHeader) (*model.ZipFile, error)
	CreateZip([]*multipart.FileHeader) (*bytes.Buffer, error)
	SendFile(*multipart.FileHeader, []string) error
}

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (f *FileService) AcceptZip(file *multipart.FileHeader) (*model.ZipFile, error){
	
	zipFile := &model.ZipFile{}
	zipFile.Name = file.Filename
	zipFile.ZipSize = float64(file.Size)

	zipData, err := readFileData(file)
	if err != nil {
		return nil, err
	}

	if err := processZip(zipFile, zipData); err != nil {
		return nil, err
	}

	return zipFile, nil
}

func (f *FileService) CreateZip(files []*multipart.FileHeader) (*bytes.Buffer, error){
	
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	defer zipWriter.Close()

	for _, f := range files {
		if err := addFile(zipWriter, f); err != nil{
			return nil, err
		}
	}

	return &buf, nil
}

func (f *FileService) SendFile(file *multipart.FileHeader, emails []string) error{

	content, err := readFileData(file)
	if err != nil {
		return err
	}
	
	email, err := composeEmail(emails, content, file.Filename)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")

	return email.Send("smtp.gmail.com:587", auth)
}