package service

import (
	"gosftp/config"
	"io"
	"log"
	"path"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SftpService struct {
	client *sftp.Client
}

func NewSftpService(config *config.SftpConfig) (*SftpService, error) {
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	log.Print("Username:", config.Username)
	log.Print("Username:", config.Password)
	log.Print("Username:", config.Host)
	log.Print("Username:", config.Port)
	
	conn, err := ssh.Dial("tcp", config.Host+":"+config.Port, sshConfig)
	if err != nil {
		return nil, err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		conn.Close() // Close the connection if creating the SFTP client fails
		return nil, err
	}

	return &SftpService{
		client: client,
	}, nil
}

func (s *SftpService) UploadFile(src io.Reader, filename string, remoteDir string) error {
	remoteFile, err := s.client.Create(path.Join(remoteDir, filename))
	if err != nil {
		return err
	}
	defer remoteFile.Close()

	// Use a buffer for efficient IO operations
	buffer := make([]byte, 32*1024) // 32KB buffer, you can adjust this size based on your use case

	_, err = io.CopyBuffer(remoteFile, src, buffer)
	if err != nil {
		return err
	}

	return nil
}

func (s *SftpService) Close() error {
	return s.client.Close()
}