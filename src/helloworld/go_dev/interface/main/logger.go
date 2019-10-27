package main

type LogWriter interface {
	Write(data interface{}) error
}

type logger struct {
	writerList []LogWriter
}

func (l *logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

func (l *logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

func NewLogger() *logger {
	return &logger{}
}
