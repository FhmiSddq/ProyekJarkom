// Package bootstrap loads the necessary config
package bootstrap

import (
	"log"

	"https://github.com/FhmiSddq/ProyekJarkom/internal/app/chat/handler"
	"https://github.com/FhmiSddq/ProyekJarkom/internal/app/chat/interface/tcp"
	"https://github.com/FhmiSddq/ProyekJarkom/internal/infra/env"
)

func Start() {
	log.Println("starting app")

	config := env.New()

	chatHandler := handler.New(config)

	log.Println("app started")

	tcp.NewChat(config, chatHandler)
}
