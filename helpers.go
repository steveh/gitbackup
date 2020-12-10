package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v32/github"
	gitlab "github.com/xanzy/go-gitlab"
)

func getUsername(client interface{}, service string) string {

	if client == nil {
		log.Fatalf("Couldn't acquire a client to talk to %s", service)
	}

	if service == "github" {
		ctx := context.Background()
		user, _, err := client.(*github.Client).Users.Get(ctx, "")
		if err != nil {
			log.Fatal("Error retrieving username", err.Error())
		}
		return *user.Login
	}

	if service == "gitlab" {
		user, _, err := client.(*gitlab.Client).Users.CurrentUser()
		if err != nil {
			log.Fatal("Error retrieving username", err.Error())
		}
		return user.Username
	}

	if service == "bitbucket" {
		bitbucketUsername := os.Getenv("BITBUCKET_USERNAME")
		if bitbucketUsername == "" {
			log.Fatal("BITBUCKET_USERNAME environment variable not set")
		}
		return bitbucketUsername
	}

	return ""
}
