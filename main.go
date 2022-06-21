// TODO:
//   - add proper error handling
//	 - use a MultiWriter for the logging, i.e.,
//			mw := io.MultiWriter(os.Stdout, logFile)
//			logrus.SetOutput(mw)

package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

func getClient() *github.Client {
	tc := getOAuthClient()
	return github.NewClient(tc)
}

var ctx context.Context
var singleContext *bool

func getContext() context.Context {
	if singleContext == nil {
		ctx = context.Background()
		b := true
		singleContext = &b
	}
	return ctx
}

func getOAuthClient() *http.Client {
	apiToken, isSet := os.LookupEnv("GITHUB_TOKEN")
	if apiToken == "" || !isSet {
		panic("[ERROR] Must set $GITHUB_TOKEN!")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiToken},
	)
	return oauth2.NewClient(getContext(), ts)
}

func getConfigs(filename string) ([]Organization, error) {
	content, err := getFileContents(filename)
	if err != nil {
		panic(err)
	}

	var organizations []Organization
	extension := filepath.Ext(filename)
	if extension == ".json" {
		err = json.Unmarshal(content, &organizations)
	} else if extension == ".yaml" {
		err = yaml.Unmarshal(content, &organizations)
	} else {
		err = errors.New("[ERROR] File extension not recognized, must be either `json` or `yaml`.")
	}

	return organizations, err
}

func getFileContents(filename string) ([]byte, error) {
	f, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	return ioutil.ReadFile(f)
}

func main() {
	filename := flag.String("file", "", "Path to GitHub config file (json or yaml).")
	destroy := flag.Bool("destroy", false, "Should destroy all projects listed in the given GitHub config file.")
	flag.Parse()

	if *filename != "" {
		configs, err := getConfigs(*filename)
		if err != nil {
			panic(err)
		}
		p := NewProvisioner(configs)
		p.ProcessConfigs(*destroy)
	} else {
		fmt.Println("Please pass a config.")
	}
}
