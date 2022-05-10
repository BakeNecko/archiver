package cmd

import (
	"archiver/lib/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

// TODO: take extension from file
const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer CloseFile(r)

	data, err := io.ReadAll(r) // TODO: Add threading
	if err != nil {
		handleErr(err)
	}

	unpacked := vlc.Decode(string(data))

	err = os.WriteFile(unpackedFileName(filePath), []byte(unpacked), 0644)
	if err != nil {
		handleErr(err)
	}
}

// TODO: refactor this
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension

}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}