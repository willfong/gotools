package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Format of the JSON response from GitHub
type GithubKeys struct {
	ID  int
	Key string
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Need to specify username and path to file. \n\nUsage: sshkeys <github_username> <file_to_write_to>\nExample: ./sshkeys willfong .ssh/authorized_keys")
	}

	username := os.Args[1]
	filepath := os.Args[2]
	githubKeys := "https://api.github.com/users/" + username + "/keys"

	fmt.Println("Getting keys from GitHub: " + githubKeys)
	fmt.Println("Writing to: " + filepath)

	resp, err := http.Get(githubKeys)

	if err != nil {
		log.Fatalln("Error downloading keys from: " + githubKeys)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	//THANKS: https://coderwall.com/p/4c2zig/decode-top-level-json-array-into-a-slice-of-structs-in-golang
	keys := make([]GithubKeys, 0)
	json.Unmarshal(body, &keys)

	var authorizedKeys []string

	for _, line := range keys {
		authorizedKeys = append(authorizedKeys, line.Key)
	}

	authkeyFile := []byte(strings.Join([]string(authorizedKeys), "\n") + "\n")
	ioutil.WriteFile(filepath, authkeyFile, 0600)
}
