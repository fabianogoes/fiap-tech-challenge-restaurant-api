package rest

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
	"github.com/fabianogoes/fiap-challenge/frameworks/rest/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttendantHandler struct {
	UseCase ports.AttendantUseCasePort
}

func NewAttendantHandler(useCase ports.AttendantUseCasePort) *AttendantHandler {
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
		return
	}

	attendant, err := h.UseCase.GetAttendantById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ToAttendantResponse(attendant))
}

func (h *AttendantHandler) CreateAttendant(c *gin.Context) {
	var request dto.CreateAttendantRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	attendant, err := h.UseCase.CreateAttendant(request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
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
		return
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	attendant, err := h.UseCase.GetAttendantById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	attendant.Name = request.Name
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
		return
	}

	err = h.UseCase.DeleteAttendant(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := fmt.Sprintf("Attendant %d deleted", id)

	c.JSON(http.StatusNoContent, gin.H{
		"message": response,
	})
}
