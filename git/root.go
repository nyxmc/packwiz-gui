package git

import (
	"github.com/go-git/go-git/v5"
	"os"
	"path/filepath"
)

var (
	Repository = GetRepo()
)

func GetRepo() *git.Repository {
	// TODO: Handle error
	home, _ := os.UserHomeDir()

	// TODO: Make this configurable
	path := filepath.Join(home, "Documents", "darkmoon-modpack")

	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		// TODO: Handle error
		repo, _ := git.PlainOpen(path)
		return repo
	}

	// TODO: Handle error
	repo, _ := git.PlainClone(path, false, &git.CloneOptions{
		// TODO: Make this configurable
		URL: "https://github.com/nyxmc/darkmoon-modpack",
	})

	return repo
}

func CommitAndPushRepo() {
	Repository.Ad

	// TODO: Handle error
	// TODO: Auth?
	// Repository.Push(nil)
}
