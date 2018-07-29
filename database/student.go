package database

type Student struct {
	first string
	last  string
	grade uint
}

var Roster = map[string][]Student{
	"fall 2018": []Student{{
		first: "Dominic",
		last:  "Augello",
		grade: 12,
	}, {
		first: "Renne",
		last:  "Cabacungan",
		grade: 12,
	}, {
		first: "Zachary",
		last:  "Evens",
		grade: 11,
	}, {
		first: "William",
		last:  "Fello",
		grade: 11,
	}, {
		first: "Luke",
		last:  "Krasinski",
		grade: 12,
	}, {
		first: "Jonathan",
		last:  "Krystopolski-Czernics",
		grade: 12,
	}, {
		first: "Jack",
		last:  "McAllister",
		grade: 11,
	}, {
		first: "Neil",
		last:  "Shipley",
		grade: 12,
	}, {
		first: "John",
		last:  "Speers",
		grade: 11,
	}, {
		first: "Liam",
		last:  "Sullivan",
		grade: 12,
	}},
	"spring 2019": []Student{{
		first: "Jacob",
		last:  "Billisits",
		grade: 12,
	}, {
		first: "Luke",
		last:  "Crist",
		grade: 10,
	}, {
		first: "Ryan",
		last:  "Haver",
		grade: 11,
	}, {
		first: "Christian",
		last:  "Farls",
		grade: 10,
	}, {
		first: "Thomas",
		last:  "Kujawinski",
		grade: 11,
	}, {
		first: "Simon",
		last:  "Herbert",
		grade: 12,
	}, {
		first: "Timothy",
		last:  "McClelland",
		grade: 12,
	}, {
		first: "Sykai",
		last:  "Napper",
		grade: 10,
	}, {
		first: "Owen",
		last:  "O'Malley",
		grade: 10,
	}, {
		first: "Michael",
		last:  "Pistolesi",
		grade: 11,
	}, {
		first: "Dominic",
		last:  "Polsinelli",
		grade: 11,
	}, {
		first: "Connor",
		last:  "Sweeney",
		grade: 12,
	}, {
		first: "Richard",
		last:  "Tillman",
		grade: 12,
	}},
}
