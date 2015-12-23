package mime

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func relToTestingPwd(f string) string {
	cwd := os.Getenv("TESTING_PWD")
	return filepath.Join(cwd, f)
}

func TestMime(t *testing.T) {

	Convey("Should find many lines in actual mime file.", t, func() {
		filename := relToTestingPwd("files/mime.types")
		txt, err := fileToString(filename)
		So(txt, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Should parse lines a,b from string.", t, func() {
		lines := toLines("a\nb")
		So(len(lines), ShouldEqual, 2)
		So(lines[0], ShouldEqual, "a\n")
		So(lines[1], ShouldEqual, "b")
	})

	Convey("Should should parse 3 lines from string ending in new-line.", t, func() {
		lines := toLines("a\nb\n\n")
		So(len(lines), ShouldEqual, 3)
		So(lines[0], ShouldEqual, "a\n")
		So(lines[1], ShouldEqual, "b\n")
		So(lines[2], ShouldEqual, "\n")
	})

	Convey("Should have non-null MimeType instance after parsing.", t, func() {
		mimeTypes, err := Parse("file")
		So(err, ShouldBeNil)
		So(mimeTypes, ShouldNotBeNil)
	})

	Convey("Should find .css as text/css.", t, func() {
		mimeTypes, _ := Parse("file")
		So(mimeTypes.Get(".css"), ShouldEqual, "text/css")
	})

	//	Convey("Should find .js as text/javascript.", t, func() {
	//		mimeTypes, _ := Parse("file")
	//		So(mimeTypes.Get(".js"), ShouldEqual, "text/javascript")
	//	})

}
