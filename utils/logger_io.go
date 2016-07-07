package utils

import (
	"bufio"
	"io"
	"log"
)

type waitWriteCloser struct {
	io.WriteCloser
	done chan bool
}

func (wc *waitWriteCloser) Close() error {
	err := wc.Close()
	<-wc.done
	return err
}

func LoggerIO(dest *log.Logger) io.WriteCloser {
	pipe_r, pipe_w := io.Pipe()
	buf_r := bufio.NewReader(pipe_r)
	done := make(chan bool)
	go func() {
		for {
			line, err := buf_r.ReadString('\n')
			if err != nil {
				done <- true
				return
			}
			dest.Output(1, line)
		}
	}()
	return &waitWriteCloser{pipe_w, done}
}
