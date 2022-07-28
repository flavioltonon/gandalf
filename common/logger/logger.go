package logger

type Logger interface {
	Info(message string, fields ...Field)
	Error(message string, fields ...Field)
}

type Field struct {
	name  string
	value string
}

func String(name, value string) Field { return Field{name: name, value: value} }
func Error(err error) Field           { return Field{name: "error", value: err.Error()} }

func (f Field) Name() string  { return f.name }
func (f Field) Value() string { return f.value }
