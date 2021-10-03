package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	level := os.Getenv("LOG_LEVEL")
	//fmt.Printf("%v", level)
	if string(level) == "Debug" {
		log.SetLevel(log.DebugLevel)
	}

	/* for _, env := range os.Environ() {
		// env is
		envPair := strings.SplitN(env, "=", 2)
		key := envPair[0]
		value := envPair[1]

		fmt.Printf("%s : %s\n", key, value)
		os.LookupEnv(key)
	} */
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"user": "admin",
	}).Info()

	log.Warn("This is a warning")
	log.Error("An error occured! " + level)
	log.Debug(level, " level set here")
	log.Panic("Exiting!!")
}
