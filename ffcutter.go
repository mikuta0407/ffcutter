package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

type options struct {
	inputFileName   string
	outputFileName  string
	startTimeString string
	endTimeString   string
	audioOnly       bool
	help            bool
	dryrun          bool
}

func main() {
	opt := options{}
	flag.StringVarP(&opt.inputFileName, "input", "i", "", "input file name")
	flag.StringVarP(&opt.outputFileName, "output", "o", "", "output file name")
	flag.StringVarP(&opt.startTimeString, "start", "s", "", "start time (1:23:45 or 1h23m45s)")
	flag.StringVarP(&opt.endTimeString, "end", "e", "", "end time (1:23:45 or 1h23m45s)")
	flag.BoolVarP(&opt.audioOnly, "audio", "a", false, "audio only mode")
	flag.BoolVarP(&opt.help, "help", "h", false, "show help message")
	flag.BoolVar(&opt.dryrun, "dryrun", false, "dryrun mode (print command only)")
	flag.Parse()
	if opt.help {
		flag.PrintDefaults()
		return
	}

	// オプションバリデーション validate.go
	if validateErr := flagValidate(opt); len(validateErr) != 0 {
		for _, v := range validateErr {
			fmt.Println(v)
		}
		os.Exit(1)
	}

	// 時間パース timeparse.go
	startSec, duration, err := timeStringToSeconds(opt.startTimeString, opt.endTimeString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ffmpegコマンド実行 command.go
	if err := execFFmpegCommand(opt.inputFileName, opt.outputFileName, startSec, duration, opt.audioOnly, opt.dryrun); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
