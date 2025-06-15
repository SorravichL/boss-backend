// internal/handlers/supplier.go
package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "boss-backend/internal/models"
)

// SupplierHandler handles CRUD operations for the Supplier model.
type SupplierHandler struct {
    DB *gorm.DB
}

// NewSupplierHandler initializes a new SupplierHandler.
func NewSupplierHandler(db *gorm.DB) *SupplierHandler {
    return &SupplierHandler{DB: db}
}

// List returns all suppliers in JSON format.
func (h *SupplierHandler) List(c *gin.Context) {
    var list []models.Supplier
    h.DB.Find(&list)
    c.JSON(http.StatusOK, list)
}

// Create adds a new supplier to the database.
func (h *SupplierHandler) Create(c *gin.Context) {
    var sup models.Supplier
    if err := c.ShouldBindJSON(&sup); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.DB.Create(&sup).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, sup)
}

// Get fetches a supplier by its ID.
func (h *SupplierHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var sup models.Supplier
    if err := h.DB.First(&sup, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, sup)
}

// Update modifies fields of an existing supplier.
func (h *SupplierHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var sup models.Supplier
    if err := h.DB.First(&sup, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    var input struct {
        SupplierName   *string `json:"SupplierName"`
        CreditTerms    *string `json:"CreditTerms"`
        ContactChannel *string `json:"ContactChannel"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updates := map[string]interface{}{}
    if input.SupplierName != nil {
        updates["supplier_name"] = *input.SupplierName
    }
    if input.CreditTerms != nil {
        updates["credit_terms"] = input.CreditTerms
    }
    if input.ContactChannel != nil {
        updates["contact_channel"] = input.ContactChannel
    }

    if err := h.DB.Model(&sup).Updates(updates).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, sup)
}

// Delete removes a supplier record by ID.
func (h *SupplierHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.DB.Delete(&models.Supplier{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}