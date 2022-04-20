package cat

import (
	"os"
	"reflect"
	"testing"

	cli "github.com/tejzpr/zcli/v2"
	"github.com/tejzpr/zetools/commands"
)

func Test_catCommand_Name(t *testing.T) {
	tests := []struct {
		name string
		b    *catCommand
		want commands.CommandName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &catCommand{}
			if got := b.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("catCommand.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_catCommand_Aliases(t *testing.T) {
	tests := []struct {
		name string
		b    *catCommand
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &catCommand{}
			if got := b.Aliases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("catCommand.Aliases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_catCommand_Usage(t *testing.T) {
	tests := []struct {
		name string
		b    *catCommand
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &catCommand{}
			if got := b.Usage(); got != tt.want {
				t.Errorf("catCommand.Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_catCommand_Subcommands(t *testing.T) {
	tests := []struct {
		name string
		b    *catCommand
		want []*cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &catCommand{}
			if got := b.Subcommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("catCommand.Subcommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_catCommand_Flags(t *testing.T) {
	tests := []struct {
		name string
		b    *catCommand
		want []cli.Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &catCommand{}
			if got := b.Flags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("catCommand.Flags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_catCommand_Action(t *testing.T) {
	tests := []struct {
		name    string
		b       *catCommand
		args    []string
		wantErr bool
	}{
		{
			name:    "cat basic",
			b:       &catCommand{},
			args:    []string{"cat", "./test.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &catCommand{}
			app := &cli.App{
				Usage: "Unified commandline tools for common tasks",
				Commands: []*cli.Command{
					&cli.Command{
						Name:            string(b.Name()),
						Aliases:         b.Aliases(),
						Usage:           b.Usage(),
						Subcommands:     b.Subcommands(),
						Flags:           b.Flags(),
						Action:          b.Action,
						SkipFlagParsing: false,
					},
				},
			}
			args := os.Args[0:1]
			args = append(args, tt.args...)
			if err := app.Run(args); (err != nil) != tt.wantErr {
				t.Errorf("catCommand.Action() error = %v, wantErr %v", err.Error(), tt.wantErr)
			}
		})
	}
}
