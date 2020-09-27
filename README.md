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
	lgr := mylogging.LoggerInit()
	
  var (
    filepath string = "/path/to/your/log/file"
    fileSizeLimit int64 = 1024*1024*5
    outputToConsole bool = false
  )
  
  lgr.WrToSingleFile(outputToConsole, filepath, fileSizeLimit)
  
  // Open channel, Open file
	lgr.Listen()

	lgr.Debug("My first log write!")
	lgr.Critical("CRITICAL MESSAGE! WATCH OUT", fmt.Sprintf("%#v", errors.New("New error Check this out")))
	lgr.Critical("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Urna et pharetra 
            pharetra massa massa. Aenean euismod elementum nisi quis eleifend. Purus in mollis nunc sed id semper risus.")

  // Write pending lines, Close channel, Save file,
	lgr.StopListen()
}


```
