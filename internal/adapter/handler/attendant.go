package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fiap/challenge-gofood/internal/core/port"
	"github.com/gin-gonic/gin"
)

type AttendantHandler struct {
	UseCase port.AttendantUseCasePort
}

func NewAttendantHandler(useCase port.AttendantUseCasePort) *AttendantHandler {
	return &AttendantHandler{useCase}
}

type FindAttendantResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
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

	var response []FindAttendantResponse
	for _, attendant := range attendants {
		response = append(response, FindAttendantResponse{
			ID:   attendant.ID,
			Name: attendant.Name,
		})
	}

	c.JSON(http.StatusOK, response)
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

	response := FindAttendantResponse{
		ID:   attendant.ID,
		Name: attendant.Name,
	}

	c.JSON(http.StatusOK, response)
}

type CreateAttendantRequest struct {
	Name string `json:"name"`
}

type CreateAttendantResponse struct {
	ID uint `json:"id"`
}

func (h *AttendantHandler) CreateAttendant(c *gin.Context) {
	var request CreateAttendantRequest
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

	response := CreateAttendantResponse{
		ID: attendant.ID,
	}

	c.JSON(http.StatusCreated, response)
}

type UpdateAttendantRequest struct {
	Nome string `json:"name"`
}

func (h *AttendantHandler) UpdateAttendant(c *gin.Context) {
	var request UpdateAttendantRequest
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

	_, err = h.UseCase.UpdateAttendant(attendant)

	response := fmt.Sprintf("Attendant[%d] - %s updated", attendant.ID, request.Nome)

	c.JSON(http.StatusAccepted, gin.H{
		"message": response,
	})
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
