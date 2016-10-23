package crush

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

var helloWorldCrushed = []byte(`package main;import"fmt";func main(){fmt.Println("Hello, world!")}`)

func Test_HelloWorld_OutputIsHelloWorld(t *testing.T) {
	assert.Equal(t, "Hello, world!\n", getOutput(t, helloWorldBytes(t)))
}

func Test_PreCrushedHelloWorld_OutputIsHelloWorld(t *testing.T) {
	assert.Equal(t, "Hello, world!\n", getOutput(t, helloWorldCrushed))
}

func Test_HelloWorldCrushed_OutputIsHelloWorld(t *testing.T) {
	assert.Equal(t, "Hello, world!\n", getOutputFromCrushed(t, helloWorldBytes(t)))
}

func helloWorldBytes(t *testing.T) []byte {
	helloWorld, err := ioutil.ReadFile("./test_assets/hello_world.go")
	if err != nil {
		t.Fatal(err)
	}
	return helloWorld
}

func getOutputFromCrushed(t *testing.T, src []byte) string {
	crushed, err := Bytes(src)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(crushed))
	out, err := getOutputWithErr(crushed)
	if err != nil {
		t.Fatalf("run failed: %s: %s", err, string(crushed))
	}
	return out
}

func getOutput(t *testing.T, src []byte) string {
	out, err := getOutputWithErr(src)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func getOutputWithErr(src []byte) (string, error) {
	err := ioutil.WriteFile("./testfile.go", []byte(src), 0777)
	if err != nil {
		return "", err
	}
	defer os.Remove("./testfile.go")
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd := exec.Command("go", "run", "./testfile.go")
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err = cmd.Run()
	if err != nil {
		errBytes, rErr := ioutil.ReadAll(stderr)
		if rErr != nil {
			return "", rErr
		}
		return "", fmt.Errorf("error during 'go run': %s: %s", err, string(errBytes))
	}
	outBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	return string(outBytes), nil
}
