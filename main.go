package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	signal.Ignore(syscall.SIGINT)
}

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		m := make(map[string]any)
		err := json.Unmarshal([]byte(s), &m)
		if err != nil {
			// just display log normally if it is not in JSON format
			fmt.Println(s)
			continue
		}

		level := m["level"].(string)
		switch level {
		case "INFO":
			fmt.Print(Green)
		case "WARN":
			fmt.Print(Magenta)
		case "ERROR":
			fmt.Print(Red)
		case "DEBUG":
			fmt.Print(Blue)
		}
		fmt.Printf("%s%s => ", level, Reset)

		tStr, ok := m["time"].(string)
		if !ok {
			fmt.Println(s)
			log.Println("^time was not a string")
			continue
		}
		t, err := time.Parse(time.RFC3339Nano, tStr)
		if err != nil {
			fmt.Println(s)
			log.Println("^could not parse time")
			continue
		}
		fmt.Printf("%s", t.Format(time.DateTime))

		fmt.Printf(` "%s"`, m["msg"])

		for k, v := range m {
			switch k {
			case "time", "level", "msg":
				continue
			}

			fmt.Printf(" %s[%v]", k, v)
		}

		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"
