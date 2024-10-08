package ssh

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateSSHKey(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "machine-test-")
	if err != nil {
		t.Fatal(err)
	}

	filename := filepath.Join(tmpDir, "sshkey")

	if err := GenerateSSHKey(filename); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filename); err != nil {
		t.Fatalf("expected ssh key at %s", filename)
	}

	// cleanup
	_ = os.RemoveAll(tmpDir)
}
