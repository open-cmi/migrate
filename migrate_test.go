package migrate

import "testing"

func Test_IsMigrateCommand(t *testing.T) {
	if !IsMigrateCommand("help") {
		t.Errorf("test migrate command error\n")
	}
	if !IsMigrateCommand("list") {
		t.Errorf("test migrate command error\n")
	}
	if !IsMigrateCommand("generate") {
		t.Errorf("test migrate command error\n")
	}
	if !IsMigrateCommand("up") {
		t.Errorf("test migrate command error\n")
	}
	if !IsMigrateCommand("down") {
		t.Errorf("test migrate command error\n")
	}
	if !IsMigrateCommand("current") {
		t.Errorf("test migrate command error\n")
	}
	if IsMigrateCommand("test") {
		t.Errorf("test migrate command error\n")
	}
}
