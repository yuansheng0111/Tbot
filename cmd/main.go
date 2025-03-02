package main

import (
	"github.com/yuansheng0111/Tbot/internal/browser"
	"github.com/yuansheng0111/Tbot/internal/config"
)

func main() {
	cfg := config.Load_sysinfo()

	browser.New(cfg)
}
