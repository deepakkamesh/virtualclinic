// Package script handles prescriptions (aka scripts).
package script

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Patient struct {
	Name  string    // Name of Patient (mandatory).
	DOB   time.Time // Date of birth (mandatory).
	Sex   string    // Sex of patient (mandatory). Acceptable values M/F
	Phone string    // Phone number.
}

// Script struct represents a single script. All string should be lowercase.
type Script struct {
	TimeStamp time.Time // Time of the script.
	Patient   *Patient  // Patient details.
	Script    string    // Prescription details.
	ScriptID  string    // Unique chracter based script ID.
}

type ScriptManager struct {
	scriptIDLen int     // Length of script ID string.
	timezone    string  // Time zone for date time.
	db          *sql.DB // SQLLIte DB driver.
}

// NewScriptManager returns a new initialized Scripts struct.
func NewScriptManager(scriptIDLen int, timezone, dbfile string) (*ScriptManager, error) {
	db, err := sql.Open("sqlite", dbfile) // see https://github.com/mattn/go-sqlite3/blob/master/README.md#faq
	if err != nil {
		return nil, err
	}

	return &ScriptManager{
		scriptIDLen: scriptIDLen,
		timezone:    timezone,
		db:          db,
	}, nil
}

/* CREATE TABLE scripts (timestamp DATETIME NOT NULL,
name TEXT NOT NULL, dob DATETIME NOT NULL, sex TEXT NOT NULL,
 phone TEXT, script TEXT NOT NULL, scriptid TEXT NOT NULL);*/

// NewScript saves script in DB and returns the script id.
// script ID is a random 10 character string.
func (s *ScriptManager) NewScript(sc *Script) (string, error) {

	//validate Sane values.
	if sc.TimeStamp.IsZero() || sc.Patient.Name == "" ||
		sc.Patient.DOB.IsZero() || sc.Patient.Sex == "" ||
		sc.Script == "" || !(sc.Patient.Sex == "M" || sc.Patient.Sex == "F") {
		return "", fmt.Errorf("cannot have null values for mandatory fields of a script")
	}

	scriptID, err := randomStringCrypto(10)
	if err != nil {
		return "", err
	}

	if _, err = s.db.Exec("INSERT INTO scripts VALUES(?,?,?,?,?,?,?);",
		sc.TimeStamp, sc.Patient.Name, sc.Patient.DOB, sc.Patient.Sex, sc.Patient.Phone, sc.Script, scriptID); err != nil {
		return "", err
	}

	return scriptID, nil
}

// DeleteScript deletes the script ID from the database.
func (s *ScriptManager) DeleteScript(scriptID string) error {
	if _, err := s.db.Exec("DELETE from scripts WHERE scriptid=?", scriptID); err != nil {
		return err
	}
	return nil
}

// DeleteAllScripts deletes all scripts.
func (s *ScriptManager) DeleteAllScripts() error {
	if _, err := s.db.Exec("DELETE from scripts"); err != nil {
		return err
	}
	return nil
}

// Script returns the script for the scriptID.
func (s *ScriptManager) Script(scriptID string) (*Script, error) {
	qry := "SELECT * FROM scripts WHERE scriptid=?"
	rows, err := s.db.Query(qry, scriptID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sc := Script{Patient: &Patient{}}

	if !rows.Next() {
		return nil, fmt.Errorf("record not found")
	}
	if err := rows.Scan(&sc.TimeStamp, &sc.Patient.Name, &sc.Patient.DOB, &sc.Patient.Sex, &sc.Patient.Phone, &sc.Script, &scriptID); err != nil {
		return nil, err
	}
	return &sc, nil
}

// FindPatient returns patient details for any substring matches of nameSubstr. Empty input will
// return all Patient records.
func (s *ScriptManager) FindPatient(nameSubstr string) ([]Patient, error) {

	qry := "SELECT DISTINCT name,dob,sex,phone FROM scripts WHERE name LIKE '%" + nameSubstr + "%'"
	rows, err := s.db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	patients := []Patient{}

	for rows.Next() {
		p := Patient{}
		if err := rows.Scan(&p.Name, &p.DOB, &p.Sex, &p.Phone); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, nil
}

// Scripts returns all scripts for a Patient. It works by doing exact match
// of Name, DOB & Sex
func (s *ScriptManager) Scripts(p Patient, limit int) ([]Script, error) {
	qry := "SELECT * FROM scripts WHERE name=? AND dob=? AND sex=? LIMIT ?"
	rows, err := s.db.Query(qry, p.Name, p.DOB, p.Sex, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	scripts := []Script{}

	for rows.Next() {
		s := Script{Patient: &Patient{}}
		if err := rows.Scan(&s.TimeStamp, &s.Patient.Name, &s.Patient.DOB, &s.Patient.Sex, &s.Patient.Phone, &s.Script, &s.ScriptID); err != nil {
			return nil, err
		}

		scripts = append(scripts, s)
	}

	return scripts, nil

}

// randomStringCrypto generates a random string of length.
func randomStringCrypto(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}
