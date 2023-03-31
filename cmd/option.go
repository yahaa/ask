package cmd

import "fmt"

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
	return fmt.Sprintf("%s/context.db", o.ConfigSavePath)
}
