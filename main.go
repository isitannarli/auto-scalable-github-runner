package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github-actions-basic-auto-scaling-with-self-hosted-runners/clients"
	"github-actions-basic-auto-scaling-with-self-hosted-runners/utils"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type GithubRequestBody struct {
	Action      string `json:"action"`
	WorkflowJob struct {
		Id          int64       `json:"id"`
		RunId       int         `json:"run_id"`
		RunUrl      string      `json:"run_url"`
		NodeId      string      `json:"node_id"`
		HeadSha     string      `json:"head_sha"`
		Url         string      `json:"url"`
		HtmlUrl     string      `json:"html_url"`
		Status      string      `json:"status"`
		Conclusion  interface{} `json:"conclusion"`
		StartedAt   time.Time   `json:"started_at"`
		CompletedAt interface{} `json:"completed_at"`
		Name        string      `json:"name"`
		Steps       []struct {
			Name        string      `json:"name"`
			Status      string      `json:"status"`
			Conclusion  interface{} `json:"conclusion"`
			Number      int         `json:"number"`
			StartedAt   time.Time   `json:"started_at"`
			CompletedAt interface{} `json:"completed_at"`
		} `json:"steps"`
		CheckRunUrl     string   `json:"check_run_url"`
		Labels          []string `json:"labels"`
		RunnerId        int      `json:"runner_id"`
		RunnerName      string   `json:"runner_name"`
		RunnerGroupId   int      `json:"runner_group_id"`
		RunnerGroupName string   `json:"runner_group_name"`
	} `json:"workflow_job"`
	Repository struct {
		Id       int    `json:"id"`
		NodeId   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HtmlUrl          string      `json:"html_url"`
		Description      string      `json:"description"`
		Fork             bool        `json:"fork"`
		Url              string      `json:"url"`
		ForksUrl         string      `json:"forks_url"`
		KeysUrl          string      `json:"keys_url"`
		CollaboratorsUrl string      `json:"collaborators_url"`
		TeamsUrl         string      `json:"teams_url"`
		HooksUrl         string      `json:"hooks_url"`
		IssueEventsUrl   string      `json:"issue_events_url"`
		EventsUrl        string      `json:"events_url"`
		AssigneesUrl     string      `json:"assignees_url"`
		BranchesUrl      string      `json:"branches_url"`
		TagsUrl          string      `json:"tags_url"`
		BlobsUrl         string      `json:"blobs_url"`
		GitTagsUrl       string      `json:"git_tags_url"`
		GitRefsUrl       string      `json:"git_refs_url"`
		TreesUrl         string      `json:"trees_url"`
		StatusesUrl      string      `json:"statuses_url"`
		LanguagesUrl     string      `json:"languages_url"`
		StargazersUrl    string      `json:"stargazers_url"`
		ContributorsUrl  string      `json:"contributors_url"`
		SubscribersUrl   string      `json:"subscribers_url"`
		SubscriptionUrl  string      `json:"subscription_url"`
		CommitsUrl       string      `json:"commits_url"`
		GitCommitsUrl    string      `json:"git_commits_url"`
		CommentsUrl      string      `json:"comments_url"`
		IssueCommentUrl  string      `json:"issue_comment_url"`
		ContentsUrl      string      `json:"contents_url"`
		CompareUrl       string      `json:"compare_url"`
		MergesUrl        string      `json:"merges_url"`
		ArchiveUrl       string      `json:"archive_url"`
		DownloadsUrl     string      `json:"downloads_url"`
		IssuesUrl        string      `json:"issues_url"`
		PullsUrl         string      `json:"pulls_url"`
		MilestonesUrl    string      `json:"milestones_url"`
		NotificationsUrl string      `json:"notifications_url"`
		LabelsUrl        string      `json:"labels_url"`
		ReleasesUrl      string      `json:"releases_url"`
		DeploymentsUrl   string      `json:"deployments_url"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		PushedAt         time.Time   `json:"pushed_at"`
		GitUrl           string      `json:"git_url"`
		SshUrl           string      `json:"ssh_url"`
		CloneUrl         string      `json:"clone_url"`
		SvnUrl           string      `json:"svn_url"`
		Homepage         interface{} `json:"homepage"`
		Size             int         `json:"size"`
		StargazersCount  int         `json:"stargazers_count"`
		WatchersCount    int         `json:"watchers_count"`
		Language         interface{} `json:"language"`
		HasIssues        bool        `json:"has_issues"`
		HasProjects      bool        `json:"has_projects"`
		HasDownloads     bool        `json:"has_downloads"`
		HasWiki          bool        `json:"has_wiki"`
		HasPages         bool        `json:"has_pages"`
		ForksCount       int         `json:"forks_count"`
		MirrorUrl        interface{} `json:"mirror_url"`
		Archived         bool        `json:"archived"`
		Disabled         bool        `json:"disabled"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		License          interface{} `json:"license"`
		Forks            int         `json:"forks"`
		OpenIssues       int         `json:"open_issues"`
		Watchers         int         `json:"watchers"`
		DefaultBranch    string      `json:"default_branch"`
	} `json:"repository"`
	Organization struct {
		Login            string `json:"login"`
		Id               int    `json:"id"`
		NodeId           string `json:"node_id"`
		Url              string `json:"url"`
		ReposUrl         string `json:"repos_url"`
		EventsUrl        string `json:"events_url"`
		HooksUrl         string `json:"hooks_url"`
		IssuesUrl        string `json:"issues_url"`
		MembersUrl       string `json:"members_url"`
		PublicMembersUrl string `json:"public_members_url"`
		AvatarUrl        string `json:"avatar_url"`
		Description      string `json:"description"`
	} `json:"organization"`
	Sender struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
}

type DockerEnvironments struct {
	RunnerOs           string `json:"RUNNER_OS"`
	RunnerArchitecture string `json:"RUNNER_ARCHITECTURE"`
	RunnerVersion      string `json:"RUNNER_VERSION"`
	RunnerName         string `json:"RUNNER_NAME"`
	RunnerToken        string `json:"RUNNER_TOKEN"`
	RunnerWorkdir      string `json:"RUNNER_WORKDIR"`
	GithubOwner        string `json:"GITHUB_OWNER"`
	GithubRepository   string `json:"GITHUB_REPOSITORY"`
	ContainerName      string `json:"CONTAINER_NAME"`
}

var (
	runnerWorkdir = "./runners"
	runnerVersion = "2.294.0"
)

func writeEnvFile(entity DockerEnvironments) (bool, error) {
	dir, err := os.Getwd()

	if err != nil {
		return false, err
	}

	envFile, err := os.Create(filepath.Join(dir, runnerWorkdir, utils.Slugify(entity.GithubRepository), "./.env"))

	if err != nil {
		return false, err
	}

	defer envFile.Close()

	out, err := json.Marshal(entity)

	if err != nil {
		return false, err
	}

	var d interface{}

	if err := json.Unmarshal(out, &d); err != nil {
		return false, err
	}

	environments := d.(map[string]interface{})

	writer := bufio.NewWriter(envFile)

	for key, value := range environments {
		result := fmt.Sprintf("%s=%s", key, value)

		if _, err := writer.WriteString(fmt.Sprintf("%s\n", result)); err != nil {
			log.Fatal("Failed to write to file")
		}
	}

	writer.Flush()

	return true, nil
}

func hookHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	var payload GithubRequestBody

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	runnerToken, err := clients.GithubActionRunnerTokenGenerator(payload.Repository.Owner.Login, payload.Repository.Name)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	environments := &DockerEnvironments{
		RunnerOs:           "linux",
		RunnerArchitecture: "x64",
		RunnerVersion:      runnerVersion,
		RunnerName:         "self-hosted",
		RunnerToken:        *runnerToken,
		RunnerWorkdir:      runnerWorkdir,
		ContainerName:      "self-hosted",
		GithubOwner:        payload.Repository.Owner.Login,
		GithubRepository:   payload.Repository.FullName,
	}

	if err := utils.EnsureDir(runnerWorkdir); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := utils.CopyDir("./templates", filepath.Join(runnerWorkdir, utils.Slugify(environments.GithubRepository))); err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := writeEnvFile(*environments); err != nil {
		fmt.Println("File write failed with error: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusNoContent)
	return

	out, err := exec.Command("docker", "compose", "--env-file", fmt.Sprintf("./%s", filepath.Join(runnerWorkdir, "/.env")), "up", "--build", "-d").Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}

func main() {
	api := http.NewServeMux()
	api.HandleFunc("/api/hooks", hookHandler)

	err := http.ListenAndServe(":8080", api)

	log.Fatal(err)
}
