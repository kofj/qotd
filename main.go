package main

import (
	"io"
	"log"
	"math/rand"
	"net"
)

var quotes = []string{
	`天行健，君子当自强不息。
	-- 《周易》
`,

	`地势坤，君子以厚德载物。
	-- 《周易》
`,
	`窈窕淑女，君子好逑。
	-- 《国风·周南·关雎》
`,
}

func main() {
	l, err := net.Listen("tcp", ":1717")
	if err != nil {
		log.Fatal(err)
	}
	listen(l)
}

func listen(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("%s: %s", l.Addr(), err)
			continue
		}

		q := quotes[rand.Intn(len(quotes))]
		io.WriteString(c, q)
		c.Close()
	}
}
