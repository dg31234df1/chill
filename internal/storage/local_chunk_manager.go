// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"context"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"go.uber.org/zap"
	"golang.org/x/exp/mmap"

	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/merr"
)

// LocalChunkManager is responsible for read and write local file.
type LocalChunkManager struct {
	localPath string
}

var _ ChunkManager = (*LocalChunkManager)(nil)

// NewLocalChunkManager create a new local manager object.
func NewLocalChunkManager(opts ...Option) *LocalChunkManager {
	c := newDefaultConfig()
	for _, opt := range opts {
		opt(c)
	}
	return &LocalChunkManager{
		localPath: c.rootPath,
	}
}

// RootPath returns lcm root path.
func (lcm *LocalChunkManager) RootPath() string {
	return lcm.localPath
}

// Path returns the path of local data if exists.
func (lcm *LocalChunkManager) Path(ctx context.Context, filePath string) (string, error) {
	exist, err := lcm.Exist(ctx, filePath)
	if err != nil {
		return "", err
	}

	if !exist {
		return "", merr.WrapErrIoKeyNotFound(filePath)
	}

	return filePath, nil
}

func (lcm *LocalChunkManager) Reader(ctx context.Context, filePath string) (FileReader, error) {
	return Open(filePath)
}

// Write writes the data to local storage.
func (lcm *LocalChunkManager) Write(ctx context.Context, filePath string, content []byte) error {
	dir := path.Dir(filePath)
	exist, err := lcm.Exist(ctx, dir)
	if err != nil {
		return err
	}
	if !exist {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return merr.WrapErrIoFailed(filePath, err)
		}
	}
	return WriteFile(filePath, content, os.ModePerm)
}

// MultiWrite writes the data to local storage.
func (lcm *LocalChunkManager) MultiWrite(ctx context.Context, contents map[string][]byte) error {
	var el error
	for filePath, content := range contents {
		err := lcm.Write(ctx, filePath, content)
		if err != nil {
			el = merr.Combine(el, errors.Wrapf(err, "write %s failed", filePath))
		}
	}
	return el
}

// Exist checks whether chunk is saved to local storage.
func (lcm *LocalChunkManager) Exist(ctx context.Context, filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, merr.WrapErrIoFailed(filePath, err)
	}
	return true, nil
}

// Read reads the local storage data if exists.
func (lcm *LocalChunkManager) Read(ctx context.Context, filePath string) ([]byte, error) {
	return ReadFile(filePath)
}

// MultiRead reads the local storage data if exists.
func (lcm *LocalChunkManager) MultiRead(ctx context.Context, filePaths []string) ([][]byte, error) {
	results := make([][]byte, len(filePaths))
	var el error
	for i, filePath := range filePaths {
		content, err := lcm.Read(ctx, filePath)
		if err != nil {
			el = merr.Combine(el, errors.Wrapf(err, "failed to read %s", filePath))
		}
		results[i] = content
	}
	return results, el
}

func (lcm *LocalChunkManager) ListWithPrefix(ctx context.Context, prefix string, recursive bool) ([]string, []time.Time, error) {
	var filePaths []string
	var modTimes []time.Time
	if recursive {
		dir := filepath.Dir(prefix)
		err := filepath.Walk(dir, func(filePath string, f os.FileInfo, err error) error {
			if strings.HasPrefix(filePath, prefix) && !f.IsDir() {
				filePaths = append(filePaths, filePath)
			}
			return nil
		})
		if err != nil {
			return nil, nil, err
		}
		for _, filePath := range filePaths {
			modTime, err2 := lcm.getModTime(filePath)
			if err2 != nil {
				return filePaths, nil, err2
			}
			modTimes = append(modTimes, modTime)
		}
		return filePaths, modTimes, nil
	}

	globPaths, err := filepath.Glob(prefix + "*")
	if err != nil {
		return nil, nil, err
	}
	filePaths = append(filePaths, globPaths...)
	for _, filePath := range filePaths {
		modTime, err2 := lcm.getModTime(filePath)
		if err2 != nil {
			return filePaths, nil, err2
		}
		modTimes = append(modTimes, modTime)
	}

	return filePaths, modTimes, nil
}

func (lcm *LocalChunkManager) ReadWithPrefix(ctx context.Context, prefix string) ([]string, [][]byte, error) {
	filePaths, _, err := lcm.ListWithPrefix(ctx, prefix, true)
	if err != nil {
		return nil, nil, err
	}
	result, err := lcm.MultiRead(ctx, filePaths)
	return filePaths, result, err
}

// ReadAt reads specific position data of local storage if exists.
func (lcm *LocalChunkManager) ReadAt(ctx context.Context, filePath string, off int64, length int64) ([]byte, error) {
	if off < 0 || length < 0 {
		return nil, io.EOF
	}

	file, err := Open(path.Clean(filePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	res := make([]byte, length)
	_, err = file.ReadAt(res, off)
	if err != nil {
		return nil, merr.WrapErrIoFailed(filePath, err)
	}
	return res, nil
}

func (lcm *LocalChunkManager) Mmap(ctx context.Context, filePath string) (*mmap.ReaderAt, error) {
	reader, err := mmap.Open(path.Clean(filePath))
	if errors.Is(err, os.ErrNotExist) {
		return nil, merr.WrapErrIoKeyNotFound(filePath, err.Error())
	}

	return reader, merr.WrapErrIoFailed(filePath, err)
}

func (lcm *LocalChunkManager) Size(ctx context.Context, filePath string) (int64, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return 0, merr.WrapErrIoKeyNotFound(filePath, err.Error())
		}
		return 0, merr.WrapErrIoFailed(filePath, err)
	}
	// get the size
	size := fi.Size()
	return size, nil
}

func (lcm *LocalChunkManager) Remove(ctx context.Context, filePath string) error {
	err := os.RemoveAll(filePath)
	return merr.WrapErrIoFailed(filePath, err)
}

func (lcm *LocalChunkManager) MultiRemove(ctx context.Context, filePaths []string) error {
	errors := make([]error, 0, len(filePaths))
	for _, filePath := range filePaths {
		err := lcm.Remove(ctx, filePath)
		errors = append(errors, err)
	}
	return merr.Combine(errors...)
}

func (lcm *LocalChunkManager) RemoveWithPrefix(ctx context.Context, prefix string) error {
	// If the prefix is empty string, the ListWithPrefix() will return all files under current process work folder,
	// MultiRemove() will delete all these files. This is a danger behavior, empty prefix is not allowed.
	if len(prefix) == 0 {
		errMsg := "empty prefix is not allowed for ChunkManager remove operation"
		log.Warn(errMsg)
		return merr.WrapErrParameterInvalidMsg(errMsg)
	}

	filePaths, _, err := lcm.ListWithPrefix(ctx, prefix, true)
	if err != nil {
		return err
	}

	return lcm.MultiRemove(ctx, filePaths)
}

func (lcm *LocalChunkManager) getModTime(filepath string) (time.Time, error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		log.Warn("stat fileinfo error",
			zap.String("filepath", filepath),
			zap.Error(err),
		)
		if os.IsNotExist(err) {
			return time.Time{}, merr.WrapErrIoKeyNotFound(filepath)
		}
		return time.Time{}, merr.WrapErrIoFailed(filepath, err)
	}

	return fi.ModTime(), nil
}
