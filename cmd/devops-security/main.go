package main

import "github.com/sirupsen/logrus"

const AppName = "devops-security"

var logger = logrus.New()
var config *Config = nil

var args struct {
	LogLevel   string         `arg:"--log-level" default:"warn"`
	MagicLinks *MagicLinksCmd `arg:"subcommand:magic-links"`
}

func main() {

}
