package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const manifest = "Godeps.json"

// Dependency is the type for a single dependency in the Godeps.json file.
type Dependency struct {
	ImportPath string
	Comment    string
	Rev        string
}

// Godeps is the file format for Godeps.json
type Godeps struct {
	ImportPath string
	GoVersion  string
	Packages   []string
	Deps       []Dependency
}

// loadGodeps loads the Godeps.json file from the specified directory and
// parses it into a Godeps struct
func loadGodeps(directory string) (*Godeps, error) {
	filename := path.Join(directory, manifest)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	deps := new(Godeps)
	err = decoder.Decode(deps)
	if err != nil {
		return nil, err
	}
	return deps, nil
}

func (g *Godeps) filter(prefix string) {
	newdeps := []Dependency{}
	for _, d := range g.Deps {
		if strings.HasPrefix(d.ImportPath, prefix) {
			continue
		}
		newdeps = append(newdeps, d)
	}
	g.Deps = newdeps
}

func (g *Godeps) save(directory string) error {
	b, err := json.MarshalIndent(g, "", "\t")
	if err != nil {
		return err
	}

	filename := path.Join(directory, manifest)
	return ioutil.WriteFile(filename, b, 0660)
}

func main() {
	// Command line arguments
	var directory string
	flag.StringVar(&directory, "directory", "./Godeps", "path to Godeps directory")
	flag.Parse()
	excludes := flag.Args()

	// Load file
	deps, err := loadGodeps(directory)
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup
	for _, e := range excludes {
		deps.filter(e)
		dir := path.Join(directory, "_workspace/src", e)
		err := os.RemoveAll(dir)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Save
	deps.save(directory)
}
