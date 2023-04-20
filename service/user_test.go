package service

import (
	"testing"
)

func TestVerifyRole(t *testing.T) {
	uSrv := UserService{}
	tests := []struct {
		name   string
		role   string
		expErr error
	}{
		{
			name:   "valid admin role",
			role:   "admin",
			expErr: nil,
		},
		{
			name:   "valid moderator role",
			role:   "moderator",
			expErr: nil,
		},
		{
			name:   "valid user role",
			role:   "user",
			expErr: nil,
		},
		{
			name:   "invalid role",
			role:   "admin1",
			expErr: ErrRoleDoesNotExist,
		},
	}

	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			err := uSrv.VerifyRole(tCase.role)
			if err != tCase.expErr {
				t.Errorf("Expected error %v, but got %v", tCase.expErr, err)
			}
		})
	}
}
