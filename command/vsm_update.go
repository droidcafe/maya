package command

import (
	"os/exec"
	"strings"

	"github.com/mitchellh/cli"
)

type VsmUpdateCommand struct {
	// To control this CLI's display
	M Meta

	// OS command to execute; <optional>
	Cmd *exec.Cmd

	// Check the help section to learn more on these variables
	plan bool
}

func (c *VsmUpdateCommand) Help() string {
	helpText := `
Usage: maya vsm-update [path-to-update-specs]

  Update the given vsm.   

General Options:

  ` + generalOptionsUsage() + `
`
	return strings.TrimSpace(helpText)
}

func (c *VsmUpdateCommand) Synopsis() string {
	return "Updates the vsm with the provided specs"
}

func (c *VsmUpdateCommand) Run(args []string) int {
	var runop int

	flags := c.M.FlagSet("vsm-update", FlagSetClient)
	flags.Usage = func() { c.M.Ui.Output(c.Help()) }

	if err := flags.Parse(args); err != nil {
		return 1
	}

	// vsm specs is required
	oargs := flags.Args()
	if len(oargs) != 1 {
		c.M.Ui.Error(c.Help())
		return 1
	}

	if c.Cmd != nil {
		// execute the provided command
		return execute(c.Cmd, c.M.Ui)
	}

	// execute vsm update
	args = append([]string{string(NomadRun)}, oargs...)
	c.Cmd = exec.Command(string(ExecNomad), args...)

	if runop = execute(c.Cmd, c.M.Ui); runop != 0 {
		c.M.Ui.Error("Error updating vsm")
	}

	return runop
}

func execute(cmd *exec.Cmd, ui cli.Ui) int {

	ic := &InternalCommand{
		Cmd: cmd,
		Ui:  ui,
	}

	return ic.Execute()
}
