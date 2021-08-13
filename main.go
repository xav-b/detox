package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type BackupFile struct {
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Extension string   `json:"extension"`
	Tags      []string `json:"tags"`
	Size      int64    `json:"size"`
	Host      string   `json:"hostname"`
}

func (f BackupFile) Info() string {
	return fmt.Sprintf("%d %s %v (%s)", f.Size, f.Path, f.Tags, f.Host)
}

// dirToTags transforms the directory path into tags
// and then trim and lower-case them
func dirToTags(directory string) []string {
	var cleanTags []string
	for _, tag := range strings.Split(filepath.Dir(directory), "/") {
		if tag != "" {
			cleanTags = append(cleanTags, strings.ToLower(tag))
		}
	}

	return cleanTags
}

func main() {
	var IGNORE = map[string]bool{
		".DS_Store": true,
	}
	// TODO: something more standard (include time) or reqd from command line
	const output = "backup-meta.json"

	root := os.Args[1]
	log.Printf("Scanning filesystem from %s\n", root)

	// reset the file with a valid but empty first line
	_ = ioutil.WriteFile(output, []byte("{}\n"), 0666)

	// subsequent lines
	file, _ := os.OpenFile(output, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			log.Printf("Walking a new directory: %s\n", info.Name())
		} else if !IGNORE[info.Name()] {
			shortpath := strings.Replace(path, root, "", 1)

			hostname, _ := os.Hostname()
			// TODO: what is info.Sys()
			myfile := BackupFile{
				Name:      info.Name(),
				Path:      shortpath,
				Extension: filepath.Ext(info.Name()),
				Tags:      dirToTags(shortpath),
				Size:      info.Size(),
				Host:      hostname,
			}
			// TODO: only if debug enabled
			fmt.Println(myfile.Info())
			data, _ := json.Marshal(myfile)
			if _, err := file.WriteString(string(data) + "\n"); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}
}
