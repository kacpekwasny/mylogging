# mylogging

### Easy start

##### Single log file output and no output to console

```
package main

import (
	"errors"
	"fmt"

	"github.com/kacpekwasny/mylogging"
)

func main() {
	var (
	    filepath string = "/path/to/your/log/file"
	    // When file becomes bigger it will be moved to filepath + ".old.log"
	    // and a new filepath file will be created
	    fileSizeLimit int64 = 1024*1024*5
	    outputToConsole bool = false
	)
	lgr := mylogging.LoggerInit()
	lgr.WrToSingleFile(outPutToConsole, filepathstring, fileSizeLimit)

	lgr.Listen()

	lgr.Debug("My first log write!")
	lgr.Critical("CRITICAL MESSAGE! WATCH OUT", fmt.Sprintf("%#v", errors.New("New error Check this out")))
	lgr.Critical(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Urna et pharetra pharetra massa massa. Aenean euismod elementum nisi quis eleifend. Purus in mollis nunc sed id semper risus.`)

	lgr.StopListen()
}



```
##### Output to file

```
DEBUG 27-9-2020; 18:10:5 >> My first log write!
CRITICAL 27-9-2020; 18:10:5 >> CRITICAL MESSAGE! WATCH OUT &errors.errorString{s:"New error Check thi
CRITICAL 27-9-2020; 18:10:5 >> s out"}
CRITICAL 27-9-2020; 18:10:5 >> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmo
CRITICAL 27-9-2020; 18:10:5 >> d tempor incididunt ut labore et dolore magna aliqua.
CRITICAL 27-9-2020; 18:10:5 >> Urna et pharetra pharetra massa massa. Aenean euismod elementum nisi q
CRITICAL 27-9-2020; 18:10:5 >> uis eleifend. Purus in mollis nunc sed id semper risus.
```

##### Rotating files - number of files wont exceed specified value, oldest will be deleted and new file created
```
```

