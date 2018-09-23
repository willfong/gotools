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

type GithubKeys struct {
	Id  int
	Key string
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Need to specify username and path to file. \n\nUsage: sshkeys <github_username> <file_to_write_to>\nExample: ./sshkeys willfong .ssh/authorized_keys\n\n")
	}

	username := os.Args[1]
	filepath := os.Args[2]
	github_keys := "https://api.github.com/users/" + username + "/keys"

	fmt.Println("Getting keys from GitHub: " + github_keys)
	fmt.Println("Writing to: " + filepath)

	resp, err := http.Get(github_keys)

	if err != nil {
		log.Fatalln("Error downloading keys from: " + github_keys)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	//THANKS: https://coderwall.com/p/4c2zig/decode-top-level-json-array-into-a-slice-of-structs-in-golang
	keys := make([]GithubKeys, 0)
	json.Unmarshal(body, &keys)

	var authorized_keys []string

	for _, line := range keys {
		authorized_keys = append(authorized_keys, line.Key)
	}

	authkey_file := []byte(strings.Join([]string(authorized_keys), "\n") + "\n")
	ioutil.WriteFile(filepath, authkey_file, 0600)
}
