package log_for_project

import (
	"fmt"
	"log"
)

func Println(message ...interface{}) {

	log.Println(RedM(message...))
}

func Print(message ...interface{}) {
	log.Print(RedM(message...))
}

func Fatal(message ...interface{}) {
	log.Fatal(RedM(message...))
}

func Fatalf(format string, message ...interface{}) {
	log.Fatalf(RedM(format, message))
}
func Fatalln(message ...interface{}) {
	log.Fatalln(RedM(message...))
}

func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

func RedM(message ...interface{}) string {
	return fmt.Sprintf("\x1b[31m%v\x1b[0m", message...)
}
