package script_test

import (
	"testing"
	"time"

	"github.com/deepakkamesh/virtualclinic/script"
)

func TestSaveScript(t *testing.T) {
	sm, err := script.NewScriptManager(10, "Asia/Kolkata", "scripts_test.db")
	if err != nil {
		t.Errorf("failed to create scriptsDB:%v", err)
	}
	loc, _ := time.LoadLocation("Asia/Kolkata")
	timeStamp := time.Now().In(loc)

	scriptID, err := sm.NewScript(&script.Script{
		TimeStamp: timeStamp,
		Patient: &script.Patient{
			Name:  "Muthu",
			DOB:   time.Date(1980, 1, 3, 0, 0, 0, 0, loc),
			Sex:   "M",
			Phone: "919812345",
		},
		Script: "Take Paracetamol \n something else \n",
	})
	if err != nil {
		t.Errorf("failed to save record")
	}

	s, err := sm.Script(scriptID)
	if err != nil {
		t.Errorf("Failed to get Script %v", err)
	}

	if s.Patient.Name != "Muthu" ||
		!s.Patient.DOB.Equal(time.Date(1980, 1, 3, 0, 0, 0, 0, loc)) ||
		s.Patient.Sex != "M" ||
		s.Patient.Phone != "919812345" ||
		s.Script != "Take Paracetamol \n something else \n" {
		t.Errorf("Got %v, Want something else", s)
	}
	if err := sm.DeleteScript(scriptID); err != nil {
		t.Errorf("Failed to delete scriptID %v", err)
	}
}

func TestFindPatient(t *testing.T) {
	s, err := script.NewScriptManager(10, "Asia/Kolkata", "scripts_test.db")
	if err != nil {
		t.Errorf("failed to create scriptsDB:%v", err)
	}
	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)

	muthu := script.Patient{"muthu", time.Date(1980, 1, 3, 0, 0, 0, 0, loc), "M", "1"}
	muthu2 := script.Patient{"muthukumar", time.Date(1980, 1, 3, 0, 0, 0, 0, loc), "M", ""}

	malar := script.Patient{"malar", time.Date(1994, 1, 1, 0, 0, 0, 0, loc), "F", "2"}
	satya := script.Patient{"satya", time.Date(1950, 1, 6, 0, 0, 0, 0, loc), "M", "3"}

	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &muthu,
			Script:    "script1 fr muthu",
		},
	)
	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &muthu,
			Script:    "script2 for muthu",
		},
	)
	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &muthu2,
			Script:    "script for muthu2",
		},
	)
	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &muthu2,
			Script:    "script for muthu2",
		},
	)
	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &malar,
			Script:    "script",
		},
	)
	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &satya,
			Script:    "script2 for muthu",
		},
	)

	patients, err := s.FindPatient("muthu")
	if err != nil {
		t.Errorf("Error finding patients %v", err)
	}

	if len(patients) != 2 {
		t.Errorf("Need 2 patients, Got %v", len(patients))
	}

	s.DeleteAllScripts()
}

func TestScripts(t *testing.T) {
	s, err := script.NewScriptManager(10, "Asia/Kolkata", "scripts_test.db")
	if err != nil {
		t.Errorf("failed to create scriptsDB:%v", err)
	}
	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)
	muthu := script.Patient{"muthu", time.Date(1980, 1, 3, 0, 0, 0, 0, loc), "M", "1"}
	muthu2 := script.Patient{"muthu", time.Date(1979, 1, 3, 0, 0, 0, 0, loc), "M", ""}

	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &muthu,
			Script:    "script1 fr muthu",
		},
	)
	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &muthu,
			Script:    "script2 for muthu",
		},
	)
	s.NewScript(
		&script.Script{
			TimeStamp: now,
			Patient:   &muthu2,
			Script:    "script for muthu2",
		},
	)

	scripts, err := s.Scripts(script.Patient{
		Name: "muthu",
		DOB:  time.Date(1980, 1, 3, 0, 0, 0, 0, loc),
		Sex:  "M",
	}, 10)
	if err != nil {
		t.Errorf("Failed to get scripts %v", err)
	}

	if len(scripts) != 2 {
		t.Errorf("Want 2 scripts, got more.")
	}

	s.DeleteAllScripts()
}
