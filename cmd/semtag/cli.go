package main

import (
	"flag"
	"fmt"
	"os"
)

type CmdArgs struct {
	Run        bool
	Prefix     string
	Suffix     string
	PathSuffix string
	Separator  string
	Version    bool
	Args       []string
}

const (
	helpMsgRun         = "control git repository. if this option is not set then git repository is not controlled"
	helpMsgPrefix      = ""
	helpMsgSuffix      = ""
	helpMsgPathSuffix  = ""
	helpMsgVersion     = "print version"
	helpMsgCompletions = "print completions file. (bash, zsh)"
)

func ParseArgs() (*CmdArgs, error) {
	opts := CmdArgs{}

	flag.Usage = flagHelpMessage
	flag.StringVar(&opts.Prefix, "prefix", "v", helpMsgPrefix)
	flag.StringVar(&opts.Suffix, "suffix", "", helpMsgSuffix)
	flag.StringVar(&opts.PathSuffix, "path", "", helpMsgPathSuffix)
	flag.BoolVar(&opts.Run, "run", false, helpMsgRun)
	flag.BoolVar(&opts.Version, "v", false, helpMsgVersion)
	flag.Parse()
	opts.Args = flag.Args()

	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &opts, nil
}

func flagHelpMessage() {
	cmd := os.Args[0]
	fmt.Fprintln(os.Stderr, fmt.Sprintf("%s convert text to '%s' style.", cmd, appName))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s [OPTIONS] [files...]", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s sample.txt", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Options:")

	flag.PrintDefaults()
}

func (c *CmdArgs) Validate() error {
	// switch c.CharCode {
	// case "utf8", "sjis":
	// 	// 何もしない
	// default:
	// 	err := errors.New("charcode must be 'utf8' or 'sjis'.")
	// 	return err
	// }
	//
	// if c.Completions != "" && !isSupportedCompletions(c.Completions) {
	// 	return fmt.Errorf("illegal completions. completions = %s", c.Completions)
	// }

	return nil
}
