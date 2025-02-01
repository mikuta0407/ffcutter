package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func timeStringToSeconds(startTimeString, endTimeString string) (startTimeSeconds, endTimeSeconds string, err error) {
	// 時間パース
	var startTimeParts []int
	var endTimeParts []int

	// 開始時間
	if isColonTime(startTimeString) {
		startTimeParts, err = parseColonTimeString(startTimeString)
	}
	if isHMSTime(startTimeString) {
		startTimeParts, err = parseHMSTimeString(startTimeString)
	}
	if err != nil {
		return
	}

	startTimeSeconds = strconv.Itoa(startTimeParts[0]*60*60 + startTimeParts[1]*60 + startTimeParts[2])

	// 終了時間
	if isColonTime(endTimeString) {
		endTimeParts, err = parseColonTimeString(endTimeString)
	}
	if isHMSTime(endTimeString) {
		endTimeParts, err = parseHMSTimeString(endTimeString)
	}
	if err != nil {
		return
	}

	endTimeSeconds = strconv.Itoa(endTimeParts[0]*60*60 + endTimeParts[1]*60 + endTimeParts[2])

	return
}

func parseColonTimeString(timeString string) ([]int, error) {
	parts := strings.Split(timeString, ":")
	result := make([]int, 3)
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result[i] = num
	}
	return result, nil
}

func parseHMSTimeString(timeString string) ([]int, error) {
	pattern := `(\d+)h(\d{2})m(\d{2})s`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(timeString)
	if len(matches) != 4 {
		return nil, errors.New("Input does not match the expected format")
	}
	result := make([]int, 3)
	for i, match := range matches[1:] {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, errors.New(fmt.Sprintln("Error converting string to int:", err))
		}
		result[i] = num
	}
	return result, nil
}
