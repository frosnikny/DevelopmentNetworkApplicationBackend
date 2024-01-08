package schemes

import (
	"awesomeProject/internal/app/ds"
)

type AllDevelopmentServicesResponse struct {
	DevelopmentServices []ds.DevelopmentService `json:"development_services"`
}

type CustomerRequests struct {
	UUID                     string `json:"uuid"`
	DevelopmentServicesCount int    `json:"development_services_count"`
}

type GetAllDevelopmentServicesResponse struct {
	DraftCustomerRequest *CustomerRequests       `json:"draft_customer_request"`
	DevelopmentServices  []ds.DevelopmentService `json:"development_services"`
}

type AllCustomerRequestsResponse struct {
	CustomerRequests []AllCustomerRequestOutputResponse `json:"customer_requests"`
}

type CustomerRequestResponse struct {
	CustomerRequest CustomerRequestOutputResponse `json:"customer_request"`
}

type UpdateCustomerRequestResponse struct {
	CustomerRequest CustomerRequestOutputResponse `json:"customer_request"`
}

type AllCustomerRequestOutputResponse struct {
	UUID              string  `json:"uuid"`
	RecordStatus      uint    `json:"record_status"`
	CreationDate      string  `json:"creation_date"`
	FormationDate     *string `json:"formation_date"`
	CompletionDate    *string `json:"completion_date"`
	WorkSpecification string  `json:"work_specification"`
	Moderator         *string `json:"moderator"`
	Creator           string  `json:"creator"`
	PaymentStatus     *string `json:"payment_status"`
}

type CustomerRequestOutputResponse struct {
	UUID                string                   `json:"uuid"`
	RecordStatus        uint                     `json:"record_status"`
	CreationDate        string                   `json:"creation_date"`
	FormationDate       *string                  `json:"formation_date"`
	CompletionDate      *string                  `json:"completion_date"`
	WorkSpecification   string                   `json:"work_specification"`
	Moderator           *string                  `json:"moderator"`
	Creator             string                   `json:"creator"`
	PaymentStatus       *string                  `json:"payment_status"`
	ServiceRequests     []ServiceRequestResponse `json:"service_requests"`
	DevelopmentServices []ds.DevelopmentService  `json:"development_services"`
}

type ServiceRequestResponse struct {
	DevelopmentServiceId string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"development_service_id"`
	WorkScope            string `gorm:"type:text"`
	WorkingDays          uint   `gorm:"type:integer"`
}

func ConvertCustomerRequestResponse(customerRequest *ds.CustomerRequest, serviceRequests []ds.ServiceRequest, developmentServices []ds.DevelopmentService) CustomerRequestOutputResponse {
	var serviceRequestsResponse = make([]ServiceRequestResponse, len(serviceRequests))
	for i, serviceRequest := range serviceRequests {
		serviceRequestsResponse[i].DevelopmentServiceId = serviceRequest.DevelopmentServiceId
		serviceRequestsResponse[i].WorkScope = serviceRequest.WorkScope
		serviceRequestsResponse[i].WorkingDays = serviceRequest.WorkingDays
	}

	var developmentServicesResponse = make([]ds.DevelopmentService, len(developmentServices))
	for i, developmentService := range developmentServices {
		developmentServicesResponse[i].UUID = developmentService.UUID
		developmentServicesResponse[i].Title = developmentService.Title
		developmentServicesResponse[i].Description = developmentService.Description
		developmentServicesResponse[i].ImageUrl = developmentService.ImageUrl
		developmentServicesResponse[i].Price = developmentService.Price
		developmentServicesResponse[i].RecordStatus = developmentService.RecordStatus
		developmentServicesResponse[i].Technology = developmentService.Technology
		developmentServicesResponse[i].DetailedPrice = developmentService.DetailedPrice
	}

	output := CustomerRequestOutputResponse{
		UUID:                customerRequest.UUID,
		RecordStatus:        customerRequest.RecordStatus,
		CreationDate:        customerRequest.CreationDate.Format("2006-01-02 15:04:05"),
		WorkSpecification:   customerRequest.WorkSpecification,
		Creator:             customerRequest.Creator.Login,
		PaymentStatus:       customerRequest.PaymentStatus,
		ServiceRequests:     serviceRequestsResponse,
		DevelopmentServices: developmentServicesResponse,
	}

	if !customerRequest.FormationDate.IsZero() { // != nil
		formationDate := customerRequest.FormationDate.Format("2006-01-02 15:04:05")
		output.FormationDate = &formationDate
	}

	if !customerRequest.CompletionDate.IsZero() { // != nil
		completionDate := customerRequest.CompletionDate.Format("2006-01-02 15:04:05")
		output.CompletionDate = &completionDate
	}

	if customerRequest.Moderator != nil {
		output.Moderator = &customerRequest.Moderator.Login
	}

	return output
}

func ConvertAllCustomerRequestResponse(customerRequest *ds.CustomerRequest) AllCustomerRequestOutputResponse {
	output := AllCustomerRequestOutputResponse{
		UUID:              customerRequest.UUID,
		RecordStatus:      customerRequest.RecordStatus,
		CreationDate:      customerRequest.CreationDate.Format("2006-01-02 15:04:05"),
		WorkSpecification: customerRequest.WorkSpecification,
		Creator:           customerRequest.Creator.Login,
		PaymentStatus:     customerRequest.PaymentStatus,
	}

	if !customerRequest.FormationDate.IsZero() { // != nil
		formationDate := customerRequest.FormationDate.Format("2006-01-02 15:04:05")
		output.FormationDate = &formationDate
	}

	if !customerRequest.CompletionDate.IsZero() { // != nil
		completionDate := customerRequest.CompletionDate.Format("2006-01-02 15:04:05")
		output.CompletionDate = &completionDate
	}

	if customerRequest.Moderator != nil {
		output.Moderator = &customerRequest.Moderator.Login
	}

	return output
}

type AuthResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
