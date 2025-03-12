package cmd

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	"github.com/ohhfishal/schedule/db"
)

type Root struct {
	Driver     string `default:"sqlite" env:"DRIVER" help:"Driver to use as for backed"`
	DataSource string `default:"schedule.db" env:"DATA_SOURCE" help:"Connection string for driver"`
	Verbose    bool   `short:"v" help:"Print more information to screen."`
	New        New    `cmd:"" help:"Create a new event."`
	Get        Get    `cmd:"" help:"Get events"`
	Delete     Delete `cmd:"" help:"Delete events by ID"`
	Edit       Edit   `cmd:"" help:"Edit an event by ID (NOT IMPLEMENTED)"`
}

type Stdout interface {
	io.Writer
	Verbose() io.Writer
}
type stdoutWriter struct {
	stdout  io.Writer
	verbose io.Writer
	builder strings.Builder
}

func (sw *stdoutWriter) Write(p []byte) (n int, err error) {
	return sw.stdout.Write(p)
}

func (sw *stdoutWriter) Verbose() io.Writer {
	if sw.verbose != nil {
		return sw.verbose
	}
	return &sw.builder
}

func Run(ctx context.Context, stdout io.Writer, args []string) error {
	stdoutWriter := &stdoutWriter{
		stdout: stdout,
	}

	var root Root
	parser, err := kong.New(
		&root,
		kong.Bind(time.Now),
		kong.BindTo(ctx, new(context.Context)),
		kong.BindTo(stdout, new(io.Writer)),
		kong.BindTo(stdoutWriter, new(Stdout)),
	)
	if err != nil {
		return err
	}

	parser.Stdout = stdout

	parsed, err := parser.Parse(args)
	if err != nil {
		return fmt.Errorf(`parsing args: %w`, err)
	}

	if root.Verbose {
		stdoutWriter.verbose = stdout
	}

	queries, err := db.Connect(ctx, root.Driver, root.DataSource)
	if err != nil {
		return err
	}

	if err := parsed.Run(ctx, stdoutWriter, queries); err != nil {
		return err
	}
	return nil
}
