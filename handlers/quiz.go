package handlers

import (
	"net/http"

	"github.com/RobbieMcKinstry/personal-website/database"
	"github.com/gorilla/mux"
)

func QuizHandler() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/quiz/{student}", studentHandler)
	return router
}

func findStudent(slug string, students []Student) (Student, bool) {
	var found = false
	var res database.Student
	for _, student := range students {
		if slug == student.StudentID {
			found = true
			res = student
			break
		}
	}
	return res, found
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var slug = vars["student"]
	students := database.Roster["fall 2018"]
	if student, ok := findStudent(); ok {
		renderTemplate(w, "quiz.tmpl", student)
	}

	http.NotFound(w, r)

}
