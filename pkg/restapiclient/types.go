/*
* Go-lang-Vcs-NEXTClan
 */

package restapiclient

type PrepareClaimDataAuthorizationRequest struct {
	OpState string `json:"op_state"`
}

type PrepareClaimDataAuthorizationResponse struct {
	RedirectURI string `json:"redirect_uri"`
}

type StoreAuthorizationCodeRequest struct {
	OpState string `json:"op_state"`
	Code    string `json:"code"`
}

type StoreAuthorizationCodeResponse struct {
	Success bool `json:"success"`
}

type PushAuthorizationRequest struct {
	OpState        string `json:"op_state"`
	CredentialType string `json:"credential_type"`
	Format         string `json:"format"`
}

type PushAuthorizationResponse struct {
	TxID string `json:"tx_id"`
}
