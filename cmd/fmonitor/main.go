package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/xapima/fmonitor"
)

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func logFatal(err error) {
	log.Fatalf("Error: webStatusChecker %s %s", time.Now(), err)
}

func main() {

	execPath, err := os.Executable()
	if err != nil {
		logFatal(errors.Wrap(err, "cause in main: "))
	}

	log.SetPrefix("fmonitor: ")
	log.SetFlags(0)
	configPath := flag.String("t", filepath.Join(filepath.Dir(execPath), "config.yml"), "path to config.yml")
	outputPath := flag.String("o", "", "output file path. If not set, it will be output to standard output")
	maxParallelNum := flag.Int("n", 200, "Parallel number")
	// verbose := flag.Bool("v", false, "show verbose")
	flag.Parse()

	// if !*verbose {
	// 	log.SetOutput(ioutil.Discard)
	// }

	if !exists(*configPath) {
		logFatal(fmt.Errorf("Error: fmonitor config.yml is not exist"))
	}
	monitor, err := fmonitor.NewMonitor()
	if err != nil {
		logFatal(errors.Wrap(err, "cause in main"))
	}
	defer monitor.Close()

	if err := monitor.Fmonitor(*configPath, *outputPath, *maxParallelNum); err != nil {
		logFatal(errors.Wrap(err, "cause in main"))
	}

}
