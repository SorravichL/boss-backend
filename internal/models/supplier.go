// internal/models/supplier.go
package models

import "time"

type Supplier struct {
    SupplierID     uint       `gorm:"primaryKey;column:supplier_id"`
    SupplierName   string     `gorm:"column:supplier_name;size:255;not null"`
    CreditTerms    *string    `gorm:"column:credit_terms;size:100"`
    ContactChannel *string    `gorm:"column:contact_channel;size:255"`
    CreatedAt      time.Time  `gorm:"autoCreateTime"`
    UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
}
