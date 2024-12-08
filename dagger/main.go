// Validate & publish fuhrmannb blog to dev.to
// https://dev.to/fuhrmannb

package main

import (
  "context"
  "fmt"

  "github.com/ettle/strcase"

  "dagger/blog/internal/dagger"
)

type Blog struct {
	// +private
	DevToToken *dagger.Secret
	// +private
	PostDir *dagger.Directory
}

func New(
	devToToken *dagger.Secret,
	// +defaultPath="./posts"
	postDir *dagger.Directory,
) *Blog {
	return &Blog{
		DevToToken: devToToken,
		PostDir:    postDir,
	}
}

// Create a new draft post on dev.to
func (m *Blog) NewPost(postName string) *dagger.Directory {
	file := fmt.Sprintf("posts/%s.md", strcase.ToKebab(postName))
	container := m.devToSetup().
		WithExec([]string{"dev", "new", file}).
		WithExec([]string{"sed", "-e", fmt.Sprintf("s|title:.*|title: %s|g", postName), "-i", file})
	return m.syncPosts(container, false).Directory("/posts")
}

// Sync all blog posts to dev.to.
//
// Note: this script does not ensure deletion of posts (only creation/update)
func (m *Blog) SyncPosts(
	ctx context.Context,
	// +optional
	reconcile bool,
) (string, error) {
	return m.syncPosts(m.devToSetup(), reconcile).Stdout(ctx)
}

func (m *Blog) devToSetup() *dagger.Container {
	return dag.Container().
		From("node:22").
		WithSecretVariable("DEVTO_TOKEN", m.DevToToken).
		// dev.to sync tool: https://github.com/sinedied/devto-cli
		WithExec([]string{"npm", "install", "-g", "@sinedied/devto-cli"}).
		WithMountedDirectory("/posts", m.PostDir)
}

func (m *Blog) syncPosts(container *dagger.Container, reconcile bool) *dagger.Container {
	devCommand := []string{"dev", "push", "-r", "fuhrmannb/blog", "-b", "main"}
	if reconcile {
		devCommand = append(devCommand, "-e")
	}

	return container.WithExec(devCommand)
}
