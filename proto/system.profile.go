package proto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SystemProfile holds the data structure of the default profiler document
// docsExamined is renamed from nscannedObjects in 3.2.0
// https://docs.mongodb.com/manual/reference/database-profiler/#system.profile.docsExamined
type SystemProfile struct {
	AllUsers        []interface{} `bson:"allUsers"`
	Client          string        `bson:"client"`
	CursorExhausted bool          `bson:"cursorExhausted"`
	DocsExamined    int           `bson:"docsExamined"`
	NscannedObjects int           `bson:"nscannedObjects"`
	ExecStats       struct {
		Advanced                    int `bson:"advanced"`
		ExecutionTimeMillisEstimate int `bson:"executionTimeMillisEstimate"`
		InputStage                  struct {
			Advanced                    int    `bson:"advanced"`
			Direction                   string `bson:"direction"`
			DocsExamined                int    `bson:"docsExamined"`
			ExecutionTimeMillisEstimate int    `bson:"executionTimeMillisEstimate"`
			Filter                      struct {
				Date struct {
					Eq string `bson:"$eq"`
				} `bson:"date"`
			} `bson:"filter"`
			Invalidates  int    `bson:"invalidates"`
			IsEOF        int    `bson:"isEOF"`
			NReturned    int    `bson:"nReturned"`
			NeedTime     int    `bson:"needTime"`
			NeedYield    int    `bson:"needYield"`
			RestoreState int    `bson:"restoreState"`
			SaveState    int    `bson:"saveState"`
			Stage        string `bson:"stage"`
			Works        int    `bson:"works"`
		} `bson:"inputStage"`
		Invalidates  int    `bson:"invalidates"`
		IsEOF        int    `bson:"isEOF"`
		LimitAmount  int    `bson:"limitAmount"`
		NReturned    int    `bson:"nReturned"`
		NeedTime     int    `bson:"needTime"`
		NeedYield    int    `bson:"needYield"`
		RestoreState int    `bson:"restoreState"`
		SaveState    int    `bson:"saveState"`
		Stage        string `bson:"stage"`
		Works        int    `bson:"works"`
	} `bson:"execStats"`
	KeyUpdates   int `bson:"keyUpdates"`
	KeysExamined int `bson:"keysExamined"`
	Locks        struct {
		Collection struct {
			AcquireCount struct {
				R int `bson:"R"`
			} `bson:"acquireCount"`
		} `bson:"Collection"`
		Database struct {
			AcquireCount struct {
				R int `bson:"r"`
			} `bson:"acquireCount"`
		} `bson:"Database"`
		Global struct {
			AcquireCount struct {
				R int `bson:"r"`
			} `bson:"acquireCount"`
		} `bson:"Global"`
		MMAPV1Journal struct {
			AcquireCount struct {
				R int `bson:"r"`
			} `bson:"acquireCount"`
		} `bson:"MMAPV1Journal"`
	} `bson:"locks"`
	Millis             int       `bson:"millis"`
	Nreturned          int       `bson:"nreturned"`
	Ns                 string    `bson:"ns"`
	NumYield           int       `bson:"numYield"`
	Op                 string    `bson:"op"`
	Protocol           string    `bson:"protocol"`
	Query              bson.D    `bson:"query"`
	UpdateObj          bson.D    `bson:"updateobj"`
	Command            bson.D    `bson:"command"`
	OriginatingCommand bson.D    `bson:"originatingCommand"`
	ResponseLength     int       `bson:"responseLength"`
	TS                 time.Time `bson:"ts"`
	User               string    `bson:"user"`
	WriteConflicts     int       `bson:"writeConflicts"`
}
