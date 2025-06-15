// internal/models/employee.go
package models

import "time"

type Employee struct {
    EmployeeID   uint      `gorm:"primaryKey;column:employee_id"`
    FullName     string    `gorm:"column:full_name;size:255;not null"`
    Position     *string   `gorm:"column:position;size:100"`
    Role         string    `gorm:"column:role;size:50;not null"`    // e.g. "viewer", "editor", "approver"
    Username     string    `gorm:"column:username;size:100;unique;not null"`
    PasswordHash string    `gorm:"column:password_hash;size:255;not null"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
