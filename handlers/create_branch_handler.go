package handlers

import (
	"encoding/json"
	"fmt"
	gh "github.com/barakb/create-branch/github"
	"github.com/barakb/create-branch/session"
	"github.com/google/go-github/github"
	"net/http"
	"strings"
)

type CreateBranchHandler struct {
}

func (h CreateBranchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	branchName := strings.Trim(strings.Replace(r.RequestURI, "/api/create_branch/", "", 1), "/")
	if strings.Contains(branchName, "/") {
		fmt.Printf("CreateBranchHandler: wrong request: %q\n", r.RequestURI)
		http.Error(w, fmt.Sprintf("CreateBranchHandler: wrong request: %q\n", r.RequestURI), http.StatusInternalServerError)
		return

	}

	client := sess.Get("*github.client").(*github.Client)
	_, done, counter := gh.CreateBranch(branchName, client)
	select {
	case <-done:
	}
	created := gh.UIBranch{Name: branchName, Quantity: int(*counter)}
	js, err := json.Marshal(created)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
