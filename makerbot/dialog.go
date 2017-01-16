package makerbot

import (
	"log"
	"os/exec"

	"github.com/skbkontur/bot"
)

// Process implements Processor method
func (c *Dialog) Process(update bot.Update) []bot.Chattable {

	answer := "Я тебя не понимаю, брат"

	command := update.MessageText()
	answer = c.execute(command)

	var messages []bot.Chattable
	messages = append(messages, bot.NewMessage(update.ChatID(), answer))
	return messages
}

// Until done
func (c *Dialog) Until(done chan bool) bot.Processor {
	c.done = done
	return c
}

func (c *Dialog) sanitize(text string) string {
	var result string
	for _, b := range text {
		if b == ' ' {
			b = '_'
		}
		if ('a' <= b && b <= 'z') || b == '_' {
			result += string(b)
		}
	}
	return result
}

func (c *Dialog) execute(text string) string {
	text = c.sanitize(text)
	out, err := exec.Command("make", text).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
