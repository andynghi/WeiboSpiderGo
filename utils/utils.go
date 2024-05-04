package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var ExecPath = getExecutePath()

func ReParse(pattern string, content string) string {
	str := regexp.MustCompile(pattern).FindAllStringSubmatch(content, -1)
	if str != nil {
		if len(str[0]) == 1 {
			return str[0][0]
		}
		return str[0][1]
	}
	return ""
}

func ReParseMayLi(pattern string, content string) [][]string {
	str := regexp.MustCompile(pattern).FindAllStringSubmatch(content, -1)
	return str
}

func ConvTime(timeStr string) string {
	now_time := time.Now()
	if strings.Contains(timeStr, "minutes ago") {
		min, _ := strconv.Atoi(ReParse(`^(\d+)minute`, timeStr))
		createdTimep := now_time.Add(-time.Duration(min) * time.Minute)
		return createdTimep.Format("2006-01-02 15:04")
	}
	if strings.Contains(timeStr, "hours ago") {
		hour, _ := strconv.Atoi(ReParse(`^(\d+)hour`, timeStr))
		createdTimep := now_time.Add(-time.Duration(hour) * time.Hour)
		return createdTimep.Format("2006-01-02 15:04")
	}
	if strings.Contains(timeStr, "Today") {
		return strings.Replace(timeStr, "Today", now_time.Format("2006-01-02"), -1)
	}
	if strings.Contains(timeStr, "month") {
		rp := strings.NewReplacer("month", "-", "day", "")
		return rp.Replace(timeStr)
	}
	return timeStr
}

func GetTargetUidList() []string {
	var uidLi []string
	file, err := os.Open(ExecPath + "/account/target.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		lineText = strings.TrimSpace(lineText)
		lineText = strings.Replace(lineText, "\uFEFF", "", -1)
		uidLi = append(uidLi, lineText)
	}
	return uidLi
}

func getExecutePath() string {
	return filepath.Dir(os.Args[0])
}
