package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// MakeParams returns a string with key-value params to sign the request
func MakeParams(params map[string]string) string {
	if params == nil {
		return ""
	}

	v := sortKeys(params)
	u := url.Values{}

	for i := range v {
		if params[v[i]] == "" {
			continue
		}
		u.Add(v[i], params[v[i]])
	}

	var urlParams *url.URL
	urlParams, err := url.Parse("")
	if err != nil {
		return ""
	}
	urlParams.RawQuery = u.Encode()
	return "&" + strings.Replace(urlParams.String(), "+", "%20", -1)[1:]
}

// SetEmptyIfFalse returns a zero string value if b is false
func SetEmptyIfFalse(b bool) string {
	s := string(strconv.AppendBool(make([]byte, 0), b))
	if s != "false" {
		return s
	}

	return ""
}

// MakeForm return the http form to send on request
func MakeForm(m map[string]string) *strings.Reader {
	vals := url.Values{}
	for k, v := range m {
		if v == "" {
			continue
		}
		vals.Set(k, v)
	}

	return strings.NewReader(vals.Encode())
}

// UnmarshalJSON returns an error if cannot unmarshal the json
// allocated on r
func UnmarshalJSON(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}