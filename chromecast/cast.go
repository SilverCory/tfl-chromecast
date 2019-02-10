package cmd

import (
	"github.com/pkg/errors"
)

func Load(file string) error {
	app, err := castApplication()
	if err != nil {
		return err
	}

	if err := app.Load(file, "", false); err != nil {
		return errors.Wrap(err, "unable to load media")
	}
}
