package config

import "github.com/vinoMamba/lazy-doc-end/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"port": config.Env("APP_PORT", "2048"),
		}
	})
}
