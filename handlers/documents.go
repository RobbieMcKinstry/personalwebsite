package handlers

import (
	"log"
	"net/http"

	minio "github.com/minio/minio-go"
)

type DigitalOceanHandler struct {
	client       *minio.Client
	AccessToken  string
	AccessSecret string
	endpoint     string
	spacename    string
}

func NewDigitalOceanHandler() *DigitalOceanHandler {
	var res = DigitalOceanHandler{
		AccessToken:  digitalOceanAccessKey,
		AccessSecret: digitalOceanAccessSecret,
		endpoint:     digitaloceanEndpoint,
		spacename:    spaceName,
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

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		renderTemplate(w, "failed_login.tmpl", nil)
		return
	}

	var username = r.FormValue("username")
	var password = r.FormValue("password")

	if username != expectedUsername || password != expectedPassword {
		w.WriteHeader(http.StatusForbidden)
		renderTemplate(w, "failed_login.tmpl", nil)
		return
	}

	// Otherwise, we can fetch the list of documents from DigitalOcean
	doneCh := make(chan struct{})
	defer close(doneCh)
	// Recursively list all objects in 'mytestbucket'
	recursive := true
	var documents = []minio.ObjectInfo{}
	for message := range handler.client.ListObjectsV2(handler.spacename, "", recursive, doneCh) {
		documents = append(documents, message)
	}
	renderTemplate(w, "documents.tmpl", documents)
}

var expectedUsername = "robbiemckinstry"
var expectedPassword = "[F*mVof78A7qheNjhwFHcovsmQB&XEZKwNMgz2T6MGCdB4nz=uMhDoAsQmdK[+ai"
var digitalOceanAccessKey = "YH37572NVSDWDCG4IELC"
var digitalOceanAccessSecret = "ttCwqhbOVxAek7avRO4JRBPWweSGrGJ7amw31V8uHhY"
var digitaloceanEndpoint = "nyc3.digitaloceanspaces.com"
var spaceName = "latex-documents"
