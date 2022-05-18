package metrics

import (
	"github.com/kataras/go-events"
)

const (
	ON_OBJECT_ADD    = events.EventName("object_added")
	ON_OBJECT_REMOVE = events.EventName("object_removed")

	ON_BLOB_ADD    = events.EventName("blob_added")
	ON_BLOB_REMOVE = events.EventName("blob_removed")
)

func RegisterEvents() {
	logger.Infof("Registering telemtry events")
	events.AddListener(ON_OBJECT_ADD, func(params ...interface{}) {
		objectsAdded.Inc()
	})

	events.AddListener(ON_OBJECT_REMOVE, func(params ...interface{}) {
		logger.Debug("Object Removed event handler triggered")
		objectsAdded.Dec()
	})

	events.AddListener(ON_BLOB_ADD, func(params ...interface{}) {
		logger.Debug("Object Removed event handler triggered")
		blobsAdded.Inc()
	})

	events.AddListener(ON_BLOB_REMOVE, func(params ...interface{}) {
		logger.Debug("Object Removed event handler triggered")
		blobsAdded.Dec()
	})
}
