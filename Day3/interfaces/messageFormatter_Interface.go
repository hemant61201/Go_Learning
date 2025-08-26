package main

type formatter1 interface {
	format() string
}

type plainText struct {
	message string
}

func (p plainText) format() string {
	return p.message
}

type bold struct {
	message string
}

func (b bold) format() string {
	return "**" + b.message + "**"
}

type code struct {
	message string
}

func (c code) format() string {
	return "`" + c.message + "`"
}

func sendMessage1(format formatter1) string {
	return format.format()
}
