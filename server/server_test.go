package server

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServer(t *testing.T) {

	Convey("Should have name: 'World'", t, func() {
		name := Name()
		So(name, ShouldEqual, "World")
	})
}
