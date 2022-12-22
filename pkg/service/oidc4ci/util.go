/*
Go-lang-Vcs-NEXTClan
*/

package oidc4ci

import "time"

func WithDocumentTTL(ttl time.Duration) func(insertOptions *InsertOptions) {
	return func(insertOptions *InsertOptions) {
		insertOptions.TTL = ttl
	}
}
