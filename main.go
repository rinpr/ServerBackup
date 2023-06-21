package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	//archive.EzCompress("1.19.3 MyFirstServer")
	//if err := zipSource("testFolder/1/1.1/1.1.txt", "1.1.zip"); err != nil {
	//	log.Fatal(err)
	//}
}

func KillProcessByName(name string) error {
	// Get the process information by name
	output, _ := exec.Command("ps", "-o", "pid,comm", "-C", name).Output()
	lines := strings.Split(string(output), "\n")

	for i, line := range lines {
		if i == 0 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		fmt.Printf("Killing process: %s (PID: %s)\n", fields[1], fields[0])
		// Kill the process
		err := exec.Command("pkill", "-f", fields[0]).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
