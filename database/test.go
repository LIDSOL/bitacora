// database/test.go
package database

import (
	"database/sql"
	"math/rand"

	"github.com/jaswdr/faker/v2"
)

func GenerateFakeData(db *sql.DB, num int) error {
	fake := faker.New()
	var (
		u     user
		us    string
		p     project
		ps    string
		l     log
		ulist []string
		plist []string
	)

	// Users
	for i := 0; i < num; i++ {
		us = fake.RandomStringWithLength(13)
		u = user{
			userID:  us,
			name:    fake.Person().FirstName(),
			surname: fake.Person().LastName(),
			uType:   userTypes[rand.Intn(len(userTypes))],
			email:   fake.Person().Contact().Email,
		}
		if err := AddUser(db, u); err != nil {
			return err
		}
		ulist = append(ulist, us)
	}

	// Projects
	for i := 0; i < num; i++ {
		ps = fake.App().Name()
		p = project{
			name:        ps,
			manager:     ulist[rand.Intn(len(ulist))],
			description: fake.Lorem().Paragraph(rand.Intn(15)),
		}
		if err := AddProject(db, p); err != nil {
			return err
		}
		plist = append(plist, ps)
	}

	// Logs
	for i := 0; i < num*10; i++ {
		pID, _ := GetProjectID(db, plist[rand.Intn(len(plist))])

		l = log{
			userID:    ulist[rand.Intn(len(ulist))],
			projectID: pID,
		}
		if err := AddLog(db, l); err != nil {
			return err
		}
	}

	return nil
}
