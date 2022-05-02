package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/encima/dgen/lib"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
)

func main() {
	err := godotenv.Load()
	pp := lib.PromPull{URI: os.Getenv("PROM_URI"), USER: os.Getenv("PROM_USER"), PASS: os.Getenv("PROM_PASS")}
	pp.Pull()
	f, err := os.Create("./sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var lines int
	flag.IntVar(&lines, "l", 10000, "default number of lines is 10000")
	flag.Parse()

	for idx := 1; idx < lines; idx++ {
		var line = fmt.Sprint("SET Key", idx, " Value", idx)
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd := exec.Command("bash", "-c", fmt.Sprintf("cat sample.txt | redis-cli -u '%s' --pipe", os.Getenv("SVC_URI")))
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}
