/*
Go-lang-Vcs-NEXTClan
*/

package requestobject

import (
	"errors"

	"github.com/trustbloc/vcs/pkg/event/spi"
)

type RequestObject struct {
	ID                       string     `json:"id"`
	Content                  string     `json:"content"`
	AccessRequestObjectEvent *spi.Event `json:"accessRequestObjectEvent"`
}

var ErrDataNotFound = errors.New("data not found")
