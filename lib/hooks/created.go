package hooks

import (
	"github.com/guark/guark/app"

	"github.com/kodah/girc/irc/manager"
	"github.com/kodah/girc/store"
)

// App created hook.
func Created(a *app.App) error {
	a.Log.Info("---- HOOK: App created! ----")

	// open the sqlite database connection
	err := store.New()
	if err != nil {
		a.Log.Error("Error while starting the database: ", err)
	}

	// open a new irc connection manager
	manager.New()
	_, err = manager.Service.NewClient(
		"girc-test",
		"",
		"girc-test",
		"gIRC Test",
		"chat.freenode.net:6667",
		0,
		0,
		0,
		0)
	if err != nil {
		a.Log.Error("Error while connecting to server: ", err)
	}

	return nil
}
