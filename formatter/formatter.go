package formatter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

func New(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		s := scanner.Text()

		m := make(map[string]any)
		err := json.Unmarshal([]byte(s), &m)
		if err != nil {
			// just display log normally if it is not in JSON format
			fmt.Fprintln(out, s)
			continue
		}

		level := m["level"].(string)
		switch level {
		case "INFO":
			fmt.Fprint(out, Green)
		case "WARN":
			fmt.Fprint(out, Magenta)
		case "ERROR":
			fmt.Fprint(out, Red)
		case "DEBUG":
			fmt.Fprint(out, Blue)
		}
		fmt.Fprintf(out, "%s%s => ", level, Reset)

		tStr, ok := m["time"].(string)
		if !ok {
			fmt.Fprintln(out, s)
			log.Println("^time was not a string")
			continue
		}
		t, err := time.Parse(time.RFC3339Nano, tStr)
		if err != nil {
			fmt.Fprintln(out, s)
			log.Println("^could not parse time")
			continue
		}
		fmt.Fprintf(out, "%s", t.Format(time.DateTime))

		fmt.Fprintf(out, ` "%s"`, m["msg"])

		for k, v := range m {
			switch k {
			case "time", "level", "msg":
				continue
			}

			fmt.Fprintf(out, " %s[%v]", k, v)
		}

		fmt.Fprintln(out)
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
