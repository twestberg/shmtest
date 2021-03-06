package shmtest

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestShm(t *testing.T) {
	//RunShell("df")
	result, err := RunShell("df -BM /dev/shm| grep shm")
	if err != nil {
		t.Fatalf("could not run shell: %v", err)
	}
	parts := strings.Fields(result)
	if len(parts) < 2 {
		t.Errorf("Not enough parts in df output %s", result)
	}
	size := parts[1]
	megabytes := 0
	fmt.Sscanf(size, "%dM", &megabytes)
	if megabytes < 900 {
		t.Errorf("shm size is too small: %s", size)
	}
	fmt.Printf("Size is %dM, %s\n", megabytes, size)
}

func RunShell(shell_script string) (out string, err error) {
	var shout []byte

	shout, err = exec.Command("/bin/bash", "-c", shell_script).CombinedOutput()
	out = string(shout)
	fmt.Printf("command output: %s\n", out)
	return
}
