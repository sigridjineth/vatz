package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestPluginInstall(t *testing.T) {
	defer os.Remove("cosmos-status")
	defer os.Remove("./vatz.db")

	pluginDir = os.Getenv("PWD")

	root := cobra.Command{}
	root.AddCommand(createPluginCommand())
	root.SetArgs([]string{
		"plugin",
		"install",
		"github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_active_status",
		"cosmos-status"})

	err := root.Execute()
	assert.Nil(t, err)
}

// TODO: Test Start.
