package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RobbieMcKinstry/personal-website/database"

	minio "github.com/minio/minio-go"
)

type DigitalOceanHandler struct {
	client       *minio.Client
	AccessToken  string
	AccessSecret string
	endpoint     string
	spacename    string
	db           *database.Handle
}

type documentPageVariables struct {
	Documents []minio.ObjectInfo
	Rosters   []string
}

func NewDigitalOceanHandler() *DigitalOceanHandler {
	var res = DigitalOceanHandler{
		AccessToken:  digitalOceanAccessKey,
		AccessSecret: digitalOceanAccessSecret,
		endpoint:     digitaloceanEndpoint,
		spacename:    spaceName,
		db:           database.NewDB(),
	}
	var ssl = true
	client, err := minio.New(res.endpoint, res.AccessToken, res.AccessSecret, ssl)
	if err != nil {
		log.Fatal(err)
	}

	res.client = client

	return &res
}

// The DocumentHandler points to the website which lists
// all LaTeX documents and allows updates of new documents.
func (handler *DigitalOceanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ok, err := handler.client.BucketExists(handler.spacename)
	if err != nil {
		fmt.Println("Error found when checking for bucket")
		fmt.Println(err)
		return
	}
	if !ok {
		fmt.Println("Bucket does not exist.")
		return
	}

	if r.Method != http.MethodGet {
		fmt.Println("Request not a GET")
		w.WriteHeader(http.StatusForbidden)
		renderTemplate(w, "failed_login.tmpl", nil)
		return
	}

	var username = r.FormValue("username")
	var password = r.FormValue("password") //
	if username != expectedUsername || password != expectedPassword {
		fmt.Println("Username/password not matching")
		w.WriteHeader(http.StatusForbidden)
		renderTemplate(w, "failed_login.tmpl", nil)
		return
	}

	// Otherwise, we can fetch the list of documents from DigitalOcean
	doneCh := make(chan struct{})
	var vars documentPageVariables
	defer close(doneCh)
	// Recursively list all objects in 'mytestbucket'
	recursive := true
	for doc := range handler.client.ListObjectsV2(handler.spacename, "", recursive, doneCh) {
		vars.Documents = append(vars.Documents, doc)
	}
	vars.Rosters = handler.db.Rosters()

	renderTemplate(w, "documents.tmpl", vars)
}

var expectedUsername = "robbiemckinstry"
var expectedPassword = "[F*mVof78A7qheNjhwFHcovsmQB&XEZKwNMgz2T6MGCdB4nz=uMhDoAsQmdK[+ai"
var digitalOceanAccessKey = "YH37572NVSDWDCG4IELC"
var digitalOceanAccessSecret = "ttCwqhbOVxAek7avRO4JRBPWweSGrGJ7amw31V8uHhY"
var digitaloceanEndpoint = "nyc3.digitaloceanspaces.com"
var spaceName = "latex-documents"
