package awhlog

import "log"

func LogInfo(message string) {
	log.Printf("INFO: %s", message)
}
func LogWarn(message string) {
	log.Printf("WARN: %s", message)
}
func LogError(message string) {
	log.Printf("ERROR: %s", message)
}
