package main

import (
	"github.com/alpha-supsys/go-cm-ucloud-webhook/src/ucloud"
	"github.com/alpha-supsys/go-common/config/env"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/cmd"
)

func main() {
	if cfg, err := env.LoadAllWithoutPrefix("UWH_"); err == nil {
		cmd.RunWebhookServer(cfg.GetString("GROUP_NAME", ""), ucloud.NewSolver(cfg))
	}
}
