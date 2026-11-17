package headers

import (
	"bytes"
	"fmt"
	"strings"
)

const crlf = "\r\n"

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return 0, false, nil
	}
	if idx == 0 {
		// the empty line
		// headers are done, consume the CRLF
		return 2, true, nil
	}

	parts := bytes.SplitN(data[:idx], []byte(":"), 2)
	if len(parts) != 2 {
		return 0, false, fmt.Errorf("invalid header line: %s", data[:idx])
	}
	key := string(parts[0])
	value := string(parts[1])

	if strings.ContainsAny(key, " \t") {
		return 0, false, fmt.Errorf("invalid header name: %s", key)
	}

	value = strings.TrimSpace(value)

	h.Set(key, value)
	return idx + 2, false, nil
}

func (h Headers) Set(key, value string) {
	h[key] = value
}
