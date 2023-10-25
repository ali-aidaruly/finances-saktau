package server

import (
	"context"

	"github.com/ali-aidaruly/finances-saktau/internal/composer"
	"github.com/ali-aidaruly/finances-saktau/internal/telegram"
)

type Server struct {
	requester requester
	responser responser

	composer composer.Composer
}

type requester interface {
	GetRequest() telegram.Request
}

type responser interface {
	Respond(response telegram.Response)
}

func NewServer(
	composer composer.Composer,

	req requester,
	res responser,
) *Server {
	return &Server{
		composer:  composer,
		requester: req,
		responser: res,
	}
}

func (s *Server) Run(ctx context.Context) {
	for {
		req := s.requester.GetRequest()
		s.processRequest(ctx, req)
	}
}

func (s *Server) processRequest(ctx context.Context, req telegram.Request) {
	msg := req.Message
	if msg == nil {
		return
	}

	respMsg := s.processMessage(ctx, *msg)
	//logrus.Info("some m")
	//
	//isCommand := msg.IsCommand()
	//command := msg.Command()
	//
	//logrus.WithFields(logrus.Fields{
	//	"command":   command,
	//	"msg":       msg.Text,
	//	"isCommand": isCommand,
	//}).Info()
	//log.Printf("[%s] %s", msg.From.UserName, msg.Text)
	//
	//fmt.Printf("%+v \n", msg.Chat)
	//
	//resp := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	//resp.ReplyToMessageID = msg.MessageID

	s.responser.Respond(telegram.Response{
		ChatID: int(msg.Chat.ID),
		Text:   respMsg,
	})
}
