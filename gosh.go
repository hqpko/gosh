package gosh

var (
	EchoTime      = false
	EchoPkg  bool = false
	EchoIn   bool = false
	EchoOut  bool = false
)

func NewSession() *Session {
	return &Session{env: make(map[string]string)}
}

func NewSessionWithEcho(tim, pkg, in, out bool) *Session {
	s := NewSession()
	s.echoTime = tim
	s.echoPkg = pkg
	s.echoIn = in
	s.echoOut = out
	return s
}

func Run(cmd string) error {
	return NewSessionWithEcho(EchoTime, EchoPkg, EchoIn, EchoOut).Run(cmd)
}
