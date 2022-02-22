package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/v42/github"
	"github.com/olekukonko/tablewriter"
)

var (
	githubUserName string
	projectsPath   string
)

var reposTitle = "## The CNCF repos\n"

func init() {
	flag.StringVar(&githubUserName, "username", "", "github user name")
	flag.StringVar(&projectsPath, "projects-path", "/projects.json", "if with stared repos")
}

type Repos struct {
	// things we want to stop referencing
	Spec ReposSpec `json:"spec"`
	// status of our unwanted dependencies
	Status ReposStatus `json:"status"`
}

type ReposSpec struct {
	// repos with descriptions or tags
	Repos map[string]string `json:"repos"`
}
type ReposStatus struct {
	// TODO add all the status of github repos to the status with date
	Repos map[string]string `json:"repos"`
}

func readFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	// Convert []byte to string and print to screen
	return string(content), err
}

func fetchAllRepos(projectsPath string, client *github.Client) []*github.Repository {
	// load repos from json
	repos, err := readFile(projectsPath)
	if err != nil {
		log.Fatalf("Error reading repos file %s: %s", repos, err)
	}
	configFromFile := &Repos{}
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(repos)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(configFromFile); err != nil {
		log.Fatalf("Error reading repos file %s: %s", projectsPath, err)
	}

	var allRepos []*github.Repository
	reposList := configFromFile.Spec.Repos
	for repo, _ := range reposList {
		if len(strings.Split(repo, "/")) < 2 {
			splits := strings.Split(repo, "/")
			// org/reponame is splits
			repository, _, err := client.Repositories.Get(context.Background(), splits[0], splits[1])
			if err != nil {
				log.Printf("WARNING: skip repo: %s for error: %s", repo, err)
				continue
			}
			allRepos = append(allRepos, repository)
		}

	}
	return allRepos
}

func makeMdTable(data [][]string, header []string) string {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	return tableString.String()
}

type RepoStatus struct {
	Name             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Lauguage         string
	ForksCount       int
	OpenIssuesCount  int
	StargazersCount  int
	SubscribersCount int
	HTMLURL          string
	//TODO lastRelease
	//TODO License
}

func makeReposString(repos []*github.Repository) string {
	reposData := [][]string{}
	for i, repo := range repos {
		reposData = append(reposData, []string{strconv.Itoa(i + 1), *repo.Name, string(*repo.StargazersCount), (*repo.UpdatedAt).String()[:10], (*repo.CreatedAt).String()[:10], string(*repo.ForksCount)})
	}
	// reposData = append(reposData, []string{"sum", "", "", "", "", strconv.Itoa(total)})
	reposString := makeMdTable(reposData, []string{"ID", "Repo", "Stars", "UpdatedAt", "CreatedAt", "ForksCount"})
	return reposTitle + reposString + "\n"
}

func main() {
	flag.Parse()
	client := github.NewClient(nil)
	repos := fetchAllRepos(projectsPath, client)
	// change sort logic here
	sort.Slice(repos[:], func(i, j int) bool {
		return *repos[j].StargazersCount < *repos[i].StargazersCount
	})

	newContentString := makeReposString(repos)

	readMeFile := path.Join(os.Getenv("GITHUB_WORKSPACE"), "README.md")
	readMeContent, err := ioutil.ReadFile(readMeFile)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`(?s)(<!--START_SECTION:my_github-->)(.*)(<!--END_SECTION:my_github-->)`)

	newContent := []byte(re.ReplaceAllString(string(readMeContent), `$1`+"\n"+newContentString+`$3`))
	err = ioutil.WriteFile(readMeFile, newContent, 0644)
	if err != nil {
		panic(err)
	}
}
