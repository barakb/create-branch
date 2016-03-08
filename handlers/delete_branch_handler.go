package handlers

import (
	"fmt"
	"github.com/barakb/github-branch/session"
	gh "github.com/barakb/github-branch/github"
	"github.com/google/go-github/github"
	"net/http"
	"strings"
	"encoding/json"
)

type DeleteBranchHandler struct {
}

func (h DeleteBranchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	branchName := strings.Trim(strings.Replace(r.RequestURI, "/api/delete_branch/", "", 1), "/")
	if strings.Contains(branchName, "/") || r.Method != "DELETE" {
		fmt.Printf("DeleteBranchHandler: wrong request: %s %q\n",r.Method, r.RequestURI)
		http.Error(w, fmt.Sprintf("DeleteBranchHandler: wrong request: %s %q\n", r.Method, r.RequestURI), http.StatusInternalServerError)
		return
	}

	client := sess.Get("*github.client").(*github.Client)
	gh.DeleteBranch(branchName, client)
	deleted := gh.UIBranch{branchName, 0}
	js, err := json.Marshal(deleted)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
