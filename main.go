package  main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	csvFileName := flag.String("file", "", "csv file")
	maxParallel := flag.Int("max-parallel", 1, "process in parallel")
	process := flag.String("process", "", "example: sleep.sh")
	sepatorCsvToProcess := flag.String("separator", ",", "csv separator")
	// TODO add arguments to process (no only csv line)
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if *maxParallel > len(csvLines) {
		log.Fatal("maxParallel flag can not be greater than number of lines of csv file")
	}

	maxParallelProcessChannel := make(chan int, *maxParallel)
	defer close(maxParallelProcessChannel)

	var wg sync.WaitGroup

	for _, line := range csvLines {
		maxParallelProcessChannel <- 1
		wg.Add(1)
		go func(line []string) {
			defer wg.Done()
			// EXECUTE COMMAND
			csvLine := strings.Join(line, *sepatorCsvToProcess)
			args := []string{*process}
			args = append(args, csvLine)

			cmd := &exec.Cmd {
				Path:   *process,
				Args:   args,
				Stdout: os.Stdout,
				Stderr: os.Stdout,
			}

			if err := cmd.Run(); err != nil {
				// TODO save errors in a channel
				fmt.Println( "Error:", err )
			}

			<-maxParallelProcessChannel

		}(line)
	}

	wg.Wait()
}