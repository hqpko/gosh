package gosh

func NewSession() *Session {
	return &Session{env: make(map[string]string), handlerIn: defHandlerIn, handlerOut: defHandlerOut, handlerErr: defHandlerErr}
}

func Run(commands ...string) error {
	s := NewSession()
	return s.Run(commands...)
}
