package handlers

import (
	"fmt"
	"github.com/barakb/github-branch/session"
	gh "github.com/barakb/github-branch/github"
	"github.com/google/go-github/github"
	"net/http"
	"strings"
)

type CreateBranchHandler struct {
}

func (h CreateBranchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	fmt.Printf("uri is: %s\n", r.RequestURI)
	branchName := strings.Trim(strings.Replace(r.RequestURI, "/api/create_branch/", "", 1), "/")
	if strings.Contains(branchName, "/") {
		fmt.Printf("CreateBranchHandler: wrong request: %q\n", r.RequestURI)
		http.Error(w, fmt.Sprintf("CreateBranchHandler: wrong request: %q\n", r.RequestURI), http.StatusInternalServerError)
		return

	}

	client := sess.Get("*github.client").(*github.Client)
	_, done := gh.CreateBranch(branchName, client)
	select{
		case <- done:
	}

	w.WriteHeader(http.StatusOK)
	//"/api/create_branch/dfsdafdsa"

}
