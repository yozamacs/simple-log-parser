package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestLogParsing(t *testing.T) {
	g := NewGomegaWithT(t)
	logFiles := []string{"./log_sample.txt"}
	logParser := logParser{
		logFiles:  logFiles,
		logMap:    make(map[string]*mapEntry),
		startTime: 1493969101.645,
		endTime:   1493969101.655,
	}
	logParser.parse()
	g.Expect(len(logParser.logMap)).To(Equal(2))
}
