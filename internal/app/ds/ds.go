package ds

import (
	"awesomeProject/internal/app/role"
	"time"
)

type User struct {
	UUID     string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Role     role.Role
	Login    string `gorm:"size:30;not null" json:"login"`
	Password string `gorm:"size:40;not null" json:"-"`
}

type DevelopmentService struct {
	UUID          string  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"uuid"`
	Title         string  `gorm:"size:100"`
	Description   string  `gorm:"type:text"`
	ImageUrl      *string `gorm:"size:100" json:"image_url"`
	Price         float32 `gorm:"type:real"`
	RecordStatus  uint    `gorm:"type:integer"`
	Technology    string  `gorm:"type:text"`
	DetailedPrice float32 `gorm:"type:real"`
}

type CustomerRequest struct {
	UUID              string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	RecordStatus      uint      `gorm:"type:integer"`
	CreationDate      time.Time `gorm:"type:timestamp"`
	FormationDate     time.Time `gorm:"type:timestamp"`
	CompletionDate    time.Time `gorm:"type:timestamp"`
	WorkSpecification string    `gorm:"type:text"`
	CreatorId         string    `gorm:"not null"`
	ModeratorId       *string   `json:"-"`
	PaymentStatus     *string   `gorm:"size:40"`

	Creator   User
	Moderator *User
}

type ServiceRequest struct {
	DevelopmentServiceId string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"development_service_id"`
	CustomerRequestId    string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"customer_request_id"`
	WorkScope            string `gorm:"type:text"`
	WorkingDays          uint   `gorm:"type:integer"`

	DevelopmentService *DevelopmentService `gorm:"foreignKey:DevelopmentServiceId"`
	CustomerRequest    *CustomerRequest    `gorm:"foreignKey:CustomerRequestId"`
}

const (
	DSWorks = iota
	DSDeleted
)

const (
	CRDraft = iota
	CRWorks
	CRCompleted
	CRDeclined
	CRDeleted
)

const PaymentStarted = "2"

const ZeroUUID = "00000000-0000-0000-0000-000000000000"
