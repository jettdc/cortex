package values

import (
	"fmt"
	"github.com/jettdc/cortex/v2/utils"
	"os"
	"os/exec"
	"strings"
)

type VimValueWriter struct{}

func (v VimValueWriter) WriteValue() (string, error) {
	tmpFile, err := os.CreateTemp(utils.GetCortexPath(), "vim-input-*.txt")
	if err != nil {
		return "", err
	}

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	cmd := exec.Command("vim", "-c", "startinsert", tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("vim exited with error: %w", err)
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(content), "\n"), nil
}
