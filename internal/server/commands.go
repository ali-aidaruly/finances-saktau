package server

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	start                 = "/start"
	addCategoryCommand    = "/add-category"
	addInvoiceCommand     = "/add-invoice"
	listCategoriesCommand = "/list-categories"
)

func (s *Server) processMessage(ctx context.Context, msg tgbotapi.Message) string {
	text := strings.TrimSpace(msg.Text)
	if text == "" {
		return ""
	}

	//msg.IsCommand()
	command := getCommand(msg.Text)
	fmt.Println(msg.Text)
	msg.Text = msg.Text[len(command):]
	fmt.Println(command, msg.Text)

	switch command {
	case start:
		return s.createUser(ctx, msg)
	case addCategoryCommand:
		return s.createCategory(ctx, msg)
	case addInvoiceCommand:
		return s.createInvoice(ctx, msg)
	case listCategoriesCommand:
		return s.getAllCategories(ctx, msg)
	default:
	}

	return ""
}

func getCommand(msg string) string {
	words := strings.Split(msg, " ")
	if len(words) == 0 {
		return ""
	}

	return strings.ToLower(words[0])
}
