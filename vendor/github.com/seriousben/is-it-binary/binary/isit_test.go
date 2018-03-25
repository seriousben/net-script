package binary

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBinaryBuffer(t *testing.T) {
	tests := []struct {
		str       string
		expectErr bool
		isBinary  bool
	}{
		{"\x00", false, true},
		{"content\nfoo\nbar\x00", false, true},
		{"content\nfoo\nbar", false, false},
	}

	for _, tt := range tests {
		b := []byte(tt.str)
		binary, err := IsBinaryBuffer(b)
		if tt.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		if tt.isBinary {
			assert.True(t, binary)
		} else {
			assert.False(t, binary)
		}
	}
}

func TestIsBinaryFile(t *testing.T) {
	tests := []struct {
		filepath  string
		expectErr bool
		isBinary  bool
	}{
		{"whoamI", false, true},
		{"image.png", false, true},
		{"text.txt", false, false},
		{"not.existing", true, false},
	}

	for _, tt := range tests {
		binary, err := IsBinaryFile(filepath.Join("testdata", tt.filepath))
		if tt.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		if tt.isBinary {
			assert.True(t, binary)
		} else {
			assert.False(t, binary)
		}
	}

}
