package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
)

// The IndexHandler points to the root of the application.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.tmpl", nil)
}

func ResearchHandler(w http.ResponseWriter, r *http.Request) {
	RenderMarkdown(w, r, "Research", "research.md")
}

func EngineeringHandler(w http.ResponseWriter, r *http.Request) {
	RenderMarkdown(w, r, "Engineering", "engineering.md")
}

func TeachingHandler(w http.ResponseWriter, r *http.Request) {
	RenderMarkdown(w, r, "Teaching", "teaching.md")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.tmpl", nil)
}

func OtherHandler(w http.ResponseWriter, r *http.Request) {
	RenderMarkdown(w, r, "Other", "other.md")
}

func KeybaseVerificationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", keybase)
}

func MemeHandler(w http.ResponseWriter, r *http.Request) {
	const memePath = "static/images/memes/*"
	memeFiles, err := filepath.Glob(memePath)
	if err != nil {
		w.WriteHeader(500)
	}
	renderTemplate(w, "memes.tmpl", memeFiles)
}

func AboutMeHandler(w http.ResponseWriter, r *http.Request) {
	var title = "About Me"
	renderMarkdown(w, r, title, garbage)
}

var garbage = `
# About me

**My name:** Robbie McKinstry

__Likes:__` + "`coding`" + `

dislikes: shitty web frameworks`

var keybase = `==================================================================
https://keybase.io/thesnowmancometh
--------------------------------------------------------------------

I hereby claim:

  * I am an admin of https://robbiemckinstry.tech
  * I am thesnowmancometh (https://keybase.io/thesnowmancometh) on keybase.
  * I have a public key with fingerprint FCF0 C407 F1E2 1EDA 368B  6F15 8D24 F66B C054 B51C

To do so, I am signing this object:

{
  "body": {
    "key": {
      "eldest_kid": "0120a674099a2189e807b00e175a95c6a1eee9e65eed74ced3c7d266cdfb78337ccc0a",
      "fingerprint": "fcf0c407f1e21eda368b6f158d24f66bc054b51c",
      "host": "keybase.io",
      "key_id": "8d24f66bc054b51c",
      "kid": "0101a0a27c2b80407642411b92d7a71c7651f0a1fd4e17f9bdf57ff8eb8cc984cea40a",
      "uid": "6d054f555c995a6ee4fc74682dfda019",
      "username": "thesnowmancometh"
    },
    "service": {
      "hostname": "robbiemckinstry.tech",
      "protocol": "https:"
    },
    "type": "web_service_binding",
    "version": 1
  },
  "ctime": 1529987986,
  "expire_in": 157680000,
  "prev": "fa51f5c8bd09c57bc8206ee2448d11ead73cef54965123959963ef97f5405b86",
  "seqno": 17,
  "tag": "signature"
}

which yields the signature:

-----BEGIN PGP MESSAGE-----
Version: Keybase OpenPGP v2.0.77
Comment: https://keybase.io/crypto

yMN+AnicbVJ7UFRVHN6VV/IQqIEGRhbnxhQxSPfevc8FGYmCzCmJAQYRZ7n33HN3
b8De9e7d3dCwAlEIscmYxAdoOiDOMGAZqaMoCE0vHo4zhYhlINPIZEqZj4Kycxn9
r/PPmfM73/f9vu93zlBEgCnU3JR0r7y/oHe5+buBTq9pAzG4ewsmqlIVZtuClcPF
DVZI0KPbyxUJs2E4QeICw1I4zwskwfGQw1kRxyHB0gJPA0YgIIQ8ZGgIJZYCULIC
ViIZBkiyyHJWKwsAwAUsFZMVlwNqbk1x6UhWBjIOKJyVCUgSUBKsDCcyMkFzEknJ
DCMCnKZEmgCI6FQ9BgOZEwUPTFNUVEMH+6K9/8E/9o0TAi6QLCBFDkedGIqkCELk
SYkVWAKwDE3IuEDIEoWyyLwoyTQryxwUOQB4DiURqEXf3kU5RkINZJqmAc/TAgMh
JQOWYjhSkiUBJ3gD6IGaS6iECK07ocel+isFF1Aroe7EqlMxdOtTADQGbCR6hNRU
UVRgJShXXB5dq0rTIXAiLbem6ipQKxDCqetuj81Q0KvcBsUPRfsjMbuouCQ0V8Tw
Qc2jqC7MRiAk0BVDnaBJnudYnmNSMfiWW9GgXTEQNMtwOFpGH+gzXkNAw6ABJ0o4
D2hWBByJo4wkRXESQUBBYq0AyjTFo6GRVp7mecYKZZ5FJZwWOQYz4m1yqUibRT4F
B9L0KA6XoHs1iFVfGCgNNJlDTcFBS4wPZwpdGvX4GybaIh8WRsZ6uy3JcdlXD09N
duh1icLynO6xBXvz/hPHCoZOuvQJsiHlbHgC54uMGf+kJCEv+Ieclq6+Unrqbol/
u1Yc03wm3zPW373+w6zrxy8evkK3pMzlX4tSGrNGduVOiOMdvgvepPUxZZc/LRps
MpfX4qasuo6ZdzZrsZn6s+zgkZ8Ptp4c35eoNC/btWJ+fm992VeXq3sG+trdMyXe
NOZQmFC0b7ZBYE93yQ0XM1+ZLH43Y4GkyPN/bo1uXWmJLU16zv90Z3VD4tEfYWZQ
QcQOatiae0K60ld7e+PYyDX910PrMn66MdUb3jNm3eEFp8wx90ZDjswesDxV9PBO
iv/9Uy+suV878vdHG1eOyttaudmZY9hczcKLUZfyXw4M/J6MSb6SM1QXUqR2Wfzp
a1sinjk//PyqoZ4D9uy0md+pyi+PXl/RePNu1pmUt/NC8m8u+/w9Xzt57i/rZN6a
cw61c7Pjm9HswmkpeXfD68ye2VvObf/eb58ozBWj32gJD7tTV3y6U6Tb1lEPDgY7
8PmmsfqgtU+81ui59EXY6r2/SR80xTuphAf1b2qfpce/am27Nb1pMPqfnZFPhgz3
7zn7R9x018erBlIKY79NcFcsKdE33D6ePtCbvzrglzB/89WanMF2X+ENpfHriKVw
e02bJd4VblmIs3XuDNifHQC2Us6MrjIz/tLC3H8GY8Bz
=/T9o
-----END PGP MESSAGE-----

And finally, I am proving ownership of this host by posting or
appending to this document.

View my publicly-auditable identity here: https://keybase.io/thesnowmancometh

==================================================================`
