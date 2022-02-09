package migrate

import "testing"

func Test_IsSubCommand(t *testing.T) {

	if !IsSubCommand("list") {
		t.Errorf("test migrate command error\n")
	}
	if !IsSubCommand("generate") {
		t.Errorf("test migrate command error\n")
	}
	if !IsSubCommand("up") {
		t.Errorf("test migrate command error\n")
	}
	if !IsSubCommand("down") {
		t.Errorf("test migrate command error\n")
	}
	if !IsSubCommand("current") {
		t.Errorf("test migrate command error\n")
	}
	if IsSubCommand("test") {
		t.Errorf("test migrate command error\n")
	}
}
