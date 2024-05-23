package copy_test

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tjamet/maroe/pkg/copy"
	"github.com/tjamet/maroe/pkg/writer"
)

func TestCopy(t *testing.T) {
	os.Remove("test/output/file_8731553c4d39ace22bf0849681a9d76ac4e5fa5a.jpeg")
	err := copy.NewCopier(writer.NewFSWriter(afero.NewOsFs())).Copy("test/output", "test/input")
	assert.NoError(t, err)
	assert.FileExists(t, "test/output/file_8731553c4d39ace22bf0849681a9d76ac4e5fa5a.jpeg")
}
