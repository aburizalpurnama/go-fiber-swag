package domain

import "github.com/google/uuid"

type Merchant struct {
	Id                                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" excel:"Id"`
	Email                             string    `gorm:"unique;not null" excel:"Email"`
	MdrPercentage                     float64   `json:"mdr_pcnt" excel:"MDR Percentage(%)"`
	InterestSharingBankPercentage     float64   `json:"interest_sharing_bank_pcnt" excel:"Bank Sharing Percentage(%)"`
	InterestSharingMerchantPercentage float64   `json:"interest_sharing_merchant_pcnt" excel:"Merchant Sharing Percentage(%)"`
}
