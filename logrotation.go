package logrotation

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logrotation struct {
	BaseFilename     string        // base filename, final name will be baseFilename.YYYYMMDD.suffix
	Suffix           string        // log filename suffix
	UseDateTree      bool          // use a date based directory structur, i.e YYYY/MM/DD/bla.suffix
	Interval         time.Duration // how often to rotate
	BaseDir          string        // defaults to "./"
	currentFileName  string
	currentStartTime time.Time
	currentFile      *os.File
	mu               sync.Mutex
}

func (l *Logrotation) makePathString(now time.Time) string {
	var baseDir string
	if l.BaseDir == "" {
		baseDir = "."
	} else {
		baseDir = l.BaseDir // possibly remove suffix / ?
	}

	if l.UseDateTree {
		year := now.Format("2006")
		month := now.Format("01")
		day := now.Format("02")
		dir := fmt.Sprintf("%s/%s/%s", year, month, day)
		dir = filepath.FromSlash(dir)
		baseDir = baseDir + "/" + dir
	}
	return baseDir
}
func (l *Logrotation) createPath(now time.Time) error {
	dir := l.makePathString(now)
	return os.MkdirAll(dir, 0744)

}
func (l *Logrotation) makeFN(now time.Time) string {
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	hour := now.Format("15")
	minute := now.Format("04")
	second := now.Format("05")

	fn := fmt.Sprintf("%s.%s%s%s-%s%s%s.%s", l.BaseFilename, year, month, day, hour, minute, second, l.Suffix)
	baseDir := l.makePathString(now)
	fn = fmt.Sprintf("%s/%s", baseDir, fn)
	fn = filepath.FromSlash(fn)
	return fn

}
func (l *Logrotation) openNewFile() error {
	now := time.Now().UTC()
	fn := l.makeFN(now)
	err := l.createPath(now)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.currentFile = f
	l.currentFileName = fn
	l.currentStartTime = now
	return nil
}

func (l *Logrotation) Write(bs []byte) (n int, e error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if (time.Since(l.currentStartTime) > l.Interval) || l.currentFile == nil {
		if l.currentFile != nil {
			if err := l.currentFile.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "error closing %s (%v)\n", l.currentFileName, err)
			}
		}
		if e = l.openNewFile(); e != nil {
			return 0, e
		}
	}
	return l.currentFile.Write(bs)
}

func (l *Logrotation) Close() error {
	return l.currentFile.Close()
}
