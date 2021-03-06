package log_test

import (
	"bytes"
	stdlog "log"
	"os"
	"strings"
	"testing"

	"github.com/tprasadtp/log"

	"github.com/stretchr/testify/assert"
)

func TestStdLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	stdlog.SetOutput(buf)
	defer func() {
		stdlog.SetOutput(os.Stderr)
	}()

	lgr := log.NewDefault()

	// Make sure levels are working
	lgr.Debug("test debug")
	if strings.Contains(buf.String(), `[DEBUG] test debug`) {
		t.Log(buf.String())
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	// Test all levels
	lgr.Level = log.TraceLevel

	lgr.Trace("test debug")
	if !strings.Contains(buf.String(), `[TRACE] test debug`) {
		t.Log(buf.String())
		t.Error("stdlib trace not logging correctly")
	}
	buf.Reset()

	lgr.Tracef("Hello %s", "World")
	if !strings.Contains(buf.String(), `[TRACE] Hello World`) {
		t.Error("stdlib trace not logging correctly")
	}
	buf.Reset()

	lgr.Tracew("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[TRACE] foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib trace not logging correctly")
	}
	buf.Reset()

	lgr.Debug("test debug")
	if !strings.Contains(buf.String(), `[DEBUG] test debug`) {
		t.Log(buf.String())
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugf("Hello %s", "World")
	if !strings.Contains(buf.String(), `[DEBUG] Hello World`) {
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	lgr.Debugw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[DEBUG] foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib debug not logging correctly")
	}
	buf.Reset()

	lgr.Info("test info")
	if !strings.Contains(buf.String(), `[INFO ] test info`) {
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Infof("Hello %s", "World")
	if !strings.Contains(buf.String(), `[INFO ] Hello World`) {
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Infow("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[INFO ] foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib info not logging correctly")
	}
	buf.Reset()

	lgr.Warn("test warn")
	if !strings.Contains(buf.String(), `[WARN ] test warn`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnf("Hello %s", "World")
	if !strings.Contains(buf.String(), `[WARN ] Hello World`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Warnw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[WARN ] foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib warn not logging correctly")
	}
	buf.Reset()

	lgr.Error("test error")
	if !strings.Contains(buf.String(), `[ERROR] test error`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	lgr.Errorf("Hello %s", "World")
	if !strings.Contains(buf.String(), `[ERROR] Hello World`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	lgr.Errorw("foo bar", log.Fields{"baz": "qux"})
	if !strings.Contains(buf.String(), `[ERROR] foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("stdlib error not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "[PANIC] test panic", func() {
		lgr.Panic("test panic")
	})
	if !strings.Contains(buf.String(), `test panic`) {
		t.Log(buf.String())
		t.Error("cli panic not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "[PANIC] Hello World", func() {
		lgr.Panicf("Hello %s", "World")
	})
	if !strings.Contains(buf.String(), `Hello World`) {
		t.Log(buf.String())
		t.Error("cli panic not logging correctly")
	}
	buf.Reset()

	assert.PanicsWithValue(t, "[PANIC] foo bar [baz=qux]", func() {
		lgr.Panicw("foo bar", log.Fields{"baz": "qux"})
	})
	if !strings.Contains(buf.String(), `foo bar [baz=qux]`) {
		t.Log(buf.String())
		t.Error("cli panic not logging correctly")
	}
	buf.Reset()
}
