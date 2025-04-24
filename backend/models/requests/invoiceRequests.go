package requests

type CreateInvoiceRequest struct {
	UserID      string  `json:"user_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type UpdateInvoiceRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}