package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/backend"
	clistate "github.com/hashicorp/terraform/command/state"
	"github.com/hashicorp/terraform/state"
	"github.com/hashicorp/terraform/terraform"
)

// TaintCommand is a cli.Command implementation that manually taints
// a resource, marking it for recreation.
type TaintCommand2 struct {
	Meta
}

func (c *TaintCommand2) Run(args []string) int {
	args = c.Meta.process(args, false)

	var allowMissing bool
	var module string
	cmdFlags := c.Meta.flagSet("taint")
	cmdFlags.BoolVar(&allowMissing, "allow-missing", false, "module")
	cmdFlags.StringVar(&module, "module", "", "module")
	cmdFlags.StringVar(&c.Meta.statePath, "state", DefaultStateFilename, "path")
	cmdFlags.StringVar(&c.Meta.stateOutPath, "state-out", "", "path")
	cmdFlags.StringVar(&c.Meta.backupPath, "backup", "", "path")
	cmdFlags.BoolVar(&c.Meta.stateLock, "lock", true, "lock state")
	cmdFlags.Usage = func() { c.Ui.Error(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	// Require the one argument for the resource to taint
	args = cmdFlags.Args()
	if len(args) != 1 {
		c.Ui.Error("The taint command expects exactly one argument.")
		cmdFlags.Usage()
		return 1
	}

	// Start of Huy's code
	return 0
}

func (c *TaintCommand2) Help() string {
	helpText := `
Usage: terraform taint [options] name

  Manually mark a resource as tainted, forcing a destroy and recreate
  on the next plan/apply.

  This will not modify your infrastructure. This command changes your
  state to mark a resource as tainted so that during the next plan or
  apply, that resource will be destroyed and recreated. This command on
  its own will not modify infrastructure. This command can be undone by
  reverting the state backup file that is created.

Options:

  -allow-missing      If specified, the command will succeed (exit code 0)
                      even if the resource is missing.

  -backup=path        Path to backup the existing state file before
                      modifying. Defaults to the "-state-out" path with
                      ".backup" extension. Set to "-" to disable backup.

  -lock=true          Lock the state file when locking is supported.

  -module=path        The module path where the resource lives. By
                      default this will be root. Child modules can be specified
                      by names. Ex. "consul" or "consul.vpc" (nested modules).

  -no-color           If specified, output won't contain any color.

  -state=path         Path to read and save state (unless state-out
                      is specified). Defaults to "terraform.tfstate".

  -state-out=path     Path to write updated state file. By default, the
                      "-state" path will be used.

`
	return strings.TrimSpace(helpText)
}

func (c *TaintCommand2) Synopsis() string {
	return "Manually mark a resource for recreation"
}

func (c *TaintCommand2) allowMissingExit(name, module string) int {
	c.Ui.Output(fmt.Sprintf(
		"The resource %s in the module %s was not found, but\n"+
			"-allow-missing is set, so we're exiting successfully.",
		name, module))
	return 0
}
