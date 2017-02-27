package cli

import (
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	green  = 32
	red    = 31
	yellow = 33
	blue   = 34
)

type CBFormatter struct {
}

func (f *CBFormatter) Format(entry *log.Entry) ([]byte, error) {

	var keys []string
	for k := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	b := &bytes.Buffer{}

	prefixFieldClashes(entry)

	if log.IsTerminal() {
		printColored(b, entry, keys)
	} else {
		f.appendKeyValue(b, "time", entry.Time.Format(time.RFC3339))
		f.appendKeyValue(b, "level", entry.Level.String())
		f.appendKeyValue(b, "msg", entry.Message)
		for _, key := range keys {
			f.appendKeyValue(b, key, entry.Data[key])
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

type ContextHook struct{}

func (hook ContextHook) Levels() []log.Level {
	return []log.Level{
		log.WarnLevel,
		log.InfoLevel,
		log.DebugLevel,
	}
}

func (hook ContextHook) Fire(entry *log.Entry) error {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(9, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/logrus") {
			_, line := fu.FileLine(pc[i] - 1)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			break
		}
	}
	return nil
}

func printColored(b *bytes.Buffer, entry *log.Entry, keys []string) {
	var levelColor int
	switch entry.Level {
	case log.WarnLevel:
		levelColor = yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = red
	case log.InfoLevel:
		levelColor = green
	default:
		levelColor = blue
	}

	levelText := strings.ToUpper(levelToString(entry))

	if runtime.GOOS == "windows" {
		fmt.Fprintf(b, "%-6s %-44s ", levelText, entry.Message)
		for _, k := range keys {
			v := entry.Data[k]
			fmt.Fprintf(b, "%s=%v", k, v)
		}
	} else {
		fun, line := getLocationInfo(entry)
		location := ""
		if len(fun) > 0 {
			location = fmt.Sprintf("[%s:%s]", fun, line)
		}
		fmt.Fprintf(b, "\x1b[%dm%s%s: \x1b[0m%-44s ", levelColor, levelText, location, entry.Message)
		//for _, k := range keys {
		//	v := entry.Data[k]
		//	fmt.Fprintf(b, " \x1b[%dm%s\x1b[0m=%v", levelColor, k, v)
		//}
	}
}

func getLocationInfo(entry *log.Entry) (fun, line string) {
	if _, ok := entry.Data["func"]; ok {
		f := entry.Data["func"].(string)
		funcParts := strings.Split(f, ".")
		for i := len(funcParts) - 1; i >= 0; i-- {
			segment := funcParts[i]
			if !strings.HasPrefix(segment, "func") {
				fun = segment
				break
			}
		}
	}
	if _, ok := entry.Data["line"]; ok {
		line = strconv.Itoa(entry.Data["line"].(int))
	}
	return fun, line
}

func levelToString(entry *log.Entry) string {
	switch entry.Level {
	case log.DebugLevel:
		return "debug"
	case log.InfoLevel:
		return "info"
	case log.WarnLevel:
		return "warn"
	case log.ErrorLevel:
		return "error"
	case log.FatalLevel:
		return "fatal"
	case log.PanicLevel:
		return "panic"
	}
	return "error"
}

func (f *CBFormatter) appendKeyValue(b *bytes.Buffer, key, value interface{}) {
	switch value.(type) {
	case string, error:
		fmt.Fprintf(b, "%v=%q ", key, value)
	default:
		fmt.Fprintf(b, "%v=%v ", key, value)
	}
}

func prefixFieldClashes(entry *log.Entry) {
	_, ok := entry.Data["time"]
	if ok {
		entry.Data["fields.time"] = entry.Data["time"]
	}

	_, ok = entry.Data["msg"]
	if ok {
		entry.Data["fields.msg"] = entry.Data["msg"]
	}

	_, ok = entry.Data["level"]
	if ok {
		entry.Data["fields.level"] = entry.Data["level"]
	}
}
