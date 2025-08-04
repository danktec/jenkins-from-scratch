package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerCopyPlugins(t *testing.T) {
	tag := "gruntwork/test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../", buildOptions)

	opts := &docker.RunOptions{Command: []string{"ls", "/usr/share/jenkins/ref/plugins.txt"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "/usr/share/jenkins/ref/plugins.txt", output)
}

func TestDockerCopyCasc(t *testing.T) {
	tag := "gruntwork/test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../", buildOptions)

	opts := &docker.RunOptions{Command: []string{"ls", "/usr/share/jenkins/ref/jenkins-casc.yml"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "/usr/share/jenkins/ref/jenkins-casc.yml", output)
}

func TestDockerRunCmd(t *testing.T) {
	tag := "gruntwork/test"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../", buildOptions)

	opts := &docker.RunOptions{Command: []string{"jenkins-plugin-cli", "-f", "/usr/share/jenkins/ref/plugins.txt"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "Done", output)
}