package parser

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/opsidian/parsley/parsley"
	"github.com/opsidian/parsley/text"
	"github.com/seriousben/net-script/internal/types"
)

func TestCommandHeaderParser(t *testing.T) {
	parser := commandHeaderParser()
	s := parsley.NewSentence(parser)

	tests := []struct {
		toParse  string
		expected commandHeader
	}{
		{
			"GET http://example.com",
			commandHeader{Method: "GET", URL: "http://example.com"},
		},
		{
			"POST http://example.com",
			commandHeader{Method: "POST", URL: "http://example.com"},
		},
	}

	for _, test := range tests {
		t.Run(test.toParse, func(t *testing.T) {
			raw, _, err := s.Evaluate(text.NewReader([]byte(test.toParse), "", false), nil)
			assert.NilError(t, err)

			parsed := raw.(commandHeader)

			assert.DeepEqual(t, parsed, test.expected)
		})
	}
}

func TestHeaderParser(t *testing.T) {
	parser := headerParser()
	s := parsley.NewSentence(parser)

	tests := []struct {
		toParse  string
		expected []string
	}{
		{
			"Content-Type: application/json",
			[]string{"Content-Type", "application/json"},
		},
		{
			"Content-Type: text/html; charset=UTF-8",
			[]string{"Content-Type", "text/html; charset=UTF-8"},
		},
	}

	for _, test := range tests {
		t.Run(test.toParse, func(t *testing.T) {
			raw, _, err := s.Evaluate(text.NewReader([]byte(test.toParse), "", false), nil)
			assert.NilError(t, err)

			parsed := raw.([]string)

			assert.DeepEqual(t, parsed, test.expected)
		})
	}
}

func TestHeadersParser(t *testing.T) {
	parser := headersParser()
	s := parsley.NewSentence(parser)

	tests := []struct {
		name     string
		toParse  []string
		expected map[string][]string
	}{
		{
			"multiple",
			[]string{
				"Size: 1337",
				"Content-Type: text/html; charset=UTF-8",
			},
			map[string][]string{
				"Size":         []string{"1337"},
				"Content-Type": []string{"text/html; charset=UTF-8"},
			},
		},
		{
			"multiple values for same header",
			[]string{
				"Size: 1337",
				"Size: 7777",
				"Content-Type: text/html; charset=UTF-8",
			},
			map[string][]string{
				"Size":         []string{"1337", "7777"},
				"Content-Type": []string{"text/html; charset=UTF-8"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			raw, _, err := s.Evaluate(text.NewReader([]byte(strings.Join(test.toParse, "\n")), "", false), nil)
			assert.NilError(t, err)

			parsed := raw.(http.Header)

			assert.DeepEqual(t, parsed, http.Header(test.expected))
		})
	}
}

func TestBodyParser(t *testing.T) {
	parser := bodyParser()
	s := parsley.NewSentence(parser)

	tests := []struct {
		name    string
		toParse []string
	}{
		{
			"multiple",
			[]string{
				"line 1",
				"line 2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := []byte(strings.Join(test.toParse, "\n"))
			raw, _, err := s.Evaluate(text.NewReader(b, "", false), nil)
			assert.NilError(t, err)

			parsed := raw.([]byte)

			assert.DeepEqual(t, string(parsed), string(b))
		})
	}
}

var exampleCommands = map[string]struct {
	lines   []string
	command types.Command
}{
	"simple GET": {
		[]string{
			"# This is a comment",
			"GET http://example.com",
		},
		types.Command{
			Comment: "This is a comment",
			Method:  "GET",
			URL:     "http://example.com",
			Headers: http.Header{},
			Body:    []byte{},
		},
	},
	"simple PUT": {
		[]string{
			"# Put",
			"PUT https://example.com/",
		},
		types.Command{
			Comment: "Put",
			Method:  "PUT",
			URL:     "https://example.com/",
			Headers: http.Header{},
			Body:    []byte{},
		},
	},
	"simple delete": {
		[]string{
			"# Delete",
			"DELETE https://example.com/",
		},
		types.Command{
			Comment: "Delete",
			Method:  "DELETE",
			URL:     "https://example.com/",
			Headers: http.Header{},
			Body:    []byte{},
		},
	},
	/*
		"multi-line comment and command only": {
			[]string{
				"# This is also a comment",
				"GET http://example.com",
			},
			types.Command{
				Comment: "This is a comment\nThis is also a comment",
				Method:  "GET",
				URL:     "http://example.com",
				Headers: http.Header{},
				Body:    []byte{},
			},
		},
	*/
	"command and single header": {
		[]string{
			"# GET example",
			"GET https://example.com/",
			"Content-Type: application/json",
		},
		types.Command{
			Comment: "GET example",
			Method:  "GET",
			URL:     "https://example.com/",
			Headers: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: []byte{},
		},
	},
	"command and multiple headers": {
		[]string{
			"# GET example",
			"GET https://example.com/",
			"Content-Type: application/json",
			"X-Special-Header: very-special",
		},
		types.Command{
			Comment: "GET example",
			Method:  "GET",
			URL:     "https://example.com/",
			Headers: http.Header{
				"Content-Type":     []string{"application/json"},
				"X-Special-Header": []string{"very-special"},
			},
			Body: []byte{},
		},
	},
	"command and body (no headers)": {
		[]string{
			"# command and body (no headers)",
			"POST https://example.com/",
			"",
			"BODY",
		},
		types.Command{
			Comment: "command and body (no headers)",
			Method:  "POST",
			URL:     "https://example.com/",
			Headers: http.Header{},
			Body:    []byte("BODY"),
		},
	},
	"command, headers and body": {
		[]string{
			"# POST example",
			"POST https://example.com/",
			"Content-Type: application/json",
			"X-Special-Header: very-special",
			"",
			"LINE 1",
			"LINE 2",
			"#LINE 3",
		},
		types.Command{
			Comment: "POST example",
			Method:  "POST",
			URL:     "https://example.com/",
			Headers: http.Header{
				"Content-Type":     []string{"application/json"},
				"X-Special-Header": []string{"very-special"},
			},
			Body: []byte(strings.Join([]string{"LINE 1", "LINE 2", "#LINE 3"}, "\n")),
		},
	},
}

func TestCommandParser(t *testing.T) {
	parser := commandParser()
	s := parsley.NewSentence(parser)

	for name, test := range exampleCommands {
		t.Run(name, func(t *testing.T) {
			lines := strings.Join(test.lines, "\n")
			fmt.Println("(", name, ")\n", lines)
			b := []byte(lines)
			raw, _, err := s.Evaluate(text.NewReader(b, "", false), nil)
			assert.NilError(t, err)

			parsed := raw.(types.Command)

			assert.DeepEqual(t, parsed, test.command)
		})
	}
}

func TestParse(t *testing.T) {
	type scenario struct {
		name     string
		toParse  []string
		expected []types.Command
	}

	tests := []scenario{}

	/*for name, exampleCommand := range exampleCommands {
		tests = append(tests, scenario{
			name:     fmt.Sprintf("single - %s", name),
			toParse:  exampleCommand.lines,
			expected: []types.Command{exampleCommand.command},
		})
	}*/

	allTypesTest := scenario{
		name:     "multiple - all command types",
		toParse:  []string{},
		expected: []types.Command{},
	}
	for _, exampleCommand := range exampleCommands {
		allTypesTest.toParse = append(allTypesTest.toParse, exampleCommand.lines...)
		allTypesTest.toParse = append(allTypesTest.toParse, "\n")
		allTypesTest.expected = append(allTypesTest.expected, exampleCommand.command)
	}
	tests = append(tests, allTypesTest)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file := strings.Join(test.toParse, "\n")
			file = strings.Replace(file, "\n\n#", "\n#", -1)
			fmt.Println("START")
			fmt.Println(file)
			// fmt.Printf("Test: %s\n----\nSTART\n----\n>>%s<<----\nEND\n----\n", test.name, file)
			b := []byte(file)
			cmds, err := Parse(b)
			assert.NilError(t, err)

			assert.DeepEqual(t, cmds, test.expected)
		})
	}
}

func TestParseFile(t *testing.T) {
	t.Skip()
	tests := []struct {
		filepath    string
		numCommands int
	}{
		{
			"./testdata/min-script.ns",
			2,
		},
		{
			"./testdata/mid-script.ns",
			2,
		},
		{
			"./testdata/httpbin.ns",
			8,
		},
	}
	for _, test := range tests {
		t.Run(test.filepath, func(t *testing.T) {
			cmds, err := ParseFile(test.filepath)
			assert.NilError(t, err)
			for _, cmd := range cmds {
				fmt.Println("vvv")
				fmt.Println(cmd.Comment)
				fmt.Println(cmd.Method, cmd.URL)
				fmt.Println(string(cmd.Body))
				fmt.Println("^^^")
			}
			assert.Equal(t, len(cmds), test.numCommands)
		})
	}
}
