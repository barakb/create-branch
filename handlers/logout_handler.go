package handlers

import (
	"fmt"
	"github.com/barakb/github-branch/session"
	"net/http"
)

type LogoutHandler struct {
}

func (h LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LogoutHandler")
	sess := session.GlobalSessions.SessionStart(w, r)
	sess.Delete("*github.client")
	sess.Delete("user")
	session.GlobalSessions.SessionDestroy(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
