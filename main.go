package main

import (
	_ "athena-server/internal/packed"

	_ "athena-server/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"athena-server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
