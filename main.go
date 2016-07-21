package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gotascii/simpledb/aggregate"
	"github.com/gotascii/simpledb/data"
	"github.com/gotascii/simpledb/kvdb"
)

func main() {
	ts := kvdb.NewStack()
	counter := &aggregate.Counts{data.NewGtreap()}
	t := &kvdb.Transaction{data.NewGtreap(), counter}

	reader := bufio.NewReader(os.Stdin)

	for {
		text, er := reader.ReadString('\n')
		if er == io.EOF {
			return
		}

		text = strings.TrimSpace(text)
		cmd := strings.Split(text, " ")

		switch {
		case cmd[0] == "SET":
			t.Set(cmd[1], cmd[2])
		case cmd[0] == "GET":
			fmt.Println(t.Get(cmd[1]))
		case cmd[0] == "UNSET":
			t.Unset(cmd[1])
		case cmd[0] == "END":
			return
		case cmd[0] == "BEGIN":
			ts = ts.Push(t)
			t = t.Copy()
		case cmd[0] == "ROLLBACK":
			ts, t = ts.Pop()
		case cmd[0] == "COMMIT":
			ts = kvdb.NewStack()
		case cmd[0] == "NUMEQUALTO":
			fmt.Println(t.Numequalto(cmd[1]))
		}
	}
}
