package github

import (
	"fmt"
	"github.com/google/go-github/github"
	"strings"
	"sync"
)

var ReposNames = []string{"CloudifySource/Cloudify-iTests-webuitf",
	"CloudifySource/Cloudify-iTests",
	"GigaSpaces-QA/devops",
	"Gigaspaces/xap-ui-desktop",
	"Gigaspaces/RESTData",
	"Gigaspaces/xap-session-sharing-manager",
	"CloudifySource/iTests-Framework",
	"GigaSpaces/mongo-datasource-itests",
	"GigaSpaces/mongo-datasource",
	"GigaSpaces/xap-mule",
	"Gigaspaces/xap-scala",
	"GigaSpaces/petclinic-jpa",
	"GigaSpaces/xap-apm-introscope",
	"GigaSpaces/xap-blobstore-mapdb",
	"GigaSpaces/xap-blobstore-rocksdb",
	"GigaSpaces/xap-cassandra",
	"GigaSpaces/xap-cpp",
	"GigaSpaces/xap-dotnet",
	"GigaSpaces/xap-example-data",
	"GigaSpaces/xap-example-helloworld",
	"GigaSpaces/xap-example-tutorials",
	"GigaSpaces/xap-example-web",
	"GigaSpaces/xap-jetty",
	"GigaSpaces/xap-jms",
	"GigaSpaces/xap-maven-plugin",
	"GigaSpaces/xap-rest",
	"GigaSpaces/xap-session-sharing-manager-itests",
	"GigaSpaces/xap-spatial",
	"GigaSpaces/xap-spring-data",
	"GigaSpaces/xap-ui-web",
	"GigaSpaces/xap"}

type Repo struct {
	owner string
	repo  string
	sha   string
}

func (repo Repo) String() string {
	return fmt.Sprintf("Repo{owner:%s, repo:%s, sha:%s}", repo.owner, repo.repo, repo.sha)
}

func createReposFromNames(names []string) []*Repo {
	ret := make([]*Repo, 0)
	for _, name := range names {
		if strings.TrimSpace(name) != "" && strings.Contains(name, "/") {
			components := strings.Split(name, "/")
			ret = append(ret, &Repo{owner: components[0], repo: components[1]})
		}
	}
	return ret
}

func fillTipSha(repos []*Repo, client *github.Client) *sync.WaitGroup {
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *Repo) {
			defer wg.Done()
			//fmt.Printf("working on repo %#v\n", repo)
			ref, _, err := client.Git.GetRef(repo.owner, repo.repo, "refs/heads/master")
			if err != nil {
				fmt.Printf("error: %#v\n", err)
			}
			//fmt.Printf("Ref: %s\nURL: %s\n", *ref.Ref, *ref.URL)
			//fmt.Printf("Type: %s\nURL: %s\nSHA: %s\n", *ref.Object.Type, *ref.Object.URL, *ref.Object.SHA)
			repo.sha = *ref.Object.SHA
		}(repo)
	}
	return &wg
}


func createBranchesWithProgress(repos []*Repo, branch string, client *github.Client) (progressChan chan RepoStatus, resChan chan map[string]interface{}) {
	progressChan = make(chan RepoStatus, len(repos))
	resChan = make(chan map[string]interface{})
	intermediateProgressChan := make(chan RepoStatus, len(repos))
	for _, repo := range repos {
		go func(repo *Repo) {
			_, _, err := client.Git.CreateRef(repo.owner, repo.repo, &github.Reference{
				Ref: github.String("refs/heads/" + branch),
				Object: &github.GitObject{
					SHA: &repo.sha,
				},
			})
			if err != nil {
				fmt.Printf("error createing branch %s in repo: %#v, error is %s:\n", branch, repo, err.Error())
				intermediateProgressChan <- RepoStatus{Name: repo.owner + "/" +repo.repo, Success: false}
			} else {
				fmt.Printf("branch %s created in repo %#v\n", branch, repo)
				intermediateProgressChan <- RepoStatus{Name: repo.owner + "/" +repo.repo, Success: true}
			}
		}(repo)
	}
	go func() {
		var statuses map[string]interface{} = make(map[string]interface{})
		for range repos {
			status := <-intermediateProgressChan
			statuses[status.Name] = status.Success
			progressChan <- status
		}
		resChan <- statuses
		close(progressChan)
		close(resChan)
		close(intermediateProgressChan)
	}()
	return progressChan, resChan
}


func deleteBranchesWithProgress(repos []*Repo, branch string, client *github.Client) (progressChan chan RepoStatus, resChan chan map[string]interface{}) {
	progressChan = make(chan RepoStatus, len(repos))
	resChan = make(chan map[string]interface{})
	intermediateProgressChan := make(chan RepoStatus, len(repos))
	for _, repo := range repos {
		go func(repo *Repo) {
			_, err := client.Git.DeleteRef(repo.owner, repo.repo, "refs/heads/" + branch)
			if err != nil {
				fmt.Printf("error deleting branch %s in repo: %#v, error is %s:\n", branch, repo, err.Error())
				intermediateProgressChan <- RepoStatus{Name: repo.repo, Success: false}
			} else {
				//fmt.Printf("branch %s deleted from repo %#v\n", branch, repo)
				intermediateProgressChan <- RepoStatus{Name: repo.repo, Success: true}
			}
		}(repo)
	}
	go func() {
		var statuses map[string]interface{} = make(map[string]interface{})
		for range repos {
			status := <-intermediateProgressChan
			statuses[status.Name] = status.Success
			progressChan <- status
		}
		resChan <- statuses
		close(progressChan)
		close(resChan)
		close(intermediateProgressChan)
	}()
	return progressChan, resChan
}

type BranchOnRepositories struct {
	Name     string       `json:"name"`
	Statuses map[string]bool `json:"statuses"`
}

func ListAllRefsAsMap(client *github.Client) chan map[string]map[string]bool {
	type branch struct {
		repo   string
		branch string
	}
	branches := make(chan branch, 100)
	repos := createReposFromNames(ReposNames)
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *Repo) {
			defer func() {
				wg.Done()
			}()
			heads, err := ListRefs(client, repo.owner, repo.repo)
			if err == nil {
				for _, head := range heads {
					branchName := strings.TrimPrefix(*head.Ref, "refs/heads/")
					branches <- branch{repo.owner + "/" + repo.repo, branchName}
				}
			} else {
				fmt.Printf("List all ref error: %#v, on repo %#v\n", err, repo)
			}
		}(repo)
	}

	// closer
	go func() {
		wg.Wait()
		close(branches)
	}()

	var res map[string]map[string]bool = make(map[string]map[string]bool)
	resChan := make(chan map[string]map[string]bool)
	go func() {
		for branch := range branches {
			if b, ok := res[branch.branch]; ok {
				b[branch.repo] = true
			}else {
				res[branch.branch] = map[string]bool{branch.repo : true}
			}
		}
		resChan <- res
		close(resChan)
	}()
	return resChan
}

func DeleteBranchWithProgress(branch string, client *github.Client) (progressChan chan RepoStatus, resChan chan map[string]interface{}) {
	repos := createReposFromNames(ReposNames)
	return deleteBranchesWithProgress(repos, branch, client)
}

func CreateBranchsWithProgress(branch string, client *github.Client) (progressChan chan RepoStatus, resChan chan map[string]interface{}) {
	repos := createReposFromNames(ReposNames)
	fillTipSha(repos, client).Wait()
	return createBranchesWithProgress(repos, branch, client)
}

func ListRefs(client *github.Client, owner string, repo string) ([]github.Reference, error) {
	refs, _, err := client.Git.ListRefs(owner, repo, &github.ReferenceListOptions{Type: "heads/"})
	//refs, _, err := client.Git.ListRefs(owner, repo, &github.ReferenceListOptions{Type:"heads"});
	return refs, err
}

type RefList struct {
}

type UIBranch struct {
	Name  string       `json:"name"`
	Repos []RepoStatus `json:"repos"`
}

type RepoStatus struct {
	Name    string `json:"name"`
	Success bool   `json:"success"`
}

type repoWithRefs struct {
	name string
	refs []github.Reference
}