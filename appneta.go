package logrus_appneta

import (
	"github.com/Sirupsen/logrus"
	"github.com/tracelytics/go-traceview/v1/tv"
	"golang.org/x/net/context"
)

const defaultErrorClass = "error"
const defaultLayerName = "logrus_appneta"

var defaultLevels = []logrus.Level{
	logrus.PanicLevel,
	logrus.FatalLevel,
	logrus.ErrorLevel,
}

// AppnetaHook sends error event to AppNeta TraceView
type AppnetaHook struct {
	FieldPrefix string
	ErrorClass  string
	LayerName   string

	levels []logrus.Level
}

// NewHook returns initialized *AppnetaHook
func NewHook() *AppnetaHook {
	return &AppnetaHook{
		ErrorClass: defaultErrorClass,
		LayerName:  defaultLayerName,
		levels:     defaultLevels,
	}
}

// Fire sends error event
func (hook *AppnetaHook) Fire(entry *logrus.Entry) error {
	d := entry.Data
	ec := hook.getErrorClass(d)
	msg := hook.getErrorMessage(entry)

	if layer, ok := hook.getLayer(d); ok {
		layer.Error(ec, msg)
		return nil
	}

	if ctx, ok := hook.getContext(d); ok {
		l, _ := tv.BeginLayer(ctx, hook.getLayerName(d))
		defer l.End()
		l.Error(ec, msg)
		return nil
	}

	return nil
}

func (hook *AppnetaHook) getErrorClass(d logrus.Fields) string {
	const key = "error_class"

	if value, ok := d[hook.FieldPrefix+key].(string); ok {
		return value
	}
	return hook.ErrorClass
}

func (hook *AppnetaHook) getErrorMessage(entry *logrus.Entry) string {
	const key = "error"

	if err, ok := entry.Data[hook.FieldPrefix+key].(error); ok {
		return err.Error()
	}
	return entry.Message
}

func (hook *AppnetaHook) getLayer(d logrus.Fields) (tv.Layer, bool) {
	const key = "layer"

	if value, ok := d[hook.FieldPrefix+key].(tv.Layer); ok {
		return value, true
	}
	return nil, false
}

func (hook *AppnetaHook) getLayerName(d logrus.Fields) string {
	const key = "layer_name"

	if value, ok := d[hook.FieldPrefix+key].(string); ok {
		return value
	}
	return hook.LayerName
}

func (hook *AppnetaHook) getContext(d logrus.Fields) (context.Context, bool) {
	const key = "context"

	if value, ok := d[hook.FieldPrefix+key].(context.Context); ok {
		return value, true
	}
	return nil, false
}

// Levels returns logging level to fire hook
func (hook *AppnetaHook) Levels() []logrus.Level {
	return hook.levels
}

// SetLevels sets logging level to fire hook
func (hook *AppnetaHook) SetLevels(levels []logrus.Level) {
	hook.levels = levels
}
