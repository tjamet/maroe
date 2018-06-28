package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	tools "github.com/tjamet/maroe/pkg/test-tools"
)

const (
	testDir      = "image/tiff/tests/raw"
	testTemplate = `
func Test{{.Provider}}{{.Name}}(t *testing.T){
	path := tools.DownloadTestRAW(t, "{{ .URL }}")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	{{ if .Exif.CreateDate -}}
	testCanReadValue(t, tiffImage, 0x0132, "{{ .Exif.CreateDate }}")
	{{- end }}
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}
`
)

var (
	linkRe              = regexp.MustCompilePOSIX(`href=["'][^"']*["']`)
	supportedExtensions = map[string]struct{}{
		".cr2": {},
		".dng": {},
		".nef": {},
		".nrw": {},
		".raf": {},
	}
	knownFailingImages = map[string]struct{}{
		"http://www.rawsamples.ch/raws/nikon/SCANNER_NIKON_LS5000.DNG":       {},
		"http://www.rawsamples.ch/raws/nikon/e5700/RAW_NIKON_E5700_SRGB.NEF": {},
		"http://www.rawsamples.ch/raws/nikon/d60/RAW_NIKON_D60.NEF":          {},
		"http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D70S.NEF":             {},
		"http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D3300.NEF":            {},
		"http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D5000.NEF":            {},
		"http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D7000.NEF":            {},
		"http://www.rawsamples.ch/rawss/fuji/RAW_FUJIFILM_X-T1.RAF":          {},
	}
)

type testTemplateValues struct {
	Provider string
	Name     string
	URL      string
	Exif     map[string]interface{}
}

func main() {
	files := map[string]map[string]interface{}{}

	outFile, err := os.Create("image/tiff/zz_generated_test.go")
	if err != nil {
		log.Fatalf("Failed to create test file: %v", err)
	}
	defer func() {
		outFile.Close()
		exec.Command("gofmt", "-w", "./image/tiff/zz_generated_test.go").Run()
	}()
	outFile.WriteString(`package tiff_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tjamet/goraw"
	_ "github.com/tjamet/goraw/all"
	"github.com/tjamet/maroe/image/tiff"
	_ "github.com/tjamet/maroe/image/tiff/exif"
	tools "github.com/tjamet/maroe/pkg/test-tools"
)
`)
	processed := map[string]struct{}{}

	for _, provider := range []string{"canon", "nikon", "fuji"} {
		tpl, err := template.New(provider).Parse(testTemplate)
		if err != nil {
			log.Fatalf("Failed to parse template: %v", err)
		}
		tpl.Funcs(sprig.FuncMap())

		resp, err := http.Get("http://www.rawsamples.ch/index.php/en/" + provider)
		if err != nil {
			log.Fatalf("Failed to get sample images: %s", err)
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalf("Failed to get sample images: %s", err)
		}
		fmt.Println(provider)
		for _, linkBytes := range linkRe.FindAll(b, -1) {
			link := string(linkBytes[6 : len(linkBytes)-1])
			if _, notACamera := knownFailingImages[link]; notACamera {
				fmt.Println("Skip not a camera file", link)
				continue
			}
			if _, alreadyProcessed := processed[link]; alreadyProcessed {
				continue
			}
			processed[link] = struct{}{}

			_, supported := supportedExtensions[strings.ToLower(filepath.Ext(link))]
			if !supported {
				fmt.Println("Skip unsupported extension", link)
				continue
			}
			if strings.HasPrefix(link, "http://www.rawsamples.ch/raws") {
				destFile, err := tools.DownloadRAW(link)
				if err != nil {
					log.Fatalf("Failed to download %s: %v", link, err)
				}
				out, err := exec.Command("exiftool", "-json", destFile).Output()
				if err == nil {
					data := []map[string]interface{}{}
					json.Unmarshal(out, &data)
					u, err := url.Parse(link)
					if err != nil {
						log.Fatalf("Failed to parse URL %s: %v", link, err)
					}
					testName := u.Path
					for _, c := range []string{"/", ".", ":", "-"} {
						testName = strings.Replace(testName, c, "", -1)
					}

					testTemplateData := testTemplateValues{
						Provider: strings.ToTitle(provider),
						Name:     testName,
						URL:      link,
						Exif:     data[0],
					}
					tpl.Execute(outFile, testTemplateData)

					files[link] = data[0]
				}
			}
		}
	}
}
