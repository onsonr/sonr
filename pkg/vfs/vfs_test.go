package vfs_test

import (
	"testing"

	"github.com/di-dao/sonr/pkg/vfs"
	"github.com/stretchr/testify/assert"
)

func TestNewVFS(t *testing.T) {
	name := "test_vfs"
	vfs := vfs.New(name)
	assert.NotNil(t, vfs)
	assert.Equal(t, name, vfs.Name())
}

func TestAdd(t *testing.T) {
	vfs := vfs.New("test_vfs")
	path := "/file1"
	data := []byte("test data")
	err := vfs.Add(path, data)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	vfs := vfs.New("test_vfs")
	path := "/file1"
	data := []byte("test data")
	err := vfs.Add(path, data)
	assert.NoError(t, err)

	retrievedData, err := vfs.Get(path)
	assert.NoError(t, err)
	assert.Equal(t, data, retrievedData)

	_, err = vfs.Get("/nonexistent")
	assert.Error(t, err)
	assert.Equal(t, "file not found", err.Error())
}

func TestRm(t *testing.T) {
	vfs := vfs.New("test_vfs")
	path := "/file1"
	data := []byte("test data")
	err := vfs.Add(path, data)
	assert.NoError(t, err)

	err = vfs.Rm(path)
	assert.NoError(t, err)

	_, err = vfs.Get(path)
	assert.Error(t, err)
	assert.Equal(t, "file not found", err.Error())
}

func TestLs(t *testing.T) {
	vfs := vfs.New("test_vfs")
	path1 := "/file1"
	data1 := []byte("test data 1")
	err := vfs.Add(path1, data1)
	assert.NoError(t, err)

	path2 := "/file2"
	data2 := []byte("test data 2")
	err = vfs.Add(path2, data2)
	assert.NoError(t, err)

	files, err := vfs.Ls("")
	assert.NoError(t, err)
	assert.Len(t, files, 2)
	assert.Contains(t, files, path1)
	assert.Contains(t, files, path2)
}
