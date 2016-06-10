package main

import (
	"bufio"
	"fmt"
	"github.com/nlopes/slack"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type config struct {
	Deployment string `yaml:"deployment"`
	Token      string `yaml:"token"`
}

// ERR OMIT
func proc_config() (*config, error) { // HLerr
	file, err := os.Open("slacker.yml")
	defer file.Close()
	if err != nil {
		return nil, err
	}

	fi, err := file.Stat() // HLerr
	if err != nil {        // HLerr
		return nil, err // HLerr
	} // HLerr

	buff := make([]byte, fi.Size())
	if _, err = file.Read(buff); err != nil {
		return nil, err
	}

	c := &config{}
	err = yaml.Unmarshal(buff, c)
	return c, err
}

// ERRE OMIT

func post(msg string, conf *config) {
	api := slack.New(conf.Token)
	params := slack.PostMessageParameters{}
	_, _, err := api.PostMessage("ossecm",
		fmt.Sprintf(
			":ossecm: *_Deployment:_ %s*\n\n%s\n",
			conf.Deployment,
			msg), params)
	if err != nil {
		panic(err)
	}
}

// READ OMIT
// YAY!  I actually know wthat this thing is returning and the types. // HLread
func read() (*[]string, func(string)) { // HLread
	msg := make([]string, 0, 3)
	header := true
	f := func(line string) {
		if header {
			if strings.HasPrefix(line, "Message-Id:") {
				header = false
			}
			return
		}
		msg = append(msg, line)
		return
	}
	return &msg, f
}

// READE OMIT

func gen_message() string {
	msg, liner := read()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		liner(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return strings.Join(*msg, "\n")
}

func main() {
	config, err := proc_config()
	if err != nil {
		panic(err)
	}
	msg := gen_message()
	post(msg, config)
}
