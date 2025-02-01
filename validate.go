package main

import (
	"errors"
	"os"
	"regexp"
)

func flagValidate(opt options) (errs []error) {
	if opt.inputFileName == "" {
		errs = append(errs, errors.New("No input file specified."))
	}
	if !isFileExists(opt.inputFileName) {
		errs = append(errs, errors.New("Input file is not found"))
		return
	}
	if opt.outputFileName == "" {
		errs = append(errs, errors.New("No output file specified."))
	}

	if opt.startTimeString == "" {
		errs = append(errs, errors.New("No start time specified."))
	}
	if !isValidateTimeString(opt.startTimeString) {
		errs = append(errs, errors.New("Start time format is invalid"))
	}

	if opt.endTimeString == "" {
		errs = append(errs, errors.New("No end time specified."))
	}
	if !isValidateTimeString(opt.endTimeString) {
		errs = append(errs, errors.New("End time format is invalid"))
	}
	return errs

}

func isValidateTimeString(timeString string) bool {
	return isColonTime(timeString) || isHMSTime(timeString)
}

func isColonTime(timeString string) bool {
	// 0:00:00形式
	patternColon := `\d+:\d{2}:\d{2}`
	reColon := regexp.MustCompile(patternColon)
	return reColon.MatchString(timeString)
}

func isHMSTime(timeString string) bool {
	// 0h00m00s形式
	patternHMS := `\d+h\d{2}m\d{2}s`
	reHMS := regexp.MustCompile(patternHMS)
	return reHMS.MatchString(timeString)
}

func isFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
