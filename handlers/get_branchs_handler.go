package handlers

import (
	"fmt"
	"github.com/barakb/github-branch/session"
	gh "github.com/barakb/github-branch/github"
	"github.com/google/go-github/github"
	"net/http"
)

type GetBranchsHandler struct {
}

func (h GetBranchsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	fmt.Printf("uri is: %s\n", r.RequestURI)
	client := sess.Get("*github.client").(*github.Client)
	gh.ListAllRefs(client)

	w.WriteHeader(http.StatusOK)
	//"/api/create_branch/dfsdafdsa"

}
