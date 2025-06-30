package handler

import (
	"go-25-27/service"
	"go-25-27/utils"
)

type Handler struct {
	AssignmentHandler AssignmentHandler
	AuthHandler       AuthHandler
	SubmissionHandler SubmissionHandler
}

func NewHandler(service service.Service, config utils.Configuration) Handler {
	return Handler{
		AssignmentHandler: NewAssignmentHandler(service, config),
		AuthHandler:       NewAuthHandler(service),
		SubmissionHandler: NewSubmissionHandler(service),
	}
}
