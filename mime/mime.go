package mime

import (
	"io/ioutil"
)

type MimeTypes struct {
	lookup map[string]mime
}

type mime struct {
	mime      string
	extension string
}

func (m *MimeTypes) Get(ext string) string {
	return "text/css"
}

func fileToString(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	return string(bytes), err
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func toLines(txt string) []string {
	lines := make([]string, 0)
	mark := 0
	for i, c := range txt {
		//		fmt.Println(c)
		if c == '\n' {
			line := txt[mark:max(mark+1, i+1)]
			lines = append(lines, line)
			mark = i + 1
			continue
		}
		if i == len(txt)-1 {
			line := txt[mark:max(mark+1, i+1)]
			//			fmt.Printf("last char [%d-%d, %c] line:'%s'\n", mark, i, c, line)
			lines = append(lines, line)
		}
	}
	return lines
}

func isIgnorableLine(line string) bool {
	return true
}

func toMimes(lines []string) []*mime {
	mimes := make([]*mime, 0)
	for _, line := range lines {
		if isIgnorableLine(line) {
			continue
		} else {
			// THis is wrong, must fix for tests
			m := &mime{
				mime:      line,
				extension: line,
			}
			mimes = append(mimes, m)
		}
	}
	return mimes
}

func toLookup(mimes []*mime) map[string]mime {
	return make(map[string]mime)
}

func Parse(file string) (*MimeTypes, error) {
	types := &MimeTypes{}
	if txt, err := fileToString(file); err == nil {
		lines := toLines(txt)
		mimes := toMimes(lines)
		types.lookup = toLookup(mimes)
	}
	return types, nil
}
