package config

import (
	"runtime"
)

type Config struct {
	SysInfo       string
	Browser_Type  string
	Browser_Path  string
	URL           string
	Ticket_number string
	Date          string
	Price         string
	Exclude       []string
	Refresh_time  float32
}

func Load_sysinfo() Config {
	cfg := Config{}
	cfg.SysInfo = runtime.GOOS
	cfg.Browser_Type = "Brace"
	cfg.Browser_Path = "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser"
	cfg.URL = "https://tixcraft.com"
	cfg.Ticket_number = "2"
	cfg.Date = "03/07"
	cfg.Price = "4480"
	cfg.Exclude = []string{"輪椅", "身障", "身心 障礙", "Restricted View", "燈柱遮蔽", "視線不完整"}
	cfg.Refresh_time = 0.1

	return cfg
}
