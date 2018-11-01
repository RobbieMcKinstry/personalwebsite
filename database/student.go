package database

type Student struct {
	First      string
	Last       string
	grade      uint
	QuizGrades []*Score
	StudentID  string
}

type Score struct {
	Date   string
	Score  float64
	Total  float64
	Number int
}

var Roster = map[string][]Student{
	"fall 2018": []Student{{
		First:     "Dominic",
		Last:      "Augello",
		grade:     12,
		StudentID: "20192606",
	}, {
		First:     "Renne",
		Last:      "Cabacungan",
		grade:     12,
		StudentID: "20192701",
	}, {
		First:     "William",
		Last:      "Fello",
		grade:     11,
		StudentID: "20201960",
	}, {
		First:     "Wylie",
		Last:      "Haddad",
		grade:     12,
		StudentID: "20192961",
	}, {
		First:     "Fantini",
		Last:      "Mike",
		grade:     12,
		StudentID: "20192886",
	}, {
		First:     "Alan",
		Last:      "Kraning",
		grade:     12,
		StudentID: "20193076",
	}, {
		First:     "Luke",
		Last:      "Krasinski",
		grade:     12,
		StudentID: "20193081",
	}, {
		First:     "Jonathan",
		Last:      "Krystopolski-Czernics",
		grade:     12,
		StudentID: "20193086",
	}, {
		First:     "Neil",
		Last:      "Shipley",
		grade:     12,
		StudentID: "20193446",
	}, {
		First:     "Connor",
		Last:      "Sweeney",
		grade:     12,
		StudentID: "20193536",
	}, {
		First:     "Jack",
		Last:      "Rooney",
		grade:     12,
		StudentID: "20193391",
	}, {
		First:     "Liam",
		Last:      "Sullivan",
		grade:     12,
		StudentID: "20193521",
	}},
	"spring 2019": []Student{{
		First: "Jacob",
		Last:  "Billisits",
		grade: 12,
	}, {
		First: "Luke",
		Last:  "Crist",
		grade: 10,
	}, {
		First: "Ryan",
		Last:  "Haver",
		grade: 11,
	}, {
		First: "Christian",
		Last:  "Farls",
		grade: 10,
	}, {
		First: "Thomas",
		Last:  "Kujawinski",
		grade: 11,
	}, {
		First: "Simon",
		Last:  "Herbert",
		grade: 12,
	}, {
		First: "Timothy",
		Last:  "McClelland",
		grade: 12,
	}, {
		First: "Sykai",
		Last:  "Napper",
		grade: 10,
	}, {
		First: "Owen",
		Last:  "O'Malley",
		grade: 10,
	}, {
		First: "Michael",
		Last:  "Pistolesi",
		grade: 11,
	}, {
		First: "Dominic",
		Last:  "Polsinelli",
		grade: 11,
	}, {
		First: "Zachary",
		Last:  "Evens",
		grade: 11,
	}, {
		First: "Richard",
		Last:  "Tillman",
		grade: 12,
	}},
}
