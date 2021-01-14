package liquibase

import (
	"context"
	"os"
	"testing"
)

func TestLiquibaseGenerator(t *testing.T) {
	t.Run("FileDoesNotExist", func(t *testing.T) {
		root := t.TempDir()
		err := os.Chdir(root)
		if err != nil {
			t.Error("could not change to temp directory")
		}

		args := []string{"ox", "generate", "liquibase", "addDevices"}

		g := Generator{}

		err = g.Generate(context.Background(), root, args)

		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}

	})

}
