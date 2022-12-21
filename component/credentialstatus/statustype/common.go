/*
* Go-lang-Vcs-NEXTClan 
*/

package statustype

type credentialSubject struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	StatusPurpose string `json:"statusPurpose,omitempty"`
	EncodedList   string `json:"encodedList"`
}
