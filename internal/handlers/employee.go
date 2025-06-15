package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "boss-backend/internal/models"
)

// EmployeeHandler handles CRUD operations for employees

type EmployeeHandler struct {
    DB *gorm.DB
}

func NewEmployeeHandler(db *gorm.DB) *EmployeeHandler {
    return &EmployeeHandler{DB: db}
}

func (h *EmployeeHandler) List(c *gin.Context) {
    var list []models.Employee
    h.DB.Find(&list)
    c.JSON(http.StatusOK, list)
}

func (h *EmployeeHandler) Create(c *gin.Context) {
    var emp models.Employee
    if err := c.ShouldBindJSON(&emp); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.DB.Create(&emp).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, emp)
}

func (h *EmployeeHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var emp models.Employee
    if err := h.DB.First(&emp, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, emp)
}

func (h *EmployeeHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var emp models.Employee
    if err := h.DB.First(&emp, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    var input struct {
        FullName     *string `json:"FullName"`
        Position     *string `json:"Position"`
        Role         *string `json:"Role"`
        Username     *string `json:"Username"`
        PasswordHash *string `json:"PasswordHash"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updates := map[string]interface{}{}
    if input.FullName != nil {
        updates["full_name"] = *input.FullName
    }
    if input.Position != nil {
        updates["position"] = input.Position
    }
    if input.Role != nil {
        updates["role"] = *input.Role
    }
    if input.Username != nil {
        updates["username"] = *input.Username
    }
    if input.PasswordHash != nil {
        updates["password_hash"] = *input.PasswordHash
    }

    if err := h.DB.Model(&emp).Updates(updates).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, emp)
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.DB.Delete(&models.Employee{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
