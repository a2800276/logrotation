# Simple file rotations

Usage: 

```
lr := Logrotation {
	"whatever",   // base filename: whatever.<YYYYMMDD-HHMMSS>.<suffix>
	"log",        // suffix       : <base filename>.<YYYYMMDD-HHMMSS>.log
	false,        // don't use a date based dir tree
	"./"          // base directory, defaults to "./"
}

log.SetOutput(lr)

```
## TODO

- more Tests
- max size
- compression and deletion
- license
