// testutil.go
package cmd

import (
	"bytes"

	"github.com/spf13/cobra"
)

// ExecCmd runs a Cobra command with args and returns combined stdout/stderr plus error.
func ExecCmd(c *cobra.Command, args ...string) (string, error) {
	var buf bytes.Buffer
	c.SetOut(&buf)
	c.SetErr(&buf)
	c.SetArgs(args)
	err := c.Execute()
	return buf.String(), err
}
