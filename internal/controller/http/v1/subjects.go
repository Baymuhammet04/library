package controller

import (
	"errors"
	"log"
	"strconv"

	"github.com/2004942/library/internal/controller/http/v1/request"
	"github.com/2004942/library/internal/controller/http/v1/response"
	"github.com/2004942/library/internal/domain"
	"github.com/2004942/library/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

var _ SubjectController = (*subjectController)(nil)

type SubjectController interface {
	CreateSubject(c *fiber.Ctx) error
	UpdateSubject(c *fiber.Ctx) error
}

type subjectController struct {
	UC usecase.SubjectUC
}

func NewSubjectUC(UC usecase.SubjectUC) *subjectController {
	return &subjectController{UC: UC}
}

func (h *subjectController) CreateSubject(c *fiber.Ctx) error {
	var req request.CreateSubjectsReq
	err := c.BodyParser(&req)
	if err != nil {
		log.Printf("(c.BodyParser): %v", err)
		return c.SendStatus(400)
	}

	ctx := c.Context()

	subject := domain.Subjects{
		NameTk: req.NameTk,
		NameEn: req.NameEn,
		NameRu: req.NameRu,
	}

	id, err := h.UC.Create(ctx,subject)
	if err != nil {
		log.Printf("(h.uc.Create): %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	response := response.CreateSubjectsRes{
		ID: id,
	}

	return c.Status(200).JSON(response)
}

func (h *subjectController) UpdateSubject(c *fiber.Ctx) error{
	subjectIDStr := c.Params("subject_id")
	subjectID, err :=strconv.Atoi(subjectIDStr)
	if err!= nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid subject id",
		})
	}

	var req request.CreateSubjectsReq
	if err := c.BodyParser(&req); err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"errors": "invalid request body",
		})
	}
	
	subject := domain.Subjects{
		ID: subjectID,
		NameTk: req.NameTk,
		NameEn: req.NameEn,
		NameRu: req.NameRu,
	}

	err=h.UC.Update(c.Context(), subject)
	if err!=nil {
		if errors.Is(err, domain.ErrorSubjectsNotFound){
			return c.Status(404).JSON(fiber.Map{
				"message": "subject not found",
			})
		}

		log.Printf("internal server errors: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
