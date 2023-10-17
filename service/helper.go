package service

import (
	"archive/zip"
	"bytes"
	"doodocs/model"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/jordan-wright/email"
)

func readFileData(file *multipart.FileHeader) ([]byte, error) {
	zp, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer zp.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, zp)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func processZip(zp *model.ZipFile, zipData []byte) error{

	zipReader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return err
	}
	for _, v := range zipReader.File{
		file := &model.File{}

		vOpen, err := v.Open()
		if err != nil{
			return err
		}
		defer vOpen.Close()

		content, err := io.ReadAll(vOpen)
		if err != nil{
			return err
		}

		file.Path = v.Name
		file.Size = float64(v.FileInfo().Size())
		file.Type = http.DetectContentType(content)
		
		zp.TotalSize += float64(v.FileInfo().Size())
		zp.Files = append(zp.Files, file)
		zp.FilesNum += 1
	}
	
	return nil
}

func addFile(zw *zip.Writer, file *multipart.FileHeader) error {
	fileWriter, err := zw.Create(file.Filename)
    if err != nil {
        log.Fatal(err)
    }

	content, err := readFileData(file)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, bytes.NewReader(content))
	if err != nil {
		return err
	}

	return nil
}

func composeEmail(emails []string, content []byte, filename string) (*email.Email, error) {
	e := email.NewEmail()
    e.From = os.Getenv("EMAIL")
    e.To = emails
    e.Subject = "The file"
    _, err := e.Attach(bytes.NewReader(content), filename, "application/octet-stream")
	if err!=nil{
		return nil, err
	}
    return e, nil
}