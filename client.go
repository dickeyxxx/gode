package gode

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// The Node version to install.
// Override this by setting client.Version.
const DefaultNodeVersion = "1.1.0"

// The NPM version to install.
// Override this by setting client.Version.
const DefaultNpmVersion = "2.4.1"

// Client is the interface between Node and Go.
// It also setups up the Node environment if needed.
type Client struct {
	RootPath    string
	NodeVersion string
	NpmVersion  string
	Registry    string
}

// NewClient creates a new Client at the specified rootPath
// The Node installation can then be setup here with client.Setup()
func NewClient(rootPath string) *Client {
	client := &Client{
		RootPath:    rootPath,
		NodeVersion: DefaultNodeVersion,
		NpmVersion:  DefaultNpmVersion,
	}

	return client
}

// NodeBase returns the directory name for the node install
func (c *Client) NodeBase() string {
	switch {
	case runtime.GOARCH == "arm":
		return "iojs-v" + c.NodeVersion + "-" + runtime.GOOS + "-armv7l"
	case runtime.GOARCH == "386":
		return "iojs-v" + c.NodeVersion + "-" + runtime.GOOS + "-x86"
	default:
		return "iojs-v" + c.NodeVersion + "-" + runtime.GOOS + "-x64"
	}
}

func (c *Client) nodeURL() string {
	switch {
	case runtime.GOOS == "windows":
		return "https://iojs.org/dist/v" + c.NodeVersion + "/win-x86/iojs.exe"
	case runtime.GOARCH == "386":
		return "https://iojs.org/dist/v" + c.NodeVersion + "/iojs-v" + c.NodeVersion + "-" + runtime.GOOS + "-x86.tar.gz"
	case runtime.GOARCH == "arm":
		return "https://iojs.org/dist/v" + c.NodeVersion + "/iojs-v" + c.NodeVersion + "-" + runtime.GOOS + "-armv7l.tar.gz"
	default:
		return "https://iojs.org/dist/v" + c.NodeVersion + "/" + c.NodeBase() + ".tar.gz"
	}
}

func (c *Client) nodePath() string {
	switch {
	case runtime.GOOS == "windows":
		return filepath.Join(c.RootPath, c.NodeBase(), "bin", "node.exe")
	default:
		return filepath.Join(c.RootPath, c.NodeBase(), "bin", "node")
	}
}

func (c *Client) npmURL() string {
	return "http://github.com/npm/npm/archive/v" + c.NpmVersion + ".zip"
}

func (c *Client) npmPath() string {
	return filepath.Join(c.RootPath, c.NodeBase(), "lib", "node_modules", "npm", "cli.js")
}

func (c *Client) tmpDir(prefix string) string {
	root := filepath.Join(c.RootPath, "tmp")
	err := os.MkdirAll(root, 0755)
	if err != nil {
		panic(err)
	}
	dir, err := ioutil.TempDir(root, prefix)
	if err != nil {
		panic(err)
	}
	return dir
}
