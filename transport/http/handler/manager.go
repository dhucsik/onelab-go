package handler

import (
	"practice/service"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(srv *service.Manager) *Manager {
	return &Manager{
		srv: srv,
	}
}
