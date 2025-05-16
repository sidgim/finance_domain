package domain

// AccountType representa la tabla tipo_cuenta
type AccountType struct {
	ID   string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique_index"`
	Name string `json:"name" gorm:"not null"` // Nombre del tipo de cuenta
}
