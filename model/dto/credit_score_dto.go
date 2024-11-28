package dto

type CreditScoreResponseDto struct {
	BaseModelResponseDto
	Customer        CustomerResponseDto `json:"customer"`
	Score           int                 `json:"score"`
	Grade           string              `json:"grade"`
	TotalCredits    int                 `json:"totalCredits"`
	ApprovedCredits int                 `json:"approvedCredits"`
	RejectedCredits int                 `json:"rejectedCredits"`
}
