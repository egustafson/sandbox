package main

import (
	log "github.com/Sirupsen/logrus" // is compatabile with std log
)

func main() {
	log.Print("Simple log message") // using golang "API" (calling Print())

	log.Info("Simple log message, Logrus 'API', calling Info()")

	log.WithFields(log.Fields{
		"widget": "framus",
	}).Info("Another cog in the machine")
}
