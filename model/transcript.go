package model

import (
	"gopkg.in/guregu/null.v3"
)

type Transcript struct {
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
	User	null.Int	`json:"user" dont:"u" key:"p"`
	File	null.String	`json:"file" dont:"u"`
	Confirm	null.Bool	`json:"confirm"`
	Time	null.Time	`json:"time"`
}

type TranscriptWithUser struct {
	Tour	null.Int	`json:"tour" key:"p"`
	File	null.String	`json:"file"`
	Confirm	null.Bool	`json:"confirm"`
	Time	null.Time	`json:"time"`
	Name	null.String	`json:"name"`
}

type TranscriptWithTour struct {
	User	null.Int	`json:"user" key:"p"`
	File	null.String	`json:"file"`
	Confirm	null.Bool	`json:"confirm"`
	Time	null.Time	`json:"time"`
	Name	null.String	`json:"name"`
	Surname	null.String	`json:"surname"`
}