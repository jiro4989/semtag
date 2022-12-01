package main

import (
	"fmt"
	"os"
)

const (
	appName = "semtag"
)

// CIでビルド時に値を埋め込む。
// 埋め込む値の設定は .goreleaser.yaml を参照。
var (
	appVersion = "dev"
	revision   = "dev"
)

const (
	exitStatusOK = iota
	exitStatusCLIError
	exitStatusConvertError
	exitStatusInputFileError
	exitStatusOutputError
)

func main() {
	args, err := ParseArgs()
	if err != nil {
		Err(err)
		os.Exit(exitStatusCLIError)
	}

	os.Exit(Main(args))
}

func Main(args *CmdArgs) int {
	if args.Version {
		msg := fmt.Sprintf("%s %s (%s)", appName, appVersion, revision)
		fmt.Println(msg)
		fmt.Println("")
		fmt.Println("author:     jiro")
		fmt.Println("repository: https://github.com/jiro4989/semtag")
		return exitStatusOK
	}

	// if args.Completions != "" {
	// 	printCompletions(args.Completions)
	// 	return
	// }

	// if args.Text != "" {
	// 	exitStatus, err := run(args.Text, args)
	// 	if err != nil {
	// 		Err(err)
	// 		os.Exit(exitStatus)
	// 	}
	// 	os.Exit(exitStatus)
	// }

	if len(args.Args) < 1 {
		return exitStatusCLIError
	}

	t, err := NewTagger(".")
	if err != nil {
		Err(err)
		return 1
	}

	bi := &bumpInput{
		Tagger: t,
		Run:    args.Run,
	}
	action := args.Args[0]
	switch action {
	case "new":
		input := &NewVersionInput{
			Prefix:    args.Prefix,
			Major:     0,
			Minor:     1,
			Patch:     0,
			Separator: "-",
			Suffix:    args.Suffix,
		}
		v := NewVersion(input)
		tag := v.String()
		if args.Run {
			if err := t.CreateTag(tag); err != nil {
				Err(err)
				return 1
			}
		}
		printMsg(tag, args.Run)
	case "major":
		err := bump(bi, func(v *Version) {
			v.BumpMajor()
		})
		if err != nil {
			Err(err)
			return 1
		}
	case "minor":
		err := bump(bi, func(v *Version) {
			v.BumpMinor()
		})
		if err != nil {
			Err(err)
			return 1
		}
	case "patch":
		err := bump(bi, func(v *Version) {
			v.BumpPatch()
		})
		if err != nil {
			Err(err)
			return 1
		}
	}

	return exitStatusOK
}

type bumpInput struct {
	Tagger Tagger
	Run    bool
}

func bump(b *bumpInput, fn func(v *Version)) error {
	input := &LatestVersionInput{
		Tagger: b.Tagger,
	}
	v, err := LatestVersion(input)
	if err != nil {
		return fmt.Errorf("failed to get latest tag: %w", err)
	}
	if b.Run {
		fn(v)
	}
	tag := v.String()
	printMsg(tag, b.Run)
	return nil
}

func wrapDryRun(msg string, run bool) string {
	if !run {
		msg = "DRYRUN: " + msg
	}
	return msg
}

func printMsg(msg string, run bool) {
	fmt.Println(wrapDryRun(msg, run))
}
