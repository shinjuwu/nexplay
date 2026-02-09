package ginweb

type ILogger interface {
	Debug(format string, v ...interface{})

	Info(format string, v ...interface{})

	Warn(format string, v ...interface{})

	Error(format string, v ...interface{})

	Print(v ...interface{})

	Println(v ...interface{})

	Printf(format string, v ...interface{})

	Fatal(v ...interface{})

	Fatalln(v ...interface{})

	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})

	Panicln(v ...interface{})

	Panicf(format string, v ...interface{})
}
