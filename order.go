package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (h *Handler) botUpdate() {
	log.Println("start receive message...")
	updates := h.botAPI.GetUpdatesChan(tgbotapi.UpdateConfig{Offset: 0, Limit: 0, Timeout: 30})
	for {
		select {
		case u := <-updates:
			if u.Message != nil {
				h.message(u.Message)
			}
		case <-h.ctx.Done():
			log.Println("bot update exit...")
			return
		}
	}
}

func checkAdmin(userId int64) bool {
	if userId == admin {
		return true
	}
	return false
}

func (h *Handler) message(msg *tgbotapi.Message) {
	h.order(msg)
}

func (h *Handler) order(msg *tgbotapi.Message) {
	var err error
	var cmd = msg.Command()
	switch cmd {
	case "id":
		err = h.reply(msg.Chat.ID, fmt.Sprintf("%d", msg.From.ID))
	case "promote":
		if !checkAdmin(msg.From.ID) {
			return
		}
		h.promote(msg)
	default:
		return
	}
	if err != nil {
		log.Printf("handler message ---> command %s, error %v, message %s\n", cmd, err, msg.Text)
	}
}

func (h *Handler) promote(msg *tgbotapi.Message) {
	chatUser := msg.CommandArguments()
	splits := strings.Split(chatUser, "/")
	if len(splits) != 2 {
		h.reply(msg.Chat.ID, "/promote chat-id/user-id")
		return
	}
	chatId, err := strconv.ParseInt(splits[0], 10, 64)
	if err != nil {
		h.reply(msg.Chat.ID, "invalid chat id")
		return
	}
	userId, err := strconv.ParseInt(splits[1], 10, 64)
	if err != nil {
		h.reply(msg.Chat.ID, "invalid user id")
		return
	}
	chat, err := h.Chat(chatId)
	if err != nil {
		h.reply(msg.Chat.ID, fmt.Sprintf("invalid chat, %v", err))
		return
	}
	user, err := h.Member(chat.ID, userId)
	if err != nil {
		h.reply(msg.Chat.ID, fmt.Sprintf("invalid user, %v", err))
		return
	}
	admins, err := h.Admins(chat.ID)
	if err != nil {
		h.reply(msg.Chat.ID, fmt.Sprintf("chat admins error %v", err))
		return
	}
	var canPromote bool
	for _, a := range admins {
		if a.User.ID == h.me.ID {
			if a.CanPromoteMembers {
				canPromote = true
				break
			}
		}
	}
	if !canPromote {
		h.reply(msg.Chat.ID, "bot can not promote chat")
		return
	}
	if err := h.Promote(chat.ID, user.User.ID, promote); err != nil {
		h.reply(msg.Chat.ID, fmt.Sprintf("promote failed %v", err))
		return
	}
	h.reply(msg.Chat.ID, "promote success")
}
