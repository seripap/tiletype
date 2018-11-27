// Package tiletype detects common map tile formats from a buffer
package tiletype

type bufferFunc func([]byte) bool

// TileHeader provides string values to content encoding and types
// that was translated from a buffer.
type TileHeader struct {
	ContentType     string
	ContentEncoding string
}

// Jpeg determines if buffer data is file type jpg.
func Jpeg(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0xFF &&
		buf[1] == 0xD8 &&
		buf[2] == 0xFF
}

// Png determines if buffer data is file type png.
func Png(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x89 && buf[1] == 0x50 &&
		buf[2] == 0x4E && buf[3] == 0x47
}

// Gif determines if buffer data is file type gif.
func Gif(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x47 && buf[1] == 0x49 && buf[2] == 0x46
}

// Webp determines if buffer data is file type webp.
func Webp(buf []byte) bool {
	return len(buf) > 10 &&
		buf[0] == 0x52 && buf[1] == 0x49 &&
		buf[2] == 0x46 && buf[3] == 0x46 &&
		buf[8] == 0x57 && buf[9] == 0x45 &&
		buf[10] == 0x42 && buf[11] == 0x50
}

// Pbf determines if buffer data is file type pbf
func Pbf(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x78 && buf[1] == 0x9C ||
		len(buf) > 2 &&
			buf[0] == 0x1F && buf[1] == 0x8B
}

// Type translates buffer to file type by buffer headers.
func Type(buf []byte) string {
	scenarios := map[string]bufferFunc{
		"jpg":  Jpeg,
		"png":  Png,
		"gif":  Gif,
		"webp": Webp,
		"pbf":  Pbf,
	}

	for key, scenario := range scenarios {
		if scenario(buf) {
			return key
		}
	}

	return "unknown"
}

// Header returns ContentType and ContentEncoding based on file
// type from a buffer.
func Headers(buf []byte) *TileHeader {
	bufferType := Type(buf)

	switch bufferType {
	case "jpg":
		return &TileHeader{ContentType: "image/jpeg"}
	case "png":
		return &TileHeader{ContentType: "image/png"}
	case "gif":
		return &TileHeader{ContentType: "image/gif"}
	case "webp":
		return &TileHeader{ContentType: "image/webp"}
	case "pbf":
		if buf[0] == 0x78 && buf[1] == 0x9C {
			return &TileHeader{ContentType: "application/x-protobuf", ContentEncoding: "defalte"}
		} else if buf[0] == 0x1F && buf[1] == 0x8B {
			return &TileHeader{ContentType: "application/x-protobuf", ContentEncoding: "gzip"}
		}
	}

	return &TileHeader{}
}
