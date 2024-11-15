package frpc

import (
	"embed"

	"github.com/LucasLan666666/frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
