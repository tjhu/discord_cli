//go:build tools
// +build tools

// A list of CLI tools.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/spf13/cobra-cli"
)
