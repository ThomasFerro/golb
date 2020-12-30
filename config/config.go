package config

import (
	"log"
	"os"
)

type ConfigurationKey string

var ConfigurationsWithFallback = map[ConfigurationKey]string{
	"POST_PATH":               "/Users/thomasferro/Documents/perso/git/readmes/posts",
	"TITLE":                   "Thomas Ferro",
	"LOCALE":                  "en",
	"POST_PAGE_TEMPLATE_PATH": "/Users/thomasferro/Documents/perso/git/golb/blog/postPageTemplate.go.html",
	"DIST_PATH":               "./dist",
}

func GetConfiguration(configurationKey ConfigurationKey) string {
	fallback := ConfigurationsWithFallback[configurationKey]
	configurationFromEnv := os.Getenv(string(configurationKey))
	if configurationFromEnv == "" {
		configurationFromEnv = fallback
		log.Printf(
			"No configuration found in env variables for %v, falling back to %v",
			configurationKey,
			fallback,
		)
	}
	return configurationFromEnv
}
