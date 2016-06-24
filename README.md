logrus_appneta
====

[![Build Status](https://travis-ci.org/evalphobia/logrus_appneta.svg?branch=master)](https://travis-ci.org/evalphobia/logrus_appneta) [![codecov](https://codecov.io/gh/evalphobia/logrus_appneta/branch/master/graph/badge.svg)](https://codecov.io/gh/evalphobia/logrus_appneta)
 [![GoDoc](https://godoc.org/github.com/evalphobia/logrus_appneta?status.svg)](https://godoc.org/github.com/evalphobia/logrus_appneta)


# AppNeta TraceView Hook for Logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>

## Usage

```go
import (
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_appneta"
)

func main() {
	hook := logrus_appneta.NewHook()
	hook.SetLevels([]logrus.Level{
		logrus.PanicLevel,
		logrus.ErrorLevel,
	})

	logrus.AddHook(hook)
}
```


## Special fields

Some logrus fields have a special meaning in this hook.

|||
|:--|:--|
|`layer`|`layer` is `tv.Layer` type and used to send error. `layer` or `context` is required to send error.|
|`context`|`context` is `context.Context` type and used to send error. `layer` or `context` is required to send error.|
|`error`|`error` is `error` type and used for error message |
|`error_class`|`error_class` is `string` type and used for error class name|

These field can have original prefix.

```go
	hook := logrus_appneta.NewHook()
	hook.FieldPrefix = "your_prefix_"
	logrus.AddHook(hook)
```
