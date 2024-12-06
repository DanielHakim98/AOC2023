package scaffold

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type scaffold struct {
	debug bool
}

func New(debug bool) *scaffold {
	return &scaffold{debug: debug}
}

func (s *scaffold) Generate(path string, day int) (err error) {
	genAbsDir, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	templateDir := filepath.Join(genAbsDir, "utils/scaffold/template/")

	dayDir := filepath.Join(genAbsDir, fmt.Sprintf("day%v/", day))
	if _, err := os.Stat(dayDir); os.IsNotExist(err) {
		err := os.Mkdir(dayDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	mainFile := filepath.Join(dayDir, fmt.Sprintf("day%v.go", day))
	testFile := filepath.Join(dayDir, fmt.Sprintf("day%v_test.go", day))
	cmdFile := filepath.Join(genAbsDir, fmt.Sprintf("cmd/day%v.go", day))

	// Cleanup if any failure occured
	defer func() {
		if err != nil {
			if _, err := os.Stat(dayDir); !os.IsNotExist(err) {
				if err := os.RemoveAll(dayDir); err != nil {
					log.Println("Error cleaning up:", err)
				}
			}
			if _, err := os.Stat(cmdFile); !os.IsNotExist(err) {
				if err := os.Remove(cmdFile); err != nil {
					log.Println("Error cleaning up:", err)
				}
			}
		}
	}()

	// Generate File
	mainTemplate := filepath.Join(templateDir, "file.tmpl")
	err = s.createFileTemplate(mainTemplate, mainFile, day)
	if err != nil {
		return err
	}

	// Generate Test
	testTemplate := filepath.Join(templateDir, "test.tmpl")
	err = s.createFileTemplate(testTemplate, testFile, day)
	if err != nil {
		return err
	}

	// Generate CMD
	cmdTemplate := filepath.Join(templateDir, "cli.tmpl")
	err = s.createFileTemplate(cmdTemplate, cmdFile, day)
	if err != nil {
		return err
	}

	return nil
}

func (s *scaffold) createFileTemplate(sourceFile string, targetFile string, data int) error {
	t, err := template.ParseFiles(sourceFile)
	if err != nil {
		log.Print(err)
		return err
	}

	f, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, struct {
		Day int
	}{
		Day: data,
	})
	if err != nil {
		return err
	}
	return nil
}
