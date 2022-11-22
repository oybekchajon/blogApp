package models

type UserRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	Email           string `json:"email"`
	Gender          string `json:"gender"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
	Type            string `json:"type"`
	Password        string `json:"password"`
}

type ResponseError struct {
	Error string `json:"error"`
}

type GetAllUsersResponse struct {
	Users []*UserRequest `json:"categories"`
	Count int32          `json:"count"`
}