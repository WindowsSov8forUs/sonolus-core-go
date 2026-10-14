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

func TestCompressWritesHTMLCharactersLiterally(t *testing.T) {
	compressed, err := codec.Compress(map[string]string{"title": "<A & B>"})
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
	if got, want := string(payload), `{"title":"<A & B>"}`; got != want {
		t.Fatalf("compressed payload = %q, want %q (JSON.stringify output)", got, want)
	}
}

func TestDecompressRejectsTrailingData(t *testing.T) {
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	_, _ = writer.Write([]byte(`{"name":"sonolus"} trailing`))
	_ = writer.Close()
	if _, err := codec.Decompress[map[string]string](buffer.Bytes()); err == nil {
		t.Fatal("Decompress accepted trailing data after the JSON value")
	}
}

func TestDecompressVerifiesChecksum(t *testing.T) {
	compressed, err := codec.Compress(map[string]string{"name": "sonolus"})
	if err != nil {
		t.Fatal(err)
	}
	corrupted := append([]byte(nil), compressed...)
	corrupted[len(corrupted)-8] ^= 0xff // CRC32 in the gzip trailer
	if _, err := codec.Decompress[map[string]string](corrupted); err == nil {
		t.Fatal("Decompress accepted a corrupted checksum")
	}
}
