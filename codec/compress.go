package codec

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
)

// Compress gzips the JSON text of data like sonolus-core's compress: level 9,
// and like JSON.stringify, '<', '>' and '&' are written literally.
func Compress[T any](data T) ([]byte, error) {
	var encoded bytes.Buffer
	encoder := json.NewEncoder(&encoded)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	payload := bytes.TrimSuffix(encoded.Bytes(), []byte("\n"))

	var buffer bytes.Buffer

	writer, err := gzip.NewWriterLevel(&buffer, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	if _, err := writer.Write(payload); err != nil {
		_ = writer.Close()
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
