package shadowos

import (
	"log"
	"testing"
)

var (
	app = &ShadowosApp{
		AddrWs:     "ws://127.0.0.1:8787",
		AddrSocks5: "127.0.0.1:2080",
		UUID:       "53881505-c10c-464a-8949-e57184a576a9",
	}
)

func TestShadowosApp_Run(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.Run()
}
