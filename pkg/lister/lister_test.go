package lister_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjamet/maroe/pkg/lister"
)

func TestFindAllPhotos(t *testing.T) {
	l := lister.Lister{}
	photos := l.Start("test")
	count := 0
	t.Run("All photos returned have valid extension", func(t *testing.T) {
		for p := range photos {
			if p.Extension != ".JPG" {
				assert.Equal(t, ".jpeg", p.Extension)
			}
			count++
		}
	})
	t.Run("All photos have been seeked", func(t *testing.T) {
		assert.Equal(t, 4, count)
	})
}
