package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	start                   = "/start"
	addCategoryCommand      = "/addcategory"
	addInvoiceCommand       = "/add"
	addSubscriptionCommand  = "/addsub"
	getCategoriesCommand    = "/getcategories"
	getInvoicesCommand      = "/get"
	getReportCommand        = "/getreport"
	getSubscriptionsCommand = "/getsubs"
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
	case addSubscriptionCommand:
		return s.createSubscription(ctx, msg)
	case getCategoriesCommand:
		return s.getAllCategories(ctx, msg)
	case getInvoicesCommand:
		return s.getInvoices(ctx, msg)
	case getReportCommand:
		return s.getReport(ctx, msg)
	case getSubscriptionsCommand:
		return s.getAllSubscriptions(ctx, msg)
	default:
		zerolog.Ctx(ctx).Log().Str("invalid command!", command).Send()
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
