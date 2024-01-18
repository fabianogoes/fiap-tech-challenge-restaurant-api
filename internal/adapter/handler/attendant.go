package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fiap/challenge-gofood/internal/adapter/handler/dto"
	"github.com/fiap/challenge-gofood/internal/core/port"
	"github.com/gin-gonic/gin"
)

type AttendantHandler struct {
	UseCase port.AttendantUseCasePort
}

func NewAttendantHandler(useCase port.AttendantUseCasePort) *AttendantHandler {
	return &AttendantHandler{useCase}
}

func (h *AttendantHandler) GetAttendants(c *gin.Context) {
	attendants, err := h.UseCase.GetAttendants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if len(attendants) == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "No attendants found",
		})
	}

	c.JSON(http.StatusOK, dto.ToAttendantResponses(attendants))
}

func (h *AttendantHandler) GetAttendant(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	attendant, err := h.UseCase.GetAttendantById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, dto.ToAttendantResponse(attendant))
}

func (h *AttendantHandler) CreateAttendant(c *gin.Context) {
	var request dto.CreateAttendantRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	attendant, err := h.UseCase.CreateAttendant(request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, dto.ToAttendantResponse(attendant))
}

func (h *AttendantHandler) UpdateAttendant(c *gin.Context) {
	var request dto.UpdateAttendantRequest
	var err error

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	attendant, err := h.UseCase.GetAttendantById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	attendant.Name = request.Nome
	attendantUpdated, err := h.UseCase.UpdateAttendant(attendant)

	c.JSON(http.StatusAccepted, dto.ToAttendantResponse(attendantUpdated))
}

func (h *AttendantHandler) DeleteAttendant(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = h.UseCase.DeleteAttendant(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := fmt.Sprintf("Attendant %d deleted", id)
	println(response)

	c.JSON(http.StatusNoContent, gin.H{
		"message": response,
	})
}
