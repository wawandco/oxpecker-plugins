package migration

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/wawandco/oxplugins/tools/migration/creator"
)

// Generator allows to identify model as a plugin
type Generator struct{}

// Name returns the name of the generator plugin
func (g Generator) Name() string {
	return "migration"
}

// Generate generates an empty [name].plush.html file
func (g Generator) Generate(ctx context.Context, root string, args []string) error {
	example := "please use `ox generate migration [type] [name] [columns?]`"
	if len(args) < 3 {
		return errors.Errorf("no type specified, %s", example)
	}

	if len(args) < 4 {
		return errors.Errorf("no name specified, %s", example)
	}

	dirPath := filepath.Join(root, "migrations")
	if !g.exists(dirPath) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	}

	creator, err := creator.CreateMigrationFor(strings.ToLower(args[2]))
	if err != nil {
		return err
	}

	if err = creator.Create(dirPath, args[3:]); err != nil {
		return errors.Wrap(err, "failed creating migrations")
	}

	return nil
}

func (g Generator) exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}