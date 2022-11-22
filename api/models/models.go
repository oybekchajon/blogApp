package models

type UserRequest struct {
	FirstName       string  `json:"first_name" binding:"required,min=2,max=30"`
	LastName        string  `json:"last_name" binding:"required,min=2,max=30"`
	PhoneNumber     *string `json:"phone_number"`
	Email           string  `json:"email" binding:"required,email"`
	Gender          *string `json:"gender" binding:"oneof=male female"`
	Username        string  `json:"username" binding:"required,min=2,max=30"`
	ProfileImageUrl *string `json:"profile_image_url"`
	Type            string  `json:"type" binding:"required,oneof=superadmin user"`
	Password        string  `json:"password" binding:"required,min=6,max=16"`
}