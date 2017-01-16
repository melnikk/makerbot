package makerbot

// Config of bot
type Config struct {
	Token  string  `yaml:"token"`
	Dialog *Dialog `yaml:"dialog"`
}

// Dialog is a conversation processor
type Dialog struct {
	File string `yaml:"file"`
	URL  string `yaml:"url"`
	done chan bool
}
