package config

import "strings"

type Options struct {
	ConfigName     string
	ConfigType     string
	ConfigPaths    []string
	EnvFiles       []string
	EnvPrefix      string
	EnvKeyReplacer *strings.Replacer
	IgnoreNotFound bool
	Silent         bool
}

func DefaultOptions() *Options {
	return &Options{
		ConfigName:     "default",
		ConfigType:     "yaml",
		ConfigPaths:    []string{"configs", "."},
		EnvFiles:       []string{".env"},
		EnvKeyReplacer: strings.NewReplacer(".", "_"),
		IgnoreNotFound: false,
		Silent:         false,
	}
}
