package model

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Email        string `gorm:"not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"password_hash,omitempty"`
}

// authenticationRequest mendefinisikan struktur permintaan untuk pendaftaran dan login.
type AuthenticationRequest struct {
	Email    string `json:"email" validate:"required,email"`    // Email harus berupa format email yang valid
	Password string `json:"password" validate:"required,min=6"` // Password harus memiliki panjang minimal 6 karakter
}
