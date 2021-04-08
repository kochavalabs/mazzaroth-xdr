package xdr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IDMarshal(t *testing.T) {
	var id ID
	rawbytes := bytes.Repeat([]byte{'a'}, 32)
	copy(id[:], rawbytes)
	b, err := json.Marshal(id)
	require.NoError(t, err)
	expected := []byte{'"'}
	expected = append(expected, bytes.Repeat([]byte{'6', '1'}, 32)...)
	expected = append(expected, '"')
	require.Equal(t, expected, b, `$ echo -n "%v" | xxd -p -c 256 # should look something like %v`, string(rawbytes), string(b))
}

func Test_SignatureMarshal(t *testing.T) {
	var sig Signature
	rawbytes := bytes.Repeat([]byte{'a'}, 64)
	copy(sig[:], rawbytes)
	b, err := json.Marshal(sig)
	require.NoError(t, err)
	expected := []byte{'"'}
	expected = append(expected, bytes.Repeat([]byte{'6', '1'}, 64)...)
	expected = append(expected, '"')
	require.Equal(t, expected, b, `$ echo -n "%v" | xxd -p -c 256 # should look like %v`, string(rawbytes), string(b))
}

func Test_HashMarshal(t *testing.T) {
	var hash Hash
	rawbytes := bytes.Repeat([]byte{'a'}, 32)
	copy(hash[:], rawbytes)
	b, err := json.Marshal(hash)
	require.NoError(t, err)
	expected := []byte{'"'}
	expected = append(expected, bytes.Repeat([]byte{'6', '1'}, 32)...)
	expected = append(expected, '"')
	require.Equal(t, expected, b, `$ echo -n "%v" | xxd -p -c 256 # should look something like %v`, string(rawbytes), string(b))
}

func Test_Hash32Marshal(t *testing.T) {
	var hash Hash32
	rawbytes := bytes.Repeat([]byte{'a'}, 32)
	copy(hash[:], rawbytes)
	b, err := json.Marshal(hash)
	require.NoError(t, err)
	expected := []byte{'"'}
	expected = append(expected, bytes.Repeat([]byte{'6', '1'}, 32)...)
	expected = append(expected, '"')
	require.Equal(t, expected, b, `$ echo -n "%v" | xxd -p -c 256 # should look something like %v`, string(rawbytes), string(b))
}

func Test_Hash64Marshal(t *testing.T) {
	var hash Hash64
	rawbytes := bytes.Repeat([]byte{'a'}, 64)
	copy(hash[:], rawbytes)
	b, err := json.Marshal(hash)
	require.NoError(t, err)
	expected := []byte{'"'}
	expected = append(expected, bytes.Repeat([]byte{'6', '1'}, 64)...)
	expected = append(expected, '"')
	require.Equal(t, expected, b, `$ echo -n "%v" | xxd -p -c 256 # should look something like %v`, string(rawbytes), string(b))
}

func Test_IDUnmarshal(t *testing.T) {
	st := fmt.Sprintf(`"%s"`, strings.Repeat("61", 32))
	var id ID
	err := json.Unmarshal([]byte(st), &id)
	require.NoError(t, err)
	expected := bytes.Repeat([]byte{'a'}, 32)
	require.Equal(t, expected, id[:], `$ echo "%v" | xxd -r  -p -c 256 #should look like %v`, st, string(expected))
}

func Test_SignatureUnmarshal(t *testing.T) {
	st := fmt.Sprintf(`"%s"`, strings.Repeat("61", 64))
	var sig Signature
	err := json.Unmarshal([]byte(st), &sig)
	require.NoError(t, err)
	expected := bytes.Repeat([]byte{'a'}, 64)
	require.Equal(t, expected, sig[:], `$ echo "%v" | xxd -r  -p -c 256 #should look like %v`, st, string(expected))
}

func Test_HashUnmarshal(t *testing.T) {
	st := fmt.Sprintf(`"%s"`, strings.Repeat("61", 32))
	var hash Hash
	err := json.Unmarshal([]byte(st), &hash)
	require.NoError(t, err)
	expected := bytes.Repeat([]byte{'a'}, 32)
	require.Equal(t, expected, hash[:], `$ echo "%v" | xxd -r  -p -c 256 #should look like %v`, st, string(expected))
}

func Test_Hash32Unmarshal(t *testing.T) {
	st := fmt.Sprintf(`"%s"`, strings.Repeat("61", 32))
	var hash Hash32
	err := json.Unmarshal([]byte(st), &hash)
	require.NoError(t, err)
	expected := bytes.Repeat([]byte{'a'}, 32)
	require.Equal(t, expected, hash[:], `$ echo "%v" | xxd -r  -p -c 256 #should look like %v`, st, string(expected))
}

func Test_Hash64Unmarshal(t *testing.T) {
	st := fmt.Sprintf(`"%s"`, strings.Repeat("61", 64))
	var hash Hash64
	err := json.Unmarshal([]byte(st), &hash)
	require.NoError(t, err)
	expected := bytes.Repeat([]byte{'a'}, 64)
	require.Equal(t, expected, hash[:], `$ echo "%v" | xxd -r  -p -c 256 #should look like %v`, st, string(expected))
}
