package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var ignoredDirs = map[string]bool{
	"vendor":       true,
	"node_modules": true,
	".git":         true,
	"bin":          true,
	"obj":          true,
}

// FileTree represents the structure of directories and files
type FileTree struct {
	Name     string     `yaml:"name"`
	Type     string     `yaml:"type"`
	Children []FileTree `yaml:"children,omitempty"`
	Content  string     `yaml:"content,omitempty"`
}

func copyToClipboard(text string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = bytes.NewBufferString(text)
	return cmd.Run()
}

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	tree := FileTree{Name: filepath.Base(root), Type: "directory"}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && ignoredDirs[info.Name()] {
			return filepath.SkipDir
		}
		addToTree(&tree, path, info, root)
		return nil
	})

	output, err := yaml.Marshal(&tree)
	if err != nil {
		fmt.Printf("Error marshalling to YAML: %v\n", err)
		return
	}
	//fmt.Println(string(output)) // Replace this
	if err := copyToClipboard(string(output)); err != nil {
		fmt.Printf("Failed to copy to clipboard: %v\n", err)
	} else {
		fmt.Println("Output copied to clipboard.")
	}

}

func addToTree(root *FileTree, path string, info os.FileInfo, base string) {
	relPath, _ := filepath.Rel(base, path)
	parts := strings.Split(relPath, string(filepath.Separator))
	current := root

	for i, part := range parts {
		found := false
		for j := range current.Children {
			if current.Children[j].Name == part {
				current = &current.Children[j]
				found = true
				break
			}
		}
		if !found {
			newNode := FileTree{Name: part}
			if i == len(parts)-1 && !info.IsDir() {
				newNode.Type = "file"
				if filepath.Ext(path) == ".go" {
					content, _ := ioutil.ReadFile(path)
					newNode.Content = "|\n" + indentContent(string(content), "  ")
				}
			} else {
				newNode.Type = "directory"
			}
			current.Children = append(current.Children, newNode)
			current = &current.Children[len(current.Children)-1]
		}
	}
}

func indentContent(content, indent string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = indent + line
	}
	return strings.Join(lines, "\n")
}
