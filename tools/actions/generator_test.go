package action

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func Test_ActionGenerator(t *testing.T) {
	g := Generator{}

	t.Run("generate action", func(t *testing.T) {
		dir := t.TempDir()
		modelsPath := filepath.Join(dir, "app", "actions")
		templatePath := filepath.Join(dir, "app", "templates")
		if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.Generate(context.Background(), dir, []string{"generate", "action", "users"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		if !g.exists(filepath.Join(modelsPath, "user.go")) {
			t.Error("'user.go' file does not exists on the path")
		}

		if !g.exists(filepath.Join(modelsPath, "user_test.go")) {
			t.Error("'user_test.go' file does not exists on the path")
		}

		if !g.exists(filepath.Join(templatePath, "user.plush.html")) {
			t.Error("'user.plush.html' file does not exists on the path")
		}
	})
	t.Run("generate action and checking the content", func(t *testing.T) {
		dir := t.TempDir()
		modelsPath := filepath.Join(dir, "app", "actions")
		templatePath := filepath.Join(dir, "app", "templates")
		if err := os.MkdirAll(modelsPath, os.ModePerm); err != nil {
			t.Errorf("creating templates folder should not be error, but got %v", err)
		}

		if err := g.Generate(context.Background(), dir, []string{"generate", "action", "users"}); err != nil {
			t.Errorf("should not be error, but got %v", err)
		}

		// Validating Files existence
		if !g.exists(filepath.Join(modelsPath, "user.go")) {
			t.Error("'user.go' file does not exists on the path")
		}

		if !g.exists(filepath.Join(modelsPath, "user_test.go")) {
			t.Error("'user_test.go' file does not exists on the path")
		}

		if !g.exists(filepath.Join(templatePath, "user.plush.html")) {
			t.Error("'user.plush.html' file does not exists on the path")
		}
		content, err := ioutil.ReadFile(filepath.Join(modelsPath, "user.go"))
		if err != nil {
			log.Fatal(err)
		}
		text := string(content)
		matched, err := regexp.MatchString(`func User`, text)

		if !matched {
			fmt.Println(text)
			t.Fatalf("File's content is not correct, %v", err)
		}
	})
}
