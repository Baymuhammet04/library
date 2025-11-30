package controller

import "github.com/gofiber/fiber/v2"

func MapRoutes(r fiber.Router, sh SubjectController,){
	r.Post("/subjects", sh.CreateSubject)
	r.Put("/subjects/subject_id", sh.UpdateSubject)
}