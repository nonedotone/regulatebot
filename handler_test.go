package main

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	botToken         = "123456789:qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop"
	deliverBot       = int64(1234567890)
	deliverChannelId = int64(-1001234567890)
	deliverGroupId   = int64(-1001234567890)
	chatId           = deliverGroupId
)

func TestNewHandler(t *testing.T) {
	ctx := context.Background()
	h := NewHandler(ctx, botToken)

	_, err := h.Chat(chatId)
	require.NoError(t, err)

	_, err = h.MembersCount(chatId)
	require.NoError(t, err)

	_, err = h.Admins(chatId)
	require.NoError(t, err)
	err = h.Promote(chatId, deliverBot)
	require.NoError(t, err)
}
