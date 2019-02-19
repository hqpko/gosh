package gosh

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	defHandlerIn = func(s *Session, in string) {
		fmt.Printf("%s < %s\n", time.Now().Format("2006-01-02 15:04:05"), in)
	}

	defHandlerOut = func(s *Session, out []byte) {
		bs := bytes.Split(out, []byte{'\n'})
		for _, b := range bs {
			fmt.Printf("%s > %s\n", time.Now().Format("2006-01-02 15:04:05"), string(b))
		}
	}

	defHandlerErr = func(s *Session, err *exec.ExitError) {
		bs := strings.Split(err.Error(), `\n`)
		for _, b := range bs {
			fmt.Printf("%s > %s\n", time.Now().Format("2006-01-02 15:04:05"), string(b))
		}
	}
)

type Session struct {
	env map[string]string
	dir string

	handlerIn  func(s *Session, cmd string)
	handlerOut func(s *Session, out []byte)
	handlerErr func(s *Session, err *exec.ExitError)
}

func (s *Session) GetDir() string {
	return s.dir
}

func (s *Session) SetDir(dir string) {
	s.dir = dir
}

func (s *Session) SetEvn(key, value string) {
	s.env[key] = value
}

func (s *Session) GetEvn(key string) (string, bool) {
	v, ok := s.env[key]
	return v, ok
}

func (s *Session) SetHandlerIn(handler func(s *Session, cmd string)) {
	if handler != nil {
		s.handlerIn = handler
	}
}

func (s *Session) SetHandlerOut(handler func(s *Session, out []byte)) {
	if handler != nil {
		s.handlerOut = handler
	}
}

func (s *Session) SetHandlerErr(handler func(s *Session, err *exec.ExitError)) {
	if handler != nil {
		s.handlerErr = handler
	}
}

func (s *Session) Run(commands ...string) error {
	cmdStr := strings.Join(commands, " && ")
	cmd := s.getCmd(cmdStr)
	s.handlerIn(s, cmdStr)
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			s.handlerErr(s, exitErr)
		}
		return err
	}
	s.handlerOut(s, out)
	return nil
}

func (s *Session) getCmd(commandStr string) *exec.Cmd {
	cmd := exec.Command("/bin/sh", "-c", commandStr)
	cmd.Dir = s.dir
	cmd.Env = os.Environ()
	for k, v := range s.env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}
	return cmd
}
