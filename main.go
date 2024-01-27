package main

import (
	_ "athena-server/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "athena-server/internal/logic/logic.go"


	"github.com/gogf/gf/v2/os/gctx"

	"athena-server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
