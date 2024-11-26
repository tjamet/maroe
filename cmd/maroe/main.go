package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/bramvdbogaerde/go-scp"
	"github.com/pkg/sftp"
	"github.com/spf13/afero"
	"github.com/spf13/afero/sftpfs"
	"github.com/tjamet/maroe/pkg/copy"
	"github.com/tjamet/maroe/pkg/writer"
	"github.com/tjamet/ssh-client-config"
	"golang.org/x/crypto/ssh"

	"github.com/docopt/docopt-go"
)

const help = `Usage: importer --input=<Folder> --output=<Folder>

Options:
	--input=<Folder>, -i    The source folder (typically the path to the card) to be synchronised
	--output=<Folder>, -o   The destination folder where photos will be named uniquely
`

func loadSSHConfig(host string) (string, *ssh.ClientConfig, error) {
	sshclientconfig.Log = func(level int, format string, args ...interface{}) {
		if level < 5 {
			log.Printf(format, args...)
		}
	}
	if host == "" {
		return "", nil, fmt.Errorf("invalid URL %s. Missing host name", host)
	}

	cfg := sshclientconfig.NewSSHClientConfig("")
	sshcfg, err := cfg.SSHClientConfig(host)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get SSH client config: %v", err)
	}
	host = cfg.DialAddr(host)
	return host, sshcfg, nil
}

func main() {
	args, err := docopt.Parse(help, os.Args[1:], true, "0.0.0", false, true)
	if err != nil {
		log.Fatal(err)
	}

	u, err := url.Parse(args["--output"].(string))
	if err != nil {
		log.Fatalf("Invalid URL %s: %v", args["--output"].(string), err)
	}
	var w writer.Writer
	switch u.Scheme {
	case "scp", "ssh":
		host, sshcfg, err := loadSSHConfig(u.Host)
		if err != nil {
			log.Fatalf("Failed to get SSH client config: %v", err)
		}

		scpClient := scp.NewClient(host, sshcfg)
		err = scpClient.Connect()
		if err != nil {
			log.Fatalf("Failed to connect to %s: %v", u.Host, err)
		}
		defer scpClient.Close()

		w = writer.NewSCPWriter(&scpClient)
		// abs paths should look like ssh://host//abs/path
		// rel paths should look like ssh://host/rel/path
		u.Path = u.Path[1:]

	case "sftp":
		host, sshcfg, err := loadSSHConfig(u.Host)
		if err != nil {
			log.Fatalf("Failed to get SSH client config: %v", err)
		}

		sshClient, err := ssh.Dial("tcp", host, sshcfg)
		if err != nil {
			log.Fatalf("Failed to connect to %s: %v", u.Host, err)
		}

		sftpClient, err := sftp.NewClient(sshClient)
		if err != nil {
			log.Fatalf("Failed to create SFTP client: %v", err)
		}

		w = writer.NewFSWriter(sftpfs.New(sftpClient))
		// abs paths should look like sftp://host//abs/path
		// rel paths should look like sftp://host/rel/path
		u.Path = u.Path[1:]

	case "":
		w = writer.NewFSWriter(afero.NewOsFs())
	default:
		log.Fatalf("Protocol %s not implemented for output URL %s", u.Scheme, args["--output"].(string))
	}
	input := args["--input"].(string)
	err = copy.NewCopier(w).Copy(u.Path, input)
	if err != nil {
		log.Fatalf("completed with errors")
	}
}
