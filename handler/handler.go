package handler

import (
	"doodocs/service"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	conf = fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
				"err": http.StatusText(http.StatusInternalServerError),
			})
		},
	}
)

type Handler struct {
	infoLogger *log.Logger
	errLogger *log.Logger
	service *service.Service
	app *fiber.App
}

func NewHandler(info, err *log.Logger, service *service.Service) *Handler {
	return &Handler{
		infoLogger: info,
		errLogger: err,
		service: service,
	}
}

func (h *Handler) Server() *fiber.App {
	h.app = fiber.New(conf)

	h.app.Use(recover.New(), h.LogRequest)

	api := h.app.Group("/api")
	
	arch := api.Group("/archive")
	arch.Post("/information", h.AcceptZip)
	arch.Post("/files", h.AcceptFiles)

	mail := api.Group("/mail")
	mail.Post("/file", h.Emailer)
	
	return h.app
}

func (h *Handler) AcceptZip(c *fiber.Ctx) error{
	file, err := c.FormFile("file")
	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("No file was submitted.")
	}

	if !isZip(file.Header.Get("Content-Type")) {
		return c.Status(fiber.StatusBadRequest).JSON(CustomError{
			Error: ErrInvalidFileFormat,
			Message: "Please submit a zip file.",
		})
	}

	zp, err := h.service.AcceptZip(file)
	if err != nil{
		h.Error(err)
		return err
	}

	return c.JSON(zp)	
}

func (h *Handler) AcceptFiles(c *fiber.Ctx) error{
	form, err := c.MultipartForm()
    if err != nil {
        return err
    }

	files := form.File["file"]

	wrongFiles := []string{}	
	for _, file := range files{
		if !allowToArchive(file.Header.Get("Content-Type"),){
			wrongFiles = append(wrongFiles, file.Filename)
		}
	}
	if len(wrongFiles) > 0{
		return c.Status(fiber.StatusBadRequest).JSON(CustomError{
			Error: ErrInvalidFileFormat,
			Message: fmt.Sprintf("These are incorrect files %v", wrongFiles),
		})
	}

	buf, err := h.service.CreateZip(files)
	if err != nil{
		h.Error(err)
		return err
	}

	c.Set("Content-Disposition", "attachment; filename=files.zip")
	c.Set("Content-Type", "application/zip")

	return c.Send(buf.Bytes())
}

func (h *Handler) Emailer(c *fiber.Ctx) error{
	file, err := c.FormFile("file")
	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("No file was submitted.")
	}

	emails := strings.Split(c.FormValue("emails"), ",")
	if len(emails) == 0{
		return c.Status(fiber.StatusBadRequest).SendString("No emails were submitted.")
	}

	if !allowToSend(file.Header.Get("Content-Type")){
		return c.Status(fiber.StatusBadRequest).JSON(CustomError{
			Error: ErrInvalidFileFormat,
			Message: "Please submit file in correct format.",
		})
	}
	
	validEmails := []string{}
	wrongEmails := []string{}

	for _, email := range emails {
		if !checkEmail(email){
			wrongEmails = append(wrongEmails, email)
		}else{
			validEmails = append(validEmails, email)
		}
	}

	err = h.service.SendFile(file, validEmails)
	if err != nil{
		h.Error(err)
		return err
	}

	if len(wrongEmails)>0{
		return c.JSON(fiber.Map{
			"message": "There are some incorrect emails.",
			"emails": wrongEmails,
		})
	}

	return c.SendString("File sent successfully to all emails!")
} 