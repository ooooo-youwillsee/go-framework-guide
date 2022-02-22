package cmd

import (
	"path"
	"testing"
)

func TestCopyVideos(t *testing.T) {
	srcDir := "C:/Users/ooooo/Development/Code/Self/go-framework-guide/experimental-code"
	destDir := "C:/Users/ooooo/Development/Code/Self/go-framework-guide/experimental-code1"

	copyVideos(srcDir, destDir, func(filePath string) bool {
		ext := path.Ext(filePath)
		return ext == ".go"
	})
}
