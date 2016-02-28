package handlers

import (
	"fmt"
	"github.com/barakb/github-branch/session"
	"net/http"
	"github.com/google/go-github/github"
	"strings"
)

type CreateBranchHandler struct {
}

func (h CreateBranchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	fmt.Printf("uri is: %s\n", r.RequestURI)
	suffix := strings.Trim(strings.Replace(r.RequestURI, "/api/create_branch/", "", 1), "/")
	if strings.Contains(suffix, "/") {
		fmt.Printf("CreateBranchHandler: wrong request: %q\n", r.RequestURI)
		http.Error(w, fmt.Sprintf("CreateBranchHandler: wrong request: %q\n", r.RequestURI), http.StatusInternalServerError)
		return;

	}

	client := sess.Get("*github.client").(*github.Client)
	user, _, err := client.Users.Get("")
	if err != nil {
		panic(fmt.Errorf("error: %v\n", err))
	}
	fmt.Printf("user is %s\n", user)

	w.WriteHeader(http.StatusOK)
	//"/api/create_branch/dfsdafdsa"

}
