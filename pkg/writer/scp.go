package writer

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	scp "github.com/bramvdbogaerde/go-scp"

	"github.com/tjamet/maroe/pkg/picture"
)

type SCPWriter struct {
	client *scp.Client
}

var _ Writer = &SCPWriter{}

func NewSCPWriter(client *scp.Client) *SCPWriter {
	return &SCPWriter{client: client}
}

func (s *SCPWriter) destinationExists(destination string) bool {
	session, err := s.client.SSHClient().NewSession()
	if err != nil {
		return false
	}
	defer session.Close()
	err = session.Run("test -f " + destination)
	if err != nil {
		return false
	}
	return true
}

func (s *SCPWriter) move(source, target string) error {
	session, err := s.client.SSHClient().NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	err = session.Run(fmt.Sprintf("mv '%s' '%s'", source, target))
	if err != nil {
		return err
	}
	return nil
}

func (s *SCPWriter) Write(photo *picture.Picture, destination string) (string, error) {
	outPath := destPath(photo, destination)
	if s.destinationExists(outPath) {
		return outPath, fmt.Errorf("path %s already exist", outPath)
	}
	tmpPath := path.Join(filepath.Dir(outPath), "."+filepath.Base(outPath))
	d := filepath.Dir(outPath)
	session, err := s.client.SSHClient().NewSession()
	if err != nil {
		return outPath, err
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		return outPath, err
	}
	stdin, err := session.StdinPipe()
	if err != nil {
		return outPath, err
	}
	defer stdin.Close()
	go io.Copy(os.Stdout, stdout)
	err = session.Run("mkdir -p " + d)
	if err != nil {
		return outPath, err
	}
	// err := s.client.Connect()
	// if err != nil {
	// 	return "", err
	// }
	// defer s.client.Close()
	err = s.client.CopyFile(context.Background(), photo.Reader, tmpPath, "0700")
	if err != nil {
		return outPath, err
	}

	err = s.move(tmpPath, outPath)
	if err != nil {
		return outPath, err
	}

	return outPath, nil
}
