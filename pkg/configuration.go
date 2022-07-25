package pkg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/natefinch/atomic"
	"go.mongodb.org/mongo-driver/mongo/address"
)

// Config fulfills the vault configuration
type Config struct {
	// Address is the address of the Vault server. This should be a complete
	// URL such as "http://vault.example.com". If you need a custom SSL
	// cert or want to enable insecure mode, you need to specify a custom
	// HttpClient.
	Address string

	HomeDir string

	ConfigPath string
}

func NewConfig(string address) (*Config) {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(fmt.Sprintf("error getting user's home directory: %v", err))
	}
	c := &Config{
		Address: address,
		HomeDir: homeDir
	}
	return c
}

// populateConfigurationPath figures out the vault path using homedir to get the user's
// home directory
func (c *Config) populateConfigurationPath() {
	c.configPath = filepath.Join(i.homeDir, ".vt")
}

func (i *Config) Path() string {
	return i.tokenPath
}

// Get gets the value of the stored token, if any
func (i *Config) Get() (string, error) {
	i.populateTokenPath()
	f, err := os.Open(i.tokenPath)
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}

// Store stores the value of the token to the file.  We always overwrite any
// existing file atomically to ensure that ownership and permissions are set
// appropriately.
func (i *InternalTokenHelper) Store(input string) error {
	i.populateTokenPath()
	tmpFile := i.tokenPath + ".tmp"
	f, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()
	defer os.Remove(tmpFile)

	_, err = io.WriteString(f, input)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	// We don't care so much about atomic writes here.  We're using this package
	// because we don't have a portable way of verifying that the target file
	// is owned by the correct user.  The simplest way of ensuring that is
	// to simply re-write it, and the simplest way to ensure that we don't
	// damage an existing working file due to error is the write-rename pattern.
	// os.Rename on Windows will return an error if the target already exists.
	return atomic.ReplaceFile(tmpFile, i.tokenPath)
}

// Erase erases the value of the token
func (i *InternalTokenHelper) Erase() error {
	i.populateTokenPath()
	if err := os.Remove(i.tokenPath); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}
