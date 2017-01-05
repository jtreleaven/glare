package glare

// Identity represents a single user resource from the Layer API
type Identity struct {
	DisplayName string                 `json:"display_name"`
	AvatarURL   string                 `json:"avatar_url"`
	FirstName   string                 `json:"first_name"`
	LastName    string                 `json:"last_name"`
	Phone       string                 `json:"phone_number"`
	Email       string                 `json:"email_address"`
	MetaData    map[string]interface{} `json:"metadata"`
}
