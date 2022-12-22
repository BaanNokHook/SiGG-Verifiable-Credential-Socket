/*
Go-lang-Vcs-NEXTClan
*/

package prometheus

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	h := NewHandler()
	require.NotNil(t, h)
	require.Equal(t, "/metrics", h.Path())
	require.Equal(t, http.MethodGet, h.Method())
	require.NotNil(t, h.Handler())
}
