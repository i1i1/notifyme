package main

import (
	"fmt"
	"os"
	"bufio"
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

func copy_arr(a []byte, b []byte, n int) {
	for i := 0; i < n; i++ {
		b[i] = a[i]
	}
}

func input() string {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 4096)
	i := 0

	for {
		var b byte
		var e error

		if b, e = r.ReadByte(); e != nil {
			break
		}

		buf[i] = b
		i += 1

		if i == len(buf) {
			tmp := make([]byte, len(buf) * 2)
			copy_arr(buf, tmp, i)
			buf = tmp
		}
	}
	return string(buf)
}

func main() {
	b := Bot{"Token from @BotFather", "Markdown"}
	id := 12345 // Get from @JsonDumpBot
	s := ""

	if len(os.Args) == 1 {
		b.send(id, 0, input());
		return
	}

	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i];
	}

	b.send(id, 0, s);
}

