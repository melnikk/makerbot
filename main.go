package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	flags "github.com/jessevdk/go-flags"
	"github.com/melnikk/makerbot/makerbot"
	"github.com/skbkontur/bot"
)

var (
	appname = "MakerBot"
	version = "latest"
)

var opts struct {
	ConfigFile string `short:"c" long:"config" default:"/etc/makerbot/config.yml" description:"Configuration file name"`
	Version    bool   `long:"version" description:"Show version info and exit"`
}

// in/out streams
var (
	in  io.Reader = os.Stdin
	out io.Writer = os.Stdout

	config *makerbot.Config
	db     bot.Database
)

func main() {
	fmt.Fprintln(out, appname)
	done := make(chan bool)
	processor := config.Dialog

	go bot.StartTelegramBot(config.Token, processor)

	<-done
}

func init() {
	flags.Parse(&opts)
	checkVersion()
	config, _ = readConfiguration(opts.ConfigFile)
}

func checkVersion() {
	if opts.Version {
		fmt.Fprintln(out, version)
		os.Exit(0)
	}
}

func readConfiguration(filename string) (*makerbot.Config, error) {
	var c makerbot.Config
	filename, _ = filepath.Abs(filename)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	return &c, err
}
