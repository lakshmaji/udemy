package reader

import (
	"io/fs"
	"os"
	"testing/fstest"
)

type fileSystemMock struct {
}

// FIXME: not using this any where. This can be removed.
func MockNewFile() fs.FS {
	return &fileSystemMock{}
}

func (f fileSystemMock) Open(name string) (fs.File, error) {
	const (
		firstBody  = "Post 1\nDescription"
		secondBody = "Post 2 \n Description"
	)

	mfs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
		"input.txt":       {Data: []byte("ADD BANANA")},
	}

	dir, err := fs.ReadDir(mfs, ".")
	if err != nil {
		return nil, err
	}

	for _, f := range dir {
		// post, err := getPost(fileSystem, f)
		postFile, err := mfs.Open(f.Name())

		if err != nil {
			//todo: needs clarification, should we totally fail if one file fails? or just ignore?
			return nil, err
		}
		defer postFile.Close()
		return postFile, nil
	}

	return nil, &fs.PathError{
		Op:   "read",
		Path: name,
		Err:  os.ErrNotExist,
	}
}
