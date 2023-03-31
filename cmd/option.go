package cmd

import (
	"fmt"
	"os"
)

type Option struct {
	APIKey         string
	Translate      string
	Polish         bool
	Debug          bool
	Check          bool
	ConfigSavePath string
	Limit          int
	History        bool
	SessionList    bool
	Session        string
}

func (o *Option) DBPath() string {
	if err := os.MkdirAll(o.ConfigSavePath, 0755); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/context.db", o.ConfigSavePath)
}
