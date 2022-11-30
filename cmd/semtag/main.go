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
	version  = "dev"
	revision = "dev"
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
		msg := fmt.Sprintf("%s %s (%s)", appName, version, revision)
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

	action := args.Args[0]
	switch action {
	case "new":
	case "major":
	case "minor":
	case "patch":
	}

	return exitStatusOK
}
