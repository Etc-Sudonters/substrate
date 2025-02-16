package dontio

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileStdFlags uint64

// populates the passed *Std with files for out, err the returned cleanup
// function will never be nil and must always be called
func FileStd(std *Std, dir string) (func(), error) {
	var opened []*os.File
	cleanup := func() {
		for _, f := range opened {
			if err := f.Sync(); err != nil {
				panic(fmt.Errorf("failed to sync %s: %w", f.Name(), err))
			}
			if err := f.Close(); err != nil {
				panic(fmt.Errorf("failed to close %s: %w", f.Name(), err))
			}
		}
	}

	opening := []string{"out", "err"}

	for _, name := range opening {
		file, fileErr := os.Create(filepath.Join(dir, name))
		if fileErr != nil {
			return cleanup, fmt.Errorf("failed to initialize std%s: %w", name, fileErr)
		}
		opened = append(opened, file)
	}

	std.Out = opened[0]
	std.Err = opened[1]
	return cleanup, nil

}
