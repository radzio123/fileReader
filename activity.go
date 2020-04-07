package flogoFileRead

import (
	"io/ioutil"
	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivSourceFile      = "sourceFile"
	ovResult          = "result"
)

var log = logger.GetLogger("activity-gzip")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done string, err error) {
	//Get constances
	sourceFile = context.GetInput(ivSourceFile).(string)

	log.Debugf("[FlogoFileRead]: %s", sourceFile)

	data, err := ioutil.ReadFile(sourceFile)

	context.SetOutput(ovResult, string(data))
	return string(data), err
}