package handlers

import "github.com/labstack/echo/v4"

// ╭───────────────────────────────────────────────────────────╮
// │                  Dexie Database Handlers                  │
// ╰───────────────────────────────────────────────────────────╯

func (a *syncAPI) FetchInitial(e echo.Context) error {
	// Implement database schema endpoint
	return nil
}

func (a *syncAPI) FetchCurrent(e echo.Context) error {
	// Implement account entries endpoint
	return nil
}

// ╭───────────────────────────────────────────────────────────╮
// │                 Group Structures                          │
// ╰───────────────────────────────────────────────────────────╯

type syncAPI struct{}

var Sync = new(syncAPI)
