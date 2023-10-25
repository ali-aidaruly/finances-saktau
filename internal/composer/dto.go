package composer

type CreateInvoice struct {
	UserTelegramId int
	Category       string
	Amount         string
	Currency       *string
}
