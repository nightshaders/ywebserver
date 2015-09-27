package test_helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const TESTING_PWD = "TESTING_PWD"

func Pwd() string {
	return os.Getenv(TESTING_PWD)
}

func TestMain(m *testing.M) {
	if Pwd() == "" {
		msg := fmt.Sprintf(
			"Need to set %s, typically done with `export %s=$(pwd)`",
			TESTING_PWD, TESTING_PWD)
		panic(msg)
	} else {
		os.Exit(m.Run())
	}
}

func Dump(err error, parms ...interface{}) {
	for _, e := range parms {
		fmt.Println(e)
	}
	fmt.Println(err)
}

func Join(root, file string) string {
	return filepath.Join(Pwd(), root, file)
}
