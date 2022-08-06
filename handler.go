package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"time"
)

var httpClient = &http.Client{Transport: http.DefaultTransport, Timeout: 60 * time.Second}

type Handler struct {
	ctx    context.Context
	botAPI *tgbotapi.BotAPI
	me     tgbotapi.User
}

func NewHandler(ctx context.Context, botToken string) (h *Handler) {
	var err error
	h = &Handler{ctx: ctx}
	h.botAPI, err = tgbotapi.NewBotAPIWithClient(botToken, tgbotapi.APIEndpoint, httpClient)
	if err != nil {
		log.Fatalf("new bot api error %v", err)
	}
	h.me, err = h.botAPI.GetMe()
	if err != nil {
		log.Fatalf("test bot api error %v", err)
	}
	return h
}

func (h *Handler) Chat(chatId int64) (tgbotapi.Chat, error) {
	return h.botAPI.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: chatId}})
}

func (h *Handler) Admins(chatId int64) ([]tgbotapi.ChatMember, error) {
	return h.botAPI.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: chatId}})
}

func (h *Handler) Member(chatId, userId int64) (tgbotapi.ChatMember, error) {
	return h.botAPI.GetChatMember(tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{ChatID: chatId, UserID: userId},
	})
}

func (h *Handler) MembersCount(chatId int64) (int, error) {
	return h.botAPI.GetChatMembersCount(tgbotapi.ChatMemberCountConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: chatId}})
}

func (h *Handler) Promote(chatId, userId int64, canPromote ...bool) error {
	chatTable := tgbotapi.PromoteChatMemberConfig{
		ChatMemberConfig:    tgbotapi.ChatMemberConfig{ChatID: chatId, UserID: userId},
		IsAnonymous:         true,
		CanManageChat:       true,
		CanChangeInfo:       true,
		CanPostMessages:     true,
		CanEditMessages:     true,
		CanDeleteMessages:   true,
		CanManageVoiceChats: true,
		CanInviteUsers:      true,
		CanRestrictMembers:  true,
		CanPinMessages:      true,
		CanPromoteMembers:   len(canPromote) > 0 && canPromote[0],
	}
	_, err := h.botAPI.Request(chatTable)
	return err
}

func (h *Handler) Leave(chatId int64) error {
	_, err := h.botAPI.Send(tgbotapi.LeaveChatConfig{ChatID: chatId})
	return err
}

func (h *Handler) reply(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := h.botAPI.Send(msg)
	if err != nil {
		log.Printf("reply chat %d, text %s, error %v\n", chatId, text, err)
	}
	return err
}
