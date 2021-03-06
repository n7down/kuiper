package blanklogger

type BlankLogger struct{}

func NewBlankLogger() *BlankLogger {
	return &BlankLogger{}
}

func (l *BlankLogger) Trace(args ...interface{}) {}

func (l *BlankLogger) Debug(args ...interface{}) {}

func (l *BlankLogger) Print(args ...interface{}) {}

func (l *BlankLogger) Info(args ...interface{}) {}

func (l *BlankLogger) Warn(args ...interface{}) {}

func (l *BlankLogger) Warning(args ...interface{}) {}

func (l *BlankLogger) Error(args ...interface{}) {}

func (l *BlankLogger) Panic(args ...interface{}) {}

func (l *BlankLogger) Fatal(args ...interface{}) {}

func (l *BlankLogger) Tracef(format string, args ...interface{}) {}

func (l *BlankLogger) Debugf(format string, args ...interface{}) {}

func (l *BlankLogger) Printf(format string, args ...interface{}) {}

func (l *BlankLogger) Infof(format string, args ...interface{}) {}

func (l *BlankLogger) Warnf(format string, args ...interface{}) {}

func (l *BlankLogger) Warningf(format string, args ...interface{}) {}

func (l *BlankLogger) Errorf(format string, args ...interface{}) {}

func (l *BlankLogger) Panicf(format string, args ...interface{}) {}

func (l *BlankLogger) Fatalf(format string, args ...interface{}) {}

func (l *BlankLogger) Traceln(args ...interface{}) {}

func (l *BlankLogger) Debugln(args ...interface{}) {}

func (l *BlankLogger) Println(args ...interface{}) {}

func (l *BlankLogger) Infoln(args ...interface{}) {}

func (l *BlankLogger) Warningln(args ...interface{}) {}

func (l *BlankLogger) Warnln(args ...interface{}) {}

func (l *BlankLogger) Errorln(args ...interface{}) {}

func (l *BlankLogger) Panicln(args ...interface{}) {}

func (l *BlankLogger) Fatalln(args ...interface{}) {}
