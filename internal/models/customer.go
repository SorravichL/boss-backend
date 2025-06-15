package models

import "time"

type Customer struct {
  CustomerID  uint      `gorm:"primaryKey;column:customer_id"`
  CompanyName string    `gorm:"column:company_name;size:255;not null"`
  ContactName *string   `gorm:"column:contact_name;size:255"`
  Address     *string   `gorm:"column:address;type:text"`
  Phone       *string   `gorm:"column:phone;size:50"`
  TaxID       *string   `gorm:"column:tax_id;size:50"`
  CreatedAt   time.Time `gorm:"autoCreateTime"`
  UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
