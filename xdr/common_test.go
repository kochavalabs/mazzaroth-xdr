package xdr

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IDMarshal(t *testing.T) {
	var id ID
	rawbytes := bytes.Repeat([]byte{'a'}, 32)
	copy(id[:], rawbytes)
	b, err := id.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, []byte(`"YWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWE="`), b, `$ echo "%v" | base64 -w0 # should look like %v`, string(rawbytes), string(b))
}

func Test_SignatureMarshal(t *testing.T) {
	var sig Signature
	rawbytes := bytes.Repeat([]byte{'a'}, 64)
	copy(sig[:], rawbytes)
	b, err := sig.MarshalJSON()
	require.NoError(t, err)
	// watch out for padding, here we have two chars of padding
	require.Equal(t, []byte(`"YWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYQ=="`), b, `$ echo "%v" | base64 -w0 # should look like %v`, string(rawbytes), string(b))
}

func Test_IDUnmarshal(t *testing.T) {
	st := `"YWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWE="`
	var id ID
	err := id.UnmarshalJSON([]byte(st))
	require.NoError(t, err)
	expected := bytes.Repeat([]byte{'a'}, 32)
	require.Equal(t, expected, id[:], `$ echo "%v" | base64 -d #should look like %v`, st, string(expected))
}
