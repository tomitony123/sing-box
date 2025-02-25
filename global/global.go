package global

import "os"

var GolbalSignals = make(chan os.Signal, 1)
