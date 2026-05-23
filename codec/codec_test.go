package codec_test

import (
	"bytes"
	"compress/gzip"
	"io"
	"testing"

	"github.com/WindowsSov8forUs/sonolus-core-go/codec"
)

func TestCompressDecompressRoundTrip(t *testing.T) {
	type payload struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}

	compressed, err := codec.Compress(payload{Name: "sonolus", Count: 2})
	if err != nil {
		t.Fatal(err)
	}

	decompressed, err := codec.Decompress[payload](compressed)
	if err != nil {
		t.Fatal(err)
	}

	if decompressed.Name != "sonolus" || decompressed.Count != 2 {
		t.Fatalf("Decompress() = %#v", decompressed)
	}
}

func TestCompressUsesJSONStringWithoutTrailingNewline(t *testing.T) {
	compressed, err := codec.Compress(map[string]string{"name": "sonolus"})
	if err != nil {
		t.Fatal(err)
	}

	reader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()

	payload, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(payload), `{"name":"sonolus"}`; got != want {
		t.Fatalf("compressed payload = %q, want %q", got, want)
	}
}
