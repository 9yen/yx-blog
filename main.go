package main

import (
	_ "yx-blog/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"yx-blog/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
