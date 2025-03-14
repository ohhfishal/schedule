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
	Edit       Edit   `cmd:"" help:"Edit an event by ID"`
}

type Stdout interface {
	io.Writer
	Verbose() io.Writer
}
type StdoutWriter struct {
	Stdout        io.Writer
	VerboseWriter io.Writer
}

func (sw *StdoutWriter) Write(p []byte) (n int, err error) {
	if sw.Stdout == nil {
		return len(p), nil
	}
	return sw.Stdout.Write(p)
}

func (sw *StdoutWriter) Verbose() io.Writer {
	if sw.VerboseWriter != nil {
		return sw.VerboseWriter
	}
	return &strings.Builder{}
}

func Run(ctx context.Context, stdout io.Writer, args []string) error {
	stdoutWriter := &StdoutWriter{
		Stdout: stdout,
	}

	var root Root
	parser, err := kong.New(
		&root,
		kong.Bind(time.Now),
		kong.BindTo(ctx, new(context.Context)),
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
		stdoutWriter.VerboseWriter = stdout
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
