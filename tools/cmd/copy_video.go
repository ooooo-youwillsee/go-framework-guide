package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
)

var (
	extTypes = map[string]struct{}{
		".avi": {},
		".mp4": {},
	}
	defaultFileFilter = func(filePath string) bool {
		extType := path.Ext(filePath)
		_, ok := extTypes[extType]
		return ok
	}
	verbose = false
	cnt     = 1

	copyVideoCmd = &cobra.Command{
		Use:   "copy_video src_dir dest_dir",
		Short: "copy all video from [src_dir] to [dest_dir] by recursively",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			srcDir, destDir := args[0], args[1]
			copyVideos(srcDir, destDir, defaultFileFilter)
			return nil
		},
	}
)

func init() {
	copyVideoCmd.Flags().BoolVarP(&verbose, "verbose", "v", true, "show verbose")
}

type FileMatcher = func(filePath string) bool

func copyVideos(srcDir, destDir string, fileMatch FileMatcher) error {
	err := checkDirExist(destDir)
	if err != nil {
		return err
	}

	dirEntries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	sort.Slice(dirEntries, func(i, j int) bool {
		return dirEntries[i].Name() < dirEntries[j].Name()
	})

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			copyVideos(path.Join(srcDir, dirEntry.Name()), destDir, fileMatch)
			continue
		}
		filePath := path.Join(srcDir, dirEntry.Name())
		if fileMatch(filePath) {
			destF, _ := os.Create(path.Join(destDir, strconv.Itoa(cnt)+path.Ext(filePath)))
			srcF, _ := os.Open(filePath)
			io.Copy(destF, srcF)
			cnt += 1
		}
	}

	return nil
}

func checkDirExist(dir string) error {
	stat, err := os.Stat(dir)
	if err == nil && !stat.IsDir() {
		os.Remove(dir)
		os.MkdirAll(dir, os.ModeDir)
		return nil
	}

	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(dir, os.ModeDir)
		return nil
	}
	return err
}
