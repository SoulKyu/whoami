// pkg/session/session.go
package session

import (
	"github.com/gorilla/sessions"
)

// Store est le magasin de sessions qui sera utilisé dans toute l'application
var Store = sessions.NewCookieStore([]byte("something-very-secret"))
