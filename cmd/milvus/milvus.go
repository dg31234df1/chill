package milvus

import (
	"flag"
	"fmt"
	"os"
)

const (
	DryRunCmd = "dry-run"
)

var (
	BuildTags = "unknown"
	BuildTime = "unknown"
	GitCommit = "unknown"
	GoVersion = "unknown"
)

type command interface {
	execute(args []string, flags *flag.FlagSet)
}

type dryRun struct{}

func (c *dryRun) execute(args []string, flags *flag.FlagSet) {}

type defaultCommand struct{}

func (c *defaultCommand) execute(args []string, flags *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "unknown command : %s\n", args[1])
	fmt.Fprintln(os.Stdout, usageLine)
}

func RunMilvus(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, usageLine)
		return
	}
	cmd := args[1]
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, usageLine)
	}

	var c command
	switch cmd {
	case RunCmd:
		c = &run{}
	case StopCmd:
		c = &stop{}
	case DryRunCmd:
		c = &dryRun{}
	case MckCmd:
		c = &mck{}
	default:
		c = &defaultCommand{}
	}

	c.execute(args, flags)
}
