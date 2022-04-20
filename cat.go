package cat

import (
	"io"
	"os"

	cli "github.com/tejzpr/zcli/v2"
	"github.com/tejzpr/zetools/commands"
)

func init() {
	commands.RegisterCommand(CatCommandName, &catCommand{}, &commands.CommandOpts{SkipFlagParsing: true})
}

// CatCommandName returns the name of the command
const CatCommandName commands.CommandName = "cat"

// catCommand is the base64 command
type catCommand struct {
}

// Name returns the name of the command
func (b *catCommand) Name() commands.CommandName {
	return CatCommandName
}

// Aliases returns the aliases of the command
func (b *catCommand) Aliases() []string {
	return []string{"cat"}
}

// Usage returns the usage of the command
func (b *catCommand) Usage() string {
	return "cat filename"
}

// Subcommands returns the subcommands of the command
func (b *catCommand) Subcommands() []*cli.Command {
	return []*cli.Command{}
}

// Flags returns the flags of the command
func (b *catCommand) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Aliases:  []string{"p"},
			Value:    "",
			Usage:    "Path of the file to cat",
			Required: false,
		},
	}
}

// Action returns the action of the command
func (b *catCommand) Action(c *cli.Context) error {
	args := c.Args()

	if args.Len() == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			return err
		}
	} else {
		for _, fname := range args.Slice() {
			fh, err := os.Open(fname)
			if err != nil {
				return err
			}

			_, err = io.Copy(os.Stdout, fh)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
