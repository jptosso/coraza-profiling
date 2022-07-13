package main

import (
	"fmt"

	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	"net/http"
	_ "net/http/pprof"

	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/corazawaf/coraza-caddy"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()
	caddycmd.Main()
}
