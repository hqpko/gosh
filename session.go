package gosh

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Session struct {
	env map[string]string
	dir string

	echoTime bool
	echoPkg  bool
	echoIn   bool
	echoOut  bool
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

func (s *Session) echo(str string) {
	infos := strings.Split(str, "\n")
	for _, info := range infos {
		if info == "" {
			continue
		}
		echo := ""
		if s.echoTime {
			echo += time.Now().Format("15:04:05") + " : "
		}
		if s.echoPkg {
			dir := s.dir
			if dir == "" {
				dir = os.Args[0]
				dir, _ = filepath.Abs(dir)
			}
			echo += dir + " > "
		}
		echo += info
		fmt.Println(echo)
	}
}

func (s *Session) Run(cmds ...string) error {
	cmd := ""
	for _, c := range cmds {
		cmd += c + " "
	}
	c := exec.Command("/bin/sh", "-c", cmd)
	c.Dir = s.dir
	c.Env = os.Environ()
	for k, v := range s.env {
		c.Env = append(c.Env, k+"="+v)
	}

	var out bytes.Buffer
	var errout bytes.Buffer
	if s.echoOut {
		c.Stdout = &out
		c.Stderr = &errout
	}
	if s.echoIn {
		s.echo(cmd)
	}

	err := c.Run()
	if s.echoOut {
		if out.Len() > 0 {
			s.echo(out.String())
		}
		if errout.Len() > 0 {
			s.echo(errout.String())
		}
	}
	return err
}
