package wallet

type TransactionRequest struct {
	SenderPrivateKey *string `json:"senderPrivateKey"`
	SenderAddress    *string `json:"senderAddress"`
	RecipientAddress *string `json:"recipientAddress"`
	SenderPublicKey  *string `json:"senderPublicKey"`
	Value            *string `json:"value"`
}

func (tr *TransactionRequest) Validate() bool {
	if tr.SenderAddress == nil ||
		tr.RecipientAddress == nil ||
		tr.Value == nil ||
		tr.SenderPrivateKey == nil ||
		tr.SenderPublicKey == nil {
		return false
	}
	return true
}
