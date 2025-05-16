package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Account representa la tabla cuenta
type Account struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null;unique_index"`
	NameAccount    string    `json:"name_account" gorm:"type:varchar(100);not null"`
	CurrentBalance float64   `json:"current_balance" gorm:"type:decimal(10,2);default:0"`

	UserID uuid.UUID `json:"user_id,omitempty" gorm:"not null"` // Clave foránea para User
	User   *User     `json:"user,omitempty"`

	AccountTypeID uuid.UUID    `json:"account_type_id" gorm:"not null"` // Clave foránea para AccountType
	AccountType   *AccountType `json:"account_type,omitempty"`          // Relación pertenece a AccountType

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // Para soft delete

}
