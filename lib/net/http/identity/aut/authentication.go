package aut

// Claim claim
type Claim map[ClaimType]string

// ClaimType claim type
type ClaimType string

const (
	// NameIdentity identity
	NameIdentity ClaimType = "github.com/picapica360/nameidentity"
	// Role role
	Role ClaimType = "github.com/picapica360/role"
	// Email email
	Email ClaimType = "github.com/picapica360/email"
	// Phone phone
	Phone ClaimType = "github.com/picapica360/phone"
)
