package main

import (

    "encoding/json"

    "net/http"

    "strings"

    "fmt"

    //"bufio"

    //"io/ioutil"

    "os"

)

func main() {
    /*
    resp, err := http.Get("http://www.google.com")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    text, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(text))
    */

    commitCode("https://github.com/CatherineLi0531/Commit-String-Test.git")
}

func commitCode(githubUrl string) (*string) {
    owner := ""
    repo := ""

    link := githubUrl
    link = strings.Replace(link, "https://github.com/", "", 1)
    link = strings.Replace(link, ".git", "", 1)
    split := strings.Split(link, "/")
    owner = split[0]
    repo = split[1]
    fmt.Println("Owner: " + owner);
    fmt.Println("Repo: " + repo);

    //fmt.Println(http.Get("/repos/:" + owner + "/:" + repo + "/comits"))

    resp, err := http.Get("https://api.github.com/repos/" + owner + "/" + repo + "/commits?per_page=1")
    if err != nil {
		fmt.Println(err)
        //os.Exit(1)
        return nil
	}
    /*text, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
        return nil
    }
    fmt.Println(string(text))
    */
    
    type githubCommit struct {
	    Sha string `json:"sha"`
	}
	var commits []githubCommit
	err2 := json.NewDecoder(resp.Body).Decode(&commits)
    if err2 != nil {
        fmt.Println(err2)
        os.Exit(1)
    }
	commit := commits[0].Sha
    
    fmt.Println(commit);

    return nil
}

