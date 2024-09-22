package preset

import (
	"cheyne.nz/unscramble/pkg/bag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

const expected = `{"A": 9,"B": 2,"C": 2,"D": 4,"E": 12,"F": 2,"G": 3,"H": 2,"I": 9,"J": 1,"K": 1,"L": 4,"M": 2,"N": 6,"O": 8,"P": 2,"Q": 1,"R": 6,"S": 4,"T": 6,"U": 4,"V": 2,"W": 2,"X": 1,"Y": 2,"Z": 1,"_": 2}`

func TestDefaultTileSet_Serialize(t *testing.T) {
	exampleBag := NewDefaultTileSet()
	contents, err := exampleBag.ToJSON()
	assert.Nil(t, err)
	require.JSONEq(t, expected, string(contents))
}

func TestDefaultTileSet_ExportImport(t *testing.T) {
	exampleBag := NewDefaultTileSet()

	file := filepath.Join(t.TempDir(), "example.json")
	exampleBag.ExportJson(file)

	importedBag, err := bag.ImportJson(file)
	contents, err := importedBag.ToJSON()
	assert.Nil(t, err)
	require.JSONEq(t, expected, string(contents))
}

func TestDefaultTileSet_Import(t *testing.T) {
	file := filepath.Join("../../..", "testdata", "letter_bag.json")

	importedBag, err := bag.ImportJson(file)
	contents, err := importedBag.ToJSON()
	assert.Nil(t, err)
	require.JSONEq(t, expected, string(contents))
}

func TestDefaultTileSet_ImportInvalid(t *testing.T) {
	file := filepath.Join("../../..", "testdata", "invalid_letter_bag.json")
	_, err := bag.ImportJson(file)
	assert.NotNil(t, err)
}

func TestDefaultTileSet_ImportInvalid2(t *testing.T) {
	file := filepath.Join("../../..", "testdata", "invalid_letter_bag_2.json")
	_, err := bag.ImportJson(file)
	assert.NotNil(t, err)
}
