// A generated module for GoDicemaster functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"fmt"
)

type GoDicemaster struct{}

// The BuildEnv function creates a container for building the dicemaster binary
func (m *GoDicemaster) BuildEnv(source *Directory, goversion string) *Container {
	build := dag.Container().From("golang:"+goversion).WithDirectory("/src", source).WithWorkdir("/src")
	return build
}

// The Build function builds the dicemaster binary for linux and darwin
func (m *GoDicemaster) Build(source *Directory) *Directory {
	gooses := []string{"linux", "darwin"}
	goarches := []string{"amd64", "arm64"}

	outputs := dag.Directory()
	for _, goos := range gooses {
		for _, goarch := range goarches {
			path := fmt.Sprintf("/build/dicemaster-%s-%s", goos, goarch)
			build := m.BuildEnv(source, "latest").
				WithEnvVariable("GOOS", goos).
				WithEnvVariable("GOARCH", goarch).
				WithExec([]string{"go", "build", "-o", path}).File(path)
			outputs = outputs.WithFile(path, build)
		}
	}
	return outputs
}
