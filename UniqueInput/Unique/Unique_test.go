package main

import (
	"bytes"
	"string"
	"testing"
)

var testOk = `1
2
3
4
5
`
var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := buffio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)

	err := uniq()
	if err != nil {
		t.Errorf(("test for Ok failed -error"), args)
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("test for Ok failed - results  not match\n %v %v", result, testOkResult)
	}
}

var testFail = `1
2
1`

func testForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("test for Ok -failed- error %v", err)
	}
}
