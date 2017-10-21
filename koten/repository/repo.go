package repository

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
)

var workingDir string = "repos"

func init() {
	if err := os.Mkdir(workingDir, 0755); err != nil {
		fmt.Println("warning:", err)
	}
}

type Entry interface {
	Description()
}

type Directory struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Children []Entry `json:"children,omitempty"`
}

func (d *Directory) Description() {
	for _, entry := range d.Children {
		entry.Description()
	}
}

func (d *Directory) Add(entries []Entry) {
	d.Children = append(d.Children, entries...)
}

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (f *File) Description() {
	fmt.Println(f.Path)
}

func Walk(root string, ignores []string) Entry {
	dir := &Directory{Name: root, Path: filepath.Join(workingDir, root)}
	entries, _ := walk(dir.Path, ignores)
	dir.Add(entries)
	//dir.Description()
	return dir
}

func walk(path string, ignores []string) ([]Entry, error) {
	entries := []Entry{}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)

	for _, name := range names {
		if isIgnore(ignores, name) {
			continue
		}

		filename := filepath.Join(path, name)
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			fmt.Println("Error")
		} else {
			if fileInfo.IsDir() {
				dir := &Directory{Name: name, Path: filename}
				_entries, _ := walk(filename, ignores)
				dir.Add(_entries)
				entries = append(entries, dir)
			} else {
				file := &File{Name: name, Path: filename}
				entries = append(entries, file)
			}
		}
	}

	return entries, nil
}

func isIgnore(ignores []string, file string) bool {
	for _, ign := range ignores {
		if ign == file {
			return true
		}
	}
	return false
}

func Clone() {
	cmd := exec.Command("git", "clone", "git@github.com:terut/pebbles.git")
	cmd.Dir = workingDir
	cmd.Run()
}

func Pull() {
	cmd := exec.Command("git", "pull")
	cmd.Dir = workingDir
	cmd.Run()
}
