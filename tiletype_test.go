package tiletype

import (
	"io/ioutil"
	"testing"
)

func openFile(filename string) []byte {
	fileBytes, _ := ioutil.ReadFile(filename)

	return fileBytes
}

func TestType(t *testing.T) {
	tests := []struct {
		fileType string
		file     []byte
	}{
		{"gif", openFile("fixtures/0.gif")},
		{"jpg", openFile("fixtures/0.jpeg")},
		{"png", openFile("fixtures/0.png")},
		{"webp", openFile("fixtures/0.webp")},
		{"pbf", openFile("fixtures/0.vector.pbf")},
		{"pbf", openFile("fixtures/0.vector.pbfz")},
		{"unknown", openFile("fixtures/unknown.txt")},
		{"webp", openFile("fixtures/tux.webp")},
		{"webp", openFile("fixtures/tux_alpha.webp")},
	}

	for _, test := range tests {
		fileType := Type(test.file)
		if fileType != test.fileType {
			t.Fatal("GIVEN:", test.fileType, "Could not validate to:", fileType)
		}
	}
}

func ExampleType() {
	// openFile is psuedo code. In this example, it returns
	// bytes of `tile.png`.
	bytes := openFile("tile.png")
	fileType := Type(bytes) // png
}

func TestHeaders(t *testing.T) {
	tests := []struct {
		contentType     string
		contentEncoding string
		file            []byte
	}{
		{"image/gif", "", openFile("fixtures/0.gif")},
		{"image/jpeg", "", openFile("fixtures/0.jpeg")},
		{"image/png", "", openFile("fixtures/0.png")},
		{"image/webp", "", openFile("fixtures/0.webp")},
		{"application/x-protobuf", "defalte", openFile("fixtures/0.vector.pbf")},
		{"application/x-protobuf", "gzip", openFile("fixtures/0.vector.pbfz")},
		{"", "", openFile("fixtures/unknown.txt")},
		{"image/webp", "", openFile("fixtures/tux.webp")},
		{"image/webp", "", openFile("fixtures/tux_alpha.webp")},
	}

	for _, test := range tests {
		headers := Headers(test.file)
		if headers.ContentType != test.contentType || headers.ContentEncoding != test.contentEncoding {
			t.Fatal("GIVEN:", test.contentType, test.contentEncoding, "Could not validate to:", headers)
		}
	}
}

func ExampleHeaders() {
	// openFile is psuedo code. In this example, it returns
	// bytes of `tile.png`.
	bytes := openFile("tile.png")
	fileHeaders := Headers(bytes) // TileHeader{contentType, contentEncoding}
}
