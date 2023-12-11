package command

import (
	"log"
	"os"

	"github.com/leonardodelira/go-git-clone/index"
	"github.com/leonardodelira/go-git-clone/objects"
	"github.com/spf13/cobra"
)

var Add = &cobra.Command{
	Use:   "add",
	Short: "add file to fit staging area",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("missing file argument")
		}

		for _, path := range args {
			file, err := os.OpenFile(path, os.O_RDONLY, 0655)
			if err != nil {
				log.Fatalf("cannot read the file: %s - err: %s", path, err)
			}
			defer file.Close()

			blob, err := objects.CreateBlob(file)
			if err != nil {
				log.Fatalf("error on create blob file: %s - err: %s", path, err)
			}
			defer blob.Close()

			err = index.AddBlob(file, blob)
			if err != nil {
				log.Fatalf("cannot add blob: %s", err)
			}
		}
	},
}
