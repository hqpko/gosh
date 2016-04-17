package gosh

var (
	EchoTime = true
	EchoPkg  = false
	EchoIn   = true
	EchoOut  = true
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
