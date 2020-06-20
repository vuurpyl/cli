package cgapp

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
)

// Config struct for app configuration
type Config struct {
	name   string
	match  string
	view   string
	folder string
}

// Create function for create app
func Create(c *Config, registry map[string]string) error {
	// Create path to folder
	folder := filepath.Join(c.folder, c.view)

	// Create match expration for frameworks/containers
	match, err := regexp.MatchString(c.match, c.name)
	ErrChecker(err)

	// Check for regexp
	if match {
		// If match, create from default template
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://github.com/" + registry[c.name],
			Progress: os.Stdout,
		})
		ErrChecker(err)

		// Show success report
		fmt.Printf(
			"\n%v[OK]%v %v was created with default template '%v'!\n",
			green, noColor,
			strings.Title(c.view),
			registry[c.name],
		)
	} else {
		// Else create from user template (from GitHub, etc)
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://" + c.name,
			Progress: os.Stdout,
		})
		ErrChecker(err)

		// Show success report
		fmt.Printf(
			"\n%v[OK]%v %v was created with user template '%v'!\n",
			green, noColor,
			strings.Title(c.view),
			c.name,
		)
	}

	// Clean
	ErrChecker(os.RemoveAll(filepath.Join(folder, ".git")))
	ErrChecker(os.RemoveAll(filepath.Join(folder, ".github")))

	return nil
}
