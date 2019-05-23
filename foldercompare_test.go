package foldercompare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSnapshot(t *testing.T) {
	var _, err = CreateSnapshot("examples/folder1")
	assert.Nil(t, err)
}

func TestCompareSnapshots(t *testing.T) {
	var f1, err1 = CreateSnapshot("examples/folder1")
	assert.Nil(t, err1)
	var f2, err2 = CreateSnapshot("examples/folder2")
	assert.Nil(t, err2)
	var f1Dup, err3 = CreateSnapshot("examples/folder1_dup")
	assert.Nil(t, err3)

	var res1 = CompareSnapshots(*f1, *f2)
	assert.Equal(t, float32(0.5), res1)

	var res2 = CompareSnapshots(*f1, *f1Dup)
	assert.Equal(t, float32(0), res2)
}

func TestMaxInt(t *testing.T) {
	assert.Equal(t, 2, maxInt(1, 2))
}
