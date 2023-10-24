package model

type (
	CreateMerchantReq struct {
		Name                              string `json:"name"`
		Email                             string `json:"email"`
		MdrPercentage                     string `json:"mdr_pcnt"`
		InterestSharingBankPercentage     string `json:"interest_sharing_bank_pcnt"`
		InterestSharingMerchantPercentage string `json:"interest_sharing_merchant_pcnt"`
	}
)
