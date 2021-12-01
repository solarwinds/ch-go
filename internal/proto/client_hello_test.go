package proto

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientHello_Encode(t *testing.T) {
	var b Buffer
	b.Encode(ClientHello{
		Name:            "ch",
		Major:           1,
		Minor:           1,
		ProtocolVersion: 41000,
		Database:        "github",
		User:            "neo",
		Password:        "",
	})
	const expHex = "000263680101a8c00206676974687562036e656f00"
	exp, _ := hex.DecodeString(expHex)
	require.Equal(t, exp, b.Buf)
}

func BenchmarkClientHello_Encode(b *testing.B) {
	buf := new(Buffer)
	h := &ClientHello{
		Name:            "ClickHouse Go Faster Client",
		Major:           1,
		Minor:           1,
		ProtocolVersion: 411337,
		Database:        "github",
		User:            "neo",
		Password:        "go faster",
	}
	h.Encode(buf)
	b.SetBytes(int64(len(buf.Buf)))
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		h.Encode(buf)
	}
}
