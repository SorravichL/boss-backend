package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "boss-backend/internal/models"
)

type CustomerHandler struct {
    DB *gorm.DB
}

func NewCustomerHandler(db *gorm.DB) *CustomerHandler {
    return &CustomerHandler{DB: db}
}

func (h *CustomerHandler) List(c *gin.Context) {
    var list []models.Customer
    h.DB.Find(&list)
    c.JSON(http.StatusOK, list)
}

func (h *CustomerHandler) Create(c *gin.Context) {
    var cust models.Customer
    if err := c.ShouldBindJSON(&cust); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.DB.Create(&cust).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, cust)
}

func (h *CustomerHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var cust models.Customer
    if err := h.DB.First(&cust, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, cust)
}

func (h *CustomerHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var cust models.Customer
    if err := h.DB.First(&cust, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    var input struct {
        CompanyName *string `json:"CompanyName"`
        ContactName *string `json:"ContactName"`
        Address     *string `json:"Address"`
        Phone       *string `json:"Phone"`
        TaxID       *string `json:"TaxID"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updates := map[string]interface{}{}
    if input.CompanyName != nil {
        updates["company_name"] = *input.CompanyName
    }
    if input.ContactName != nil {
        updates["contact_name"] = input.ContactName
    }
    if input.Address != nil {
        updates["address"] = input.Address
    }
    if input.Phone != nil {
        updates["phone"] = input.Phone
    }
    if input.TaxID != nil {
        updates["tax_id"] = input.TaxID
    }

    if err := h.DB.Model(&cust).Updates(updates).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, cust)
}

func (h *CustomerHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.DB.Delete(&models.Customer{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
