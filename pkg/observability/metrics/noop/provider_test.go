/*
Go-lang-Vcs-NEXTClan
*/

package noop

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMetrics(t *testing.T) {
	m := GetMetrics()
	require.NotNil(t, m)

	t.Run("VCS Activity", func(t *testing.T) {
		require.NotPanics(t, func() { m.SignTime(time.Second) })
		require.NotPanics(t, func() { m.CheckAuthorizationResponseTime(time.Second) })
		require.NotPanics(t, func() { m.VerifyOIDCVerifiablePresentationTime(time.Second) })
	})
}
