package command

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/leonardodelira/go-git-clone/index"
	"github.com/spf13/cobra"
)

var idx index.Index
var err error
var untrackedFiles []string
var notStagedFiles []string

var Status = &cobra.Command{
	Use:   "status",
	Short: "show the stauts files",
	Run: func(cmd *cobra.Command, args []string) {
		idx, err = index.Build()
		if err != nil {
			log.Fatalf("could not build index: %v", err)
		}

		filepath.Walk(".", classifyFiles)

		red := color.New(color.FgRed).SprintFunc()

		if len(notStagedFiles) > 0 {
			fmt.Println("Changes not staged for commit:")
			fmt.Println("  (use \"fit add <file>\" to update what will be committed)")
			fmt.Println("    " + red(strings.Join(notStagedFiles, "\n    ")))
		}

		if len(untrackedFiles) > 0 {
			fmt.Println("untracked files:")
			fmt.Println("  (use \"fit add <file>\" to update what will be committed)")
			fmt.Println("    " + red(strings.Join(untrackedFiles, "\n    ")))
		}
	},
}

func classifyFiles(path string, fileInfo fs.FileInfo, err error) error {
	if fileInfo.IsDir() && (fileInfo.Name() == ".fit" || fileInfo.Name() == ".git") {
		return filepath.SkipDir
	}

	if fileInfo.IsDir() {
		return nil
	}

	if fileInfo.Name() == "." || fileInfo.Name() == ".." {
		return nil
	}

	if _, ok := idx.Objects[path]; !ok {
		untrackedFiles = append(untrackedFiles, path)
		return nil
	}

	if idx.Objects[path].LastModified.Before(fileInfo.ModTime()) {
		notStagedFiles = append(notStagedFiles, path)
		return nil
	}

	return nil
}
