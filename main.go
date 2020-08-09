package main

import (
	"context"
	"lake/commands"
	"lake/handlers"

	"github.com/olivetree123/polar/server"
)

func main() {
	ctx := context.Background()
	s := server.NewServer("localhost:5000")
	s.Register(commands.SetKey, handlers.SetKeyValue)
	s.Register(commands.POP, handlers.SetKeyValue)
	s.Register(commands.PUSH, handlers.SetKeyValue)
	s.Listen(ctx)
}
