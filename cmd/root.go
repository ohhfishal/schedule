package cmd

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"time"

	"github.com/alecthomas/kong"
	"github.com/ohhfishal/schedule/cmd/get"
	"github.com/ohhfishal/schedule/db"
)

type Root struct {
	Driver     string  `default:"sqlite" env:"DRIVER" help:"Driver to use as for backed"`
	DataSource string  `default:"schedule.db" env:"DATA_SOURCE" help:"Connection string for driver"`
	Verbose    bool    `short:"v" help:"Print more information to screen."`
	New        New     `cmd:"" help:"Create a new event."`
	Get        get.CMD `cmd:"" help:"Get events"`
	Delete     Delete  `cmd:"" help:"Delete events by ID"`
	Edit       Edit    `cmd:"" help:"Edit an event by ID"`
}

func Run(ctx context.Context, stdout io.Writer, args []string) error {
	var root Root
	parser, err := kong.New(
		&root,
		kong.Bind(time.Now),
		kong.BindTo(ctx, new(context.Context)),
		kong.Bind(time.Now().Location()),
		kong.BindTo(stdout, new(io.Writer)),
	)
	if err != nil {
		return err
	}

	parser.Stdout = stdout

	parsed, err := parser.Parse(args)
	if err != nil {
		return fmt.Errorf(`parsing args: %w`, err)
	}
	parsed.Bind(root.Verbose)

	queries, err := db.Connect(ctx, root.Driver, root.DataSource)
	if err != nil {
		return err
	}

	if err := parsed.Run(ctx, queries); err != nil {
		return err
	}
	return nil
}
