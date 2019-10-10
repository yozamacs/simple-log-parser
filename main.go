package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type mapEntry struct {
	fivexx float64
	total  float64
}

type logParser struct {
	logFiles  []string
	logMap    map[string]*mapEntry
	startTime float64
	endTime   float64
}

func main() {
	startTime := flag.Float64("startTime", 0, "Start time for log parsing")
	endTime := flag.Float64("endTime", 0, "End time for log parsing")
	logFileNames := flag.String("logFiles", "", "Comma separated list of log files (no spaces)")
	flag.Parse()

	logFiles := strings.Split(*logFileNames, ",")

	logParser := logParser{
		logFiles:  logFiles,
		logMap:    make(map[string]*mapEntry),
		startTime: *startTime,
		endTime:   *endTime,
	}
	logParser.parse()

	fmt.Printf("Between time %v and %v:\n", startTime, endTime)
	for domain, entry := range logParser.logMap {
		avg := entry.fivexx / entry.total * 100
		fmt.Printf("%v returned %v%% 5xx errors\n", domain, avg)
	}
}

func (lp *logParser) parse() {
	for _, logFile := range lp.logFiles {
		file, err := os.Open(logFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lp.processLine(scanner.Text())
		}
	}
}

func (lp *logParser) processLine(line string) {
	entry := strings.Split(line, " | ")
	if len(entry) < 10 {
		return
	}
	responseTime, err := strconv.ParseFloat(entry[0], 64)
	if err != nil {
		panic(err)
	}
	if responseTime >= lp.startTime && responseTime < lp.endTime {
		domain := entry[2]
		_, inMap := lp.logMap[domain]
		if !inMap {
			lp.logMap[domain] = &mapEntry{0, 0}
		}
		lp.logMap[domain].total++
		matched, err := regexp.MatchString(`^5\d{2}$`, entry[4])
		if err != nil {
			panic(err)
		}
		if matched {
			lp.logMap[domain].fivexx++
		}
	}
}
