package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type RunnerTokenGenerator struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

func GithubActionRunnerTokenGenerator(githubOwner string, githubRepository string) (*string, error) {
	client := http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	var url = fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/runners/registration-token", githubOwner, githubRepository)

	request, _ := http.NewRequest("POST", url, nil)

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("GITHUB_PERSONAL_TOKEN")))

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var result RunnerTokenGenerator
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result.Token, nil
}
