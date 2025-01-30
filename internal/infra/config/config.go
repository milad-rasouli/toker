package config

import (
	"encoding/json"
	"fmt"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/milad-rasouli/toker/internal/infra/logger"
	"github.com/milad-rasouli/toker/internal/infra/redis"

	"github.com/tidwall/pretty"
	"log"
	"strings"
)

type app struct {
	Name string `json:"name" koanf:"name"`
	Port string `json:"port" koanf:"port"`
}

const prefix = "toker_"

type Config struct {
	App    app           `json:"app" koanf:"app"`
	Redis  redis.Config  `json:"redis"  koanf:"redis"`
	Logger logger.Config `json:"logger"    koanf:"logger"`
}

func New() (*Config, error) {
	var (
		err error
		k   = koanf.New(".")
	)

	err = k.Load(structs.Provider(Default(), "koanf"), nil)
	if err != nil {
		return nil, fmt.Errorf("error loading default: %w", err)
	}

	err = k.Load(file.Provider("config.toml"), toml.Parser())
	if err != nil {
		return nil, fmt.Errorf("error loading config.toml: %w", err)
	}

	err = k.Load(
		env.Provider(prefix, ".", func(source string) string {
			base := strings.ToLower(strings.TrimPrefix(source, prefix))
			return strings.ReplaceAll(base, "__", ".")
		}),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	var instance Config
	err = k.Unmarshal("", &instance)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config %w", err)
	}

	indent, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		panic(err)
	}

	indent = pretty.Color(indent, nil)

	log.Printf(`
================ Loaded Configuration ================
%s
======================================================
	`, string(indent))

	return &instance, nil
}
