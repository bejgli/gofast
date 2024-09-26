package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/bejgli/gofast/sorter"
	"github.com/fsnotify/fsnotify"
)

var CONFIGPATH = "patterns.json"

func main() {
	log.Print("Reading configuration file.")
	configFile, err := os.ReadFile(CONFIGPATH)
	if err != nil {
		log.Fatal("Couldn't open config file.")
	}

	var conf sorter.Config
	err = json.Unmarshal(configFile, &conf)
	if err != nil {
		log.Fatal(err)
	}

	badPattern := conf.CheckPatterns()
	if badPattern != "" {
		log.Fatalf("Bad regex pattern '%s'.", badPattern)
	}
	badDir := conf.CheckDirs()
	if badDir != "" {
		log.Fatalf("Directory '%s' doesn't exist.", badDir)
	}

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	log.Print("Watching: ", conf.Source)

	go watchLoop(watcher, conf)

	err = watcher.Add(conf.Source)
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}

func watchLoop(watcher *fsnotify.Watcher, conf sorter.Config) {
	for {
		select {

		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)

			if event.Has(fsnotify.Create) {
				files, err := os.ReadDir(conf.Source)
				if err != nil {
					log.Fatal("Cannot read source Dir", conf.Source)
				}

				log.Println("Sorting files.")
				err = sorter.SortFiles(files, conf)
				if err != nil {
					log.Fatal(err)
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
