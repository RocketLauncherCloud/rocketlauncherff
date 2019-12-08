package core

import "time"

type FeatureFlag struct {
	Id          string
	Name        string
	Description string
	Enabled     bool
	created     time.Time
	updated     time.Time
}
