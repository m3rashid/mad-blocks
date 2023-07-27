package block

type TransactionRequest struct {
	SenderAddress    *string `json:"senderAddress"`
	RecipientAddress *string `json:"recipientAddress"`
	SenderPublicKey  *string `json:"senderPublicKey"`
	Signature        *string `json:"signature"`
	Value            *string `json:"value"`
}

func (t *TransactionRequest) Validate() bool {
	if t.SenderAddress == nil ||
		t.Value == nil ||
		t.RecipientAddress == nil ||
		t.SenderPublicKey == nil ||
		t.Signature == nil {
		return false
	}

	return true
}
