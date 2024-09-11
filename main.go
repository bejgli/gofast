package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"

	"github.com/bejgli/download-sorter/sorter"
	"github.com/fsnotify/fsnotify"
)

func main() {
	log.Print("Reading configuration file.")
	configFile, err := os.ReadFile("patterns.json")
	if err != nil {
		log.Fatal("Couldn't open config file.")
	}

	var conf sorter.Config

	err = json.Unmarshal(configFile, &conf)
	if err != nil {
		log.Fatal(err)
	}

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Create) {
					files, err := os.ReadDir("./" + conf.Source)
					if err != nil {
						log.Fatal("Cannot read source Dir", conf.Source)
					}

					log.Println("Sorting files.")
					for _, f := range files {
						if f.IsDir() {
							continue
						}

						for _, r := range conf.Rules {
							if regexp.MustCompile(r.Pattern).MatchString(f.Name()) {
								_, err = sorter.MoveFile(conf.Source+"/"+f.Name(), r.Target+"/"+f.Name())
								if err != nil {
									log.Fatal(err)
								}
							}
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add(conf.Source)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
