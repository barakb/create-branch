package handlers

import (
	"encoding/json"
	gh "github.com/barakb/create-branch/github"
	"github.com/barakb/create-branch/session"
	"github.com/google/go-github/github"
	"net/http"
)

type GetBranchsHandler struct {
	WS *WebSocketHandler
}

func (h GetBranchsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	client := sess.Get("*github.client").(*github.Client)
	branches:= <- gh.ListAllRefsAsMap(client)
	js, err := json.Marshal(branches)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)


}
