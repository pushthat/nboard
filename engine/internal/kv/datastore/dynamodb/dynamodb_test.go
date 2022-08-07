package dynamodb

import (
	"github.com/pushthat/nboard-engine/internal/kv"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPI_Set(t *testing.T) {
	api := New("nboard-test", "eu-north-1")

	primary := kv.PK("namespace-a")
	sort := kv.SK("user-a")
	value := []byte("hello world")

	err := api.Set(primary, sort, value)
	require.Nil(t, err)

	resp, err := api.Get(primary, sort)
	require.Nil(t, err)
	require.Equal(t, value, resp)

	err = api.Delete(primary, sort)
	require.Nil(t, err)

	resp, err = api.Get(primary, sort)
	require.Nil(t, resp)
	require.Error(t, err)

}
