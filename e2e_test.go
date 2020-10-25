package defaultcertifi_test

import (
	"context"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func Test_http_in_docker(t *testing.T) {
	if _, err := exec.LookPath("docker"); err != nil {
		t.Skip("docker not found")
	}
	if runtime.GOOS == "windows" {
		t.Skip("e2e test on windows is not supported")
	}
	ctx := context.Background()
	var version string
	_version := runtime.Version()
	if strings.Contains(_version, "devel") {
		version = "latest" // There is no image tag for tip in docker hub.
	} else {
		version = _version[2:] // go1.x.y -> 1.x.y
	}
	image := "johejo/go-defaultcertifi:" + version
	build := exec.CommandContext(ctx, "docker", "build", "--build-arg", "GO_VERSION="+version, "-t", image, ".")
	build.Stderr = os.Stderr
	build.Stdout = os.Stdout
	if err := build.Run(); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		rmi := exec.CommandContext(ctx, "docker", "rmi", image)
		if err := rmi.Run(); err != nil {
			t.Fatal(err)
		}
	})

	run := exec.CommandContext(ctx, "docker", "run", "--rm", image)
	run.Stderr = os.Stderr
	run.Stdout = os.Stdout
	if err := run.Run(); err != nil {
		t.Fatal(err)
	}
}
