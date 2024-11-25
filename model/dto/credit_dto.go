package dto

type CreditRequestDto struct {
	Id               string  `json:"id"`
	AppNumber        string  `json:"appNumber"`
	CustomerID       string  `json:"customerId"`
	ProductType      string  `json:"productType"`
	LoanAmount       float64 `json:"loanAmount"`
	Tenure           int     `json:"tenure"`
	EmploymentStatus string  `json:"employmentStatus"`
	MonthlyIncome    float64 `json:"monthlyIncome"`
	Status           string  `json:"status"`
	RejectionReason  string  `json:"rejectionReason"`
	SubmittedAt      string  `json:"submittedAt"`
}

type CreditResponseDto struct {
	BaseModelResponseDto
	AppNumber        string              `json:"appNumber"`
	Customer         CustomerResponseDto `json:"customer,omitempty"`
	ProductType      string              `json:"productType"`
	LoanAmount       float64             `json:"loanAmount"`
	Tenure           int                 `json:"tenure"`
	EmploymentStatus string              `json:"employmentStatus"`
	MonthlyIncome    float64             `json:"monthlyIncome"`
	Status           string              `json:"status"`
	RejectionReason  string              `json:"rejectionReason"`
	SubmittedAt      string              `json:"submittedAt"`
}
