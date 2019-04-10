package main

import (
	"fmt"
	"os"
)


func (b *Bot) send(id, reply int, s string) {
	if s == "" {
		if err := b.SendMessage(id, reply, "-/-/-/-"); err != nil {
			fmt.Print(err)
		}
		return
	}
	for len(s) > 4096 {
		if err := b.SendMessage(id, reply, string(s[:4095])); err != nil {
			fmt.Print(err)
		}
		s = s[4096:]
	}

	if err := b.SendMessage(id, reply, s); err != nil {
		fmt.Print(err)
	}
}

func main() {
	b := Bot{"Token from @BotFather", "Markdown"}
	id := 12345 // Get from @JsonDumpBot
	s := ""
	
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i];
	}

	b.send(id, 0, s);
}

