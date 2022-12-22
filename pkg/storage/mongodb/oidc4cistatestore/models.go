/*
* Go-lang-Vcs-NEXTClan
 */

package oidc4cistatestore

import "net/url"

type AuthorizeState struct {
	RedirectURI *url.URL            `json:"redirect_uri"`
	RespondMode string              `json:"respond_mode"`
	Header      map[string][]string `json:"header"`
	Parameters  map[string][]string `json:"parameters"`
}
