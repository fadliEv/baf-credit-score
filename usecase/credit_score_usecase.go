package usecase

import (
	"baf-credit-score/model"
	"baf-credit-score/model/dto"
	"baf-credit-score/repository"
	"baf-credit-score/utils/common"
	"baf-credit-score/utils/constant"
	"math"
)

type CreditScoreUsecase interface {
	CreateOrUpdateScore(customerId string, credits []dto.CreditResponseDto) error
	FindScoreByCustomer(customerId string) (dto.CreditScoreResponseDto, error)
}

type creditScoreUsecase struct {
	repo            repository.CreditScoreRepository
}

func (c *creditScoreUsecase) CreateOrUpdateScore(customerId string, credits []dto.CreditResponseDto) error {
	// customer, err := c.creditUsecase.GetCreditsByCustomer(customerId)
	// if err != nil {
	// 	return err
	// }

	// calculate credit statics
	totalCredits := len(credits)
	approvedCredits := c.countApprovedCredits(credits)
	rejectedCredits := c.countRejectedCredits(credits)
	score := c.calculateScore(totalCredits, approvedCredits, rejectedCredits)
	grade := c.calculateGrade(score)

	creditScore := model.CreditScore{
		CustomerID:      customerId,
		Score:           score,
		Grade:           grade,
		TotalCredits:    totalCredits,
		ApprovedCredits: approvedCredits,
		RejectedCredits: rejectedCredits,
	}

	existingScore, err := c.repo.GetByCustomer(customerId)
	if err != nil && err.Error() != "record not found" {
		return err
	}
	if existingScore.ID == "" {
		return c.repo.Save(creditScore)
	} else {
		creditScore.ID = existingScore.ID
		return c.repo.Update(creditScore)
	}
}

// Helper: Count approved credits
func (c *creditScoreUsecase) countApprovedCredits(credits []dto.CreditResponseDto) int {
	count := 0
	for _, credit := range credits {
		if credit.Status == constant.APPROVED {
			count++
		}
	}
	return count
}

func (c *creditScoreUsecase) countRejectedCredits(credits []dto.CreditResponseDto) int {
	count := 0
	for _, credit := range credits {
		if credit.Status == constant.REJECTED {
			count++
		}
	}
	return count
}

func (c *creditScoreUsecase) calculateScore(totalCredits, approvedCredits, rejectedCredits int) int {
	if totalCredits == 0 {
		return 0
	}
	approvalRate := float64(approvedCredits) / float64(totalCredits)
	rejectionRate := float64(rejectedCredits) / float64(totalCredits)

	score := int(math.Min(100, approvalRate*80-rejectionRate*20))
	return score
}

func (c *creditScoreUsecase) calculateGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 75:
		return "B"
	case score >= 60:
		return "C"
	case score >= 40:
		return "D"
	default:
		return "E"
	}
}

func (c *creditScoreUsecase) FindScoreByCustomer(customerId string) (dto.CreditScoreResponseDto, error) {
	creditScore, err := c.repo.GetByCustomer(customerId)
	if err != nil {
		return dto.CreditScoreResponseDto{}, err
	}
	return c.mappingToResponse(creditScore), nil
}

func (c *creditScoreUsecase) mappingToResponse(payload model.CreditScore) dto.CreditScoreResponseDto {
	return dto.CreditScoreResponseDto{
		BaseModelResponseDto: dto.BaseModelResponseDto{
			Id:        payload.ID,
			CreatedAt: payload.CreatedAt,
			UpdatedAt: payload.UpdatedAt,
		},
		Customer: dto.CustomerResponseDto{
			BaseModelResponseDto: dto.BaseModelResponseDto{
				Id:        payload.Customer.ID,
				CreatedAt: payload.Customer.CreatedAt,
				UpdatedAt: payload.Customer.UpdatedAt,
			},
			FullName:    payload.Customer.FullName,
			PhoneNumber: payload.Customer.PhoneNumber,
			NIK:         payload.Customer.NIK,
			Address:     payload.Customer.Address,
			Status:      payload.Customer.Status,
			BirthDate:   common.FormatDateString(payload.Customer.BirthDate),
			User:        nil,
			Credits:     nil,
		},
		Score:           payload.Score,
		Grade:           payload.Grade,
		TotalCredits:    payload.TotalCredits,
		ApprovedCredits: payload.ApprovedCredits,
		RejectedCredits: payload.RejectedCredits,
	}
}

func NewCreditScoreUsecase(repo repository.CreditScoreRepository) CreditScoreUsecase {
	return &creditScoreUsecase{
		repo: repo,
	}
}
