package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/atotto/version/internal/file"

	"github.com/google/subcommands"
	"github.com/hashicorp/go-version"
)

func main() {
	subcommands.Register(&versionUp{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

type versionUp struct {
	major bool
	minor bool
	write bool
}

func (*versionUp) Name() string     { return "up" }
func (*versionUp) Synopsis() string { return "version up" }
func (*versionUp) Usage() string {
	return `up [-major] [-minor] filename:
`
}

func (c *versionUp) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.major, "major", false, "update major version (default: update patch version)")
	f.BoolVar(&c.minor, "minor", false, "update minor version (default: update patch version)")
	f.BoolVar(&c.write, "w", false, "write result to file instead of stdout")
}

func (c *versionUp) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	filename := f.Arg(0)
	if filename == "" {
		return fail("filename not found")
	}

	old, err := file.Version(filename)
	if err != nil {
		return fail("version up: %s", err)
	}

	v, err := version.NewVersion(old)
	if err != nil {
		return fail("failed to parse version: %s", err)
	}

	seg := v.Segments()
	if len(seg) < 3 {
		return fail("invalid segment")
	}
	switch {
	case c.major:
		seg[0]++
		seg[1] = 0
		seg[2] = 0
	case c.minor:
		seg[1]++
		seg[2] = 0
	default:
		seg[2]++
	}

	var new string
	if old[0] == 'v' {
		new = fmt.Sprintf("v%d.%d.%d", seg[0], seg[1], seg[2])
	} else {
		new = fmt.Sprintf("%d.%d.%d", seg[0], seg[1], seg[2])
	}

	if c.write {
		if err := file.Replace(filename, old, new); err != nil {
			return fail("failed to write file: %s", err)
		}
	} else {
		fmt.Printf(new)
	}
	return subcommands.ExitSuccess
}

func fail(format string, a ...interface{}) subcommands.ExitStatus {
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Fprintln(os.Stderr)
	return subcommands.ExitFailure
}
