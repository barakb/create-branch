package handlers

import (
	"github.com/barakb/github-branch/session"
	gh "github.com/barakb/github-branch/github"
	"github.com/google/go-github/github"
	"net/http"
	"encoding/json"
)

type GetBranchsHandler struct {
}

func (h GetBranchsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	client := sess.Get("*github.client").(*github.Client)
	branches, err := gh.ListAllRefs(client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(branches)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
