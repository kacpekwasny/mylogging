# mylogging

### Easy start

#### Single log file output and no output to console

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
<br>
<br>
<br>

#### Rotating files - number of files wont exceed specified value, oldest will be deleted and new file created
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
		dirpath               = "/your/dir/path/for/log/files"
		maxfiles              = 20
		maxfileSize     int64 = 1024 * 4
		outputToConsole       = true
	)

	lgr.RotatingFiles(outputToConsole, dirpath, maxfiles, maxfileSize)

	// Up to 20 strings will be in the buffer before your Code will wait
	// for function to save and output
	lgr.SetBuff(20)

	// lgr.Debug wont work because level is to high (Debug works only for level<=10)
	// so only Info Warning Error and Critical will work
	lgr.SetLevel(12)

	lgr.Listen()

	lgr.Debug("My first log write!")
	lgr.Critical("CRITICAL MESSAGE! WATCH OUT", fmt.Sprintf("%#v", errors.New("New error Check this out")))
	lgr.Critical(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Urna et pharetra pharetra massa massa. Aenean euismod elementum nisi quis eleifend. Purus in mollis nunc sed id semper risus.`)

	lgr.StopListen()
}
```
##### Output in console
```
CRITICAL 27-9-2020; 18:36:43 >> CRITICAL MESSAGE! WATCH OUT &errors.errorString{s:"New error Check thi
CRITICAL 27-9-2020; 18:36:43 >> s out"}
CRITICAL 27-9-2020; 18:36:43 >> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmo
CRITICAL 27-9-2020; 18:36:43 >> d tempor incididunt ut labore et dolore magna aliqua.
CRITICAL 27-9-2020; 18:36:43 >> Urna et pharetra pharetra massa massa. Aenean euismod elementum nisi q
CRITICAL 27-9-2020; 18:36:43 >> uis eleifend. Purus in mollis nunc sed id semper risus.
```

##### Output files

![Files screen shot](https://github.com/kacpekwasny/mylogging/blob/master/readmefiles/scrns.png)


<br>
<br>
<br>

#### Other functions

```
lgr.SetFormatStr("%D%-%M% and hour %H% Level of debugging: %L% >message will be at the end>")
lgr.SetMsgLengthLim(70)

//   70 == len("CRITICAL MESSAGE! WATCH OUT &errors.errorString{s:"New error Check thi")
//   prefix isn't counted

// Output
27-9 and hour H Level of debugging: CRITICAL >message will be at the end>CRITICAL MESSAGE! WATCH OUT &errors.errorString{s:"New error Check thi
27-9 and hour H Level of debugging: CRITICAL >message will be at the end>s out"}
27-9 and hour H Level of debugging: CRITICAL >message will be at the end>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmo
27-9 and hour H Level of debugging: CRITICAL >message will be at the end>d tempor incididunt ut labore et dolore magna aliqua.

lgr.SetFileNameFormat("mylogging_%D%-%M%_%h%")
// FileName
mylogging_27-9_19.1.log

lgr.SetFilePerm(0644)
// Set log file permissions
```
