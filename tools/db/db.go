// db packs all db operations under this top level command.
package db

import (
	"context"
	"errors"
	"fmt"

	pop4 "github.com/gobuffalo/pop"
	pop5 "github.com/gobuffalo/pop/v5"
	"github.com/wawandco/oxpecker/plugins"
)

var _ plugins.Command = (*Command)(nil)
var _ plugins.HelpTexter = (*Command)(nil)
var _ plugins.PluginReceiver = (*Command)(nil)
var _ plugins.Subcommander = (*Command)(nil)

var ErrConnectionNotFound = errors.New("connection not found")

type Command struct {
	subcommands []plugins.Command
}

func (c Command) Name() string {
	return "db"
}

func (c Command) ParentName() string {
	return ""
}

func (c Command) HelpText() string {
	return "database operation commands"
}

func (c *Command) Run(ctx context.Context, root string, args []string) error {
	if len(args) < 2 {
		fmt.Println("no subcommand specified, please use `db [subcommand]` to run one of the db subcommands.")
		return nil
	}

	name := args[1]
	var subcommand plugins.Command
	for _, sub := range c.subcommands {
		if sub.Name() != name {
			continue
		}

		subcommand = sub
		break
	}

	if subcommand == nil {
		return fmt.Errorf("subcommand `%v` not found", name)
	}

	return subcommand.Run(ctx, root, args)
}

func (c *Command) Receive(pls []plugins.Plugin) {
	for _, plugin := range pls {
		ptool, ok := plugin.(plugins.Command)
		if !ok || ptool.ParentName() != c.Name() {
			continue
		}

		c.subcommands = append(c.subcommands, ptool)
	}
}

func (c *Command) Subcommands() []plugins.Command {
	return c.subcommands
}

func Plugins(conns interface{}) []plugins.Plugin {
	var result []plugins.Plugin
	providers := map[string]URLProvider{}
	switch v := conns.(type) {
	case map[string]*pop4.Connection:
		for k, conn := range v {
			providers[k] = conn
		}
	case map[string]*pop5.Connection:
		for k, conn := range v {
			providers[k] = conn
		}
	default:
		fmt.Println("DB plugin ONLY receives pop v4 and v5 connections")
		return result
	}

	result = append(result, &Command{})
	result = append(result, &CreateCommand{connections: providers})
	result = append(result, &DropCommand{connections: providers})
	result = append(result, &ResetCommand{connections: providers})

	return result
}
