package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func execFFmpegCommand(inputFileName, outputFilename, startSec, endSec string, isAudio, isDryrun bool) error {
	commandsOptions := []string{
		"-ss",
		startSec,
		"-i",
		inputFileName,
		"-t",
		endSec,
		"-c",
		"copy",
	}
	if isAudio {
		commandsOptions = append(commandsOptions, "-vn")
	}
	commandsOptions = append(commandsOptions, outputFilename)

	if isDryrun {
		optionJoinedString := strings.Join(commandsOptions, " ")
		fmt.Printf("ffmpeg %s\n", optionJoinedString)
		return nil
	}

	cmd := exec.Command("ffmpeg", commandsOptions...)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
