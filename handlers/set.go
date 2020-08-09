package handlers

import (
	"context"

	"github.com/olivetree123/polar/commands"
	"github.com/olivetree123/polar/response"
	"github.com/olivetree123/polar/server"
)

// SetKeyValue 设置键值对
func SetKeyValue(ctx context.Context, push server.PushHandler, command commands.Command) *response.Response {
	return response.APIResponse(command.RequestID, command.Code, nil)
}
