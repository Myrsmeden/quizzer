package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"time"
)

type Answer struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type Question struct {
	Enquiry string   `json:"question"`
	Answers []Answer `json:"answers"`
}

func GetQuestion(db *sql.DB) Question {
	// We have a different number of types of questions
	// 1. Connect attribute with person
	// 2. Connect person with attribute
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	if randGenerator.Intn(100) > 50 {
		return GetPersonQuestion(db)
	} else {
		return GetPersonQuestion(db)
	}
}
func GetPersonQuestion(db *sql.DB) Question {

	// First get the question text
	sql := "SELECT id, name FROM persons ORDER BY RANDOM() LIMIT 1"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	question := Question{}
	id := 0
	name := ""
	for rows.Next() {

		err2 := rows.Scan(&id, &name)
		if err2 != nil {
			panic(err2)
		}
		question.Enquiry = "Vad gäller för " + name + "?"
	}

	// Now get the correct alternative

	sql = "SELECT info FROM attributes WHERE person_id = ? ORDER BY RANDOM() LIMIT 1"
	rows, err3 := db.Query(sql, id)

	if err3 != nil {
		panic(err3)
	}

	defer rows.Close()

	for rows.Next() {
		answer := Answer{}
		err4 := rows.Scan(&answer.Text)
		if err4 != nil {
			panic(err4)
		}
		answer.Correct = true
		question.Answers = append(question.Answers, answer)
	}

	// Now get three incorrect answers
	sql = "SELECT info FROM attributes WHERE person_id != ? ORDER BY RANDOM() LIMIT 3"
	rows, err = db.Query(sql, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		answer := Answer{}
		err5 := rows.Scan(&answer.Text)
		if err5 != nil {
			panic(err5)
		}

		answer.Correct = false
		question.Answers = append(question.Answers, answer)
	}
	Shuffle(question.Answers)
	return question
}

func Shuffle(slc []Answer) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}
