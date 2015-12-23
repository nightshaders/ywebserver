package server

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServer(t *testing.T) {

	Convey("Should have name: 'World'", t, func() {
		name := Name()
		So(name, ShouldEqual, "World")
	})
}
