package codec

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/json"
	"io"
)

func Decompress[T any](data []byte) (T, error) {
	var value T

	reader, err := compressedReader(data)
	if err != nil {
		return value, err
	}
	defer reader.Close()

	// Read to EOF so the gzip/zlib checksum is verified, and reject trailing
	// data after the JSON value, as JSON.parse does.
	payload, err := io.ReadAll(reader)
	if err != nil {
		return value, err
	}
	if err := json.Unmarshal(payload, &value); err != nil {
		return value, err
	}
	return value, nil
}

func compressedReader(data []byte) (io.ReadCloser, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err == nil {
		return reader, nil
	}

	return zlib.NewReader(bytes.NewReader(data))
}
