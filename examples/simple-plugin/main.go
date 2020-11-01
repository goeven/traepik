package main

import (
	"context"
	"log"
	"net/http"

	"github.com/goeven/traepik/v2/pkg/cmd"
	"github.com/goeven/traepik/v2/pkg/plugins"
)

func main() {
	traefikCmd, err := cmd.New(getPlugins())
	if err != nil {
		log.Fatalf("Failed creating command traefik: %v", err)
	}

	if err := traefikCmd.Execute(); err != nil {
		log.Fatalf("traefik: %v", err)
	}
}

func getPlugins() map[string]plugins.PluginBuilder {
	return map[string]plugins.PluginBuilder{
		"loggingPlugin": plugins.PluginBuilderFunc(func(config map[string]interface{}, middlewareName string) (plugins.Constructor, error) {
			msg, ok := config["log-message"]
			if !ok {
				msg = "hello from the logging plugin"
			}
			return func(ctx context.Context, next http.Handler) (http.Handler, error) {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					log.Printf("traepik: %s", msg)
					next.ServeHTTP(w, r)
				}), nil
			}, nil
		}),
	}
}
