package foldercompare

import (
	"os"
	"path/filepath"
	"strings"
)

// FolderSnapshot keeps the state of a folder
type FolderSnapshot struct {
	path     string
	filesMap map[string]os.FileInfo
	// size     int
}

// CreateSnapshot creates a snapshot of a folder
func CreateSnapshot(folderPath string) (*FolderSnapshot, error) {
	// get the absolute path
	var absPath, err1 = filepath.Abs(folderPath)
	if err1 != nil {
		return nil, err1
	}

	// initialize the snapshot
	var snapshot = FolderSnapshot{
		path:     absPath,
		filesMap: map[string]os.FileInfo{},
	}

	// get folder info
	// filepath.Walk flattens a directory to /dir/subdir/filename
	var err2 = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		snapshot.filesMap[strings.Replace(path, absPath, "", -1)] = info
		return nil
	})

	if err2 != nil {
		return nil, err2
	}

	return &snapshot, nil
}

// CompareSnapshots compares 2 folder states
func CompareSnapshots(firstFolder FolderSnapshot, secondFolder FolderSnapshot) float32 {
	var similiarCount float32
	for p1, f := range firstFolder.filesMap {
		var elem, ok = secondFolder.filesMap[p1]

		if ok && f.Size() == elem.Size() {
			similiarCount++
		}
	}

	// return 1 - ((similiarCount * 2) / float32(len(firstFolder.filesMap)+len(secondFolder.filesMap)))
	return 1 - similiarCount/float32(maxInt(len(firstFolder.filesMap), len(secondFolder.filesMap)))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
