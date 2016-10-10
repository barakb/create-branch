package handlers

import (
	"encoding/json"
	gh "github.com/barakb/create-branch/github"
	"github.com/barakb/create-branch/session"
	"github.com/google/go-github/github"
	"net/http"
)

type GetBranchsHandler struct {
	WS    *WebSocketHandler
	Repos []string
}

type Branches struct {
	Map   map[string]map[string]bool `json:"branches"`
	Repos []string `json:"repos"`
}

func (h GetBranchsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	client := sess.Get("*github.client").(*github.Client)
	branches := <-gh.ListAllRefsAsMap(client)
	js, err := json.Marshal(Branches{branches, h.Repos})
		if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)


	}
