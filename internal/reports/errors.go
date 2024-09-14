package reports

import "log"

func Error(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	log.Printf("[line %d]: Error (%s): %s", line, where, message)
}
