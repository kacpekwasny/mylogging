package mylogging

import (
	"errors"
	"fmt"
	cmt "imports/commontools"
	"imports/pyfuncs"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func (l *Logger) log(lines []string) error {
	// "\n" is added when writing to file
	// in function listenLoop
	for _, line := range lines {
		l.LinesCh <- line
	}
	return nil
}

// Start d
func (l *Logger) start(buff int) {
	if l.WrToFile || l.OutToCon {
		l.LinesCh = make(chan string, buff)
		if l.WrToFile {
			if l.RotFileHand {
				//
				// tworzenie nowych i usuwanie najstarszych plikow
				// jezeli rotating files, liczenie jaka maja wielkosc co
				// pewna ilosc wpisow zalezna od maksymalnej wilekosci
				// info, err := os.Stat("//")
				// info.Size()
				// sprawdzenie wielkosci pliku, a potem doawanie do zmiennej
				// wielkosci linijki
				lgFiles, err := findLogFiles(l.DirPath)
				cmt.A(err)
				l.setFileObRot(lgFiles)

			} else {
				l.setFileOb()
			}
			// jezeli jeden plik
			// jezeli kilka
			//fmt.Println("wr to file")
		}
	} else {
		panic(errors.New("logger not started as there is no output set"))
	}
}

// newLogFile d
func (l *Logger) newLogFile() {
	name := l.makeFullFname(l.makeFileName())
	fmt.Println(name)
	l.FilePath = l.DirPath + name
	l.setFileOb()
}

// listenLoop main loop for recieving creating and writing logs
func (l *Logger) listenLoop(buff int) {
	l.start(buff)
	go func() {
		defer close(l.LinesCh)
		for {
			line, err := <-l.LinesCh
			if !err {
				fmt.Println("d Logger.LinesCh is closed, end of Logger.Listen()")
				return
			}
			if l.OutToCon {
				fmt.Println(line)
			}
			if l.WrToFile {
				line = line + "\n"
				if !l.RotFileHand {
					// if write to single file
					l.handleFwr(line)
				} else {
					// if write to many files
					l.handleRot(line)
				}
			}
		}
	}()
}

// setFileOb d
func (l *Logger) setFileOb() {
	f, err := os.OpenFile(l.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, l.FilePerm)
	if err == nil {
		l.File = f
		info, err := os.Lstat(l.FilePath)
		cmt.A(err)
		l.FileSize = info.Size()
	}
	cmt.A(err)

}

// setFileObRot d
func (l *Logger) setFileObRot(lgFiles []os.FileInfo) error {
	// Set FilePath, FileSize, Find Last file, Create new...
	if lgFiles == nil || len(lgFiles) == 0 {
		l.newLogFile()
		return nil
	} else {
		f := newestLogFile(lgFiles)
		if f != nil && f.Size() < l.FileSizeLimit {
			l.FilePath = l.DirPath + f.Name()
			l.setFileOb()
		} else {
			l.newLogFile()
		}
		return nil
	}
}

// handleRot is for wr to file if rotating handler
func (l *Logger) handleRot(line string) {
	// CHeck size of file
	// if to big create new file
	// if to many files, delete oldest
	// name file by FileNameFormat

	// Checking the size and saving the file if size is to big
	// then checking if actualy the file is 2 big,
	// because maybe the string takes more memory in RAM
	// than on HDD/SSD
	n, err := l.File.WriteString(line)
	cmt.A(err)
	l.FileSize = l.FileSize + int64(n)
	if l.FileSize > l.FileSizeLimit {
		l.File.Close()

		f, err := os.Lstat(l.FilePath)
		cmt.A(err)
		if f.Size() < l.FileSizeLimit {
			l.setFileOb()
		} else {
			l.newLogFile()
		}
		// Find log files and delete the oldest one if too many log files
		files, err := findLogFiles(l.DirPath)
		cmt.A(err)
		if len(files) > l.MaxFiles {
			f = oldestLogFile(files)
			err = os.Remove(l.DirPath + f.Name())
			cmt.A(err)
		}
	}
}

// handleFwr HandleFileWrite
func (l *Logger) handleFwr(line string) {
	// check size
	n, err := l.File.WriteString(line)
	cmt.A(err)
	l.FileSize = l.FileSize + int64(n)
	if l.FileSize > l.FileSizeLimit {
		l.File.Close()
		i, err := os.Lstat(l.FilePath)
		cmt.A(err)
		if i.Size() < l.FileSizeLimit {
			l.setFileOb()
		} else {
			// Change name of file
			_, err := os.Stat(l.FilePath + ".old.log")
			if err != nil {
				os.Remove(l.FilePath + ".old.log")
			}
			os.Rename(l.FilePath, l.FilePath+".old.log")
			l.setFileOb()
		}
	}
}

// makeFileName s
func (l *Logger) makeFileName() string {
	var (
		parts []string
	)

	for _, str := range pyfuncs.Split(l.FileNameFormat, "%") {

		if str == "" {
			continue
		}

		// special strings to change for value
		_, b := cmt.InSlice(str, []string{"D", "M", "Y", "h", "m", "s", "Mon"})

		// if str is a format string
		if b {
			if str == "D" {
				parts = append(parts, stri(time.Now().Day()))
			} else if str == "M" {
				parts = append(parts, stri(int(time.Now().Month())))
			} else if str == "Mon" {
				parts = append(parts, stri(time.Now().Month()))
			} else if str == "Y" {
				parts = append(parts, stri(time.Now().Year()))
			} else if str == "h" {
				parts = append(parts, stri(time.Now().Hour()))
			} else if str == "m" {
				parts = append(parts, stri(time.Now().Minute()))
			} else if str == "s" {
				parts = append(parts, stri(time.Now().Second()))
			}
		} else {
			parts = append(parts, str)
		}
	}
	return pyfuncs.Join("", parts)
}

// makeFullFname d
func (l *Logger) makeFullFname(name string) string {
	files, err := findLogFiles(l.DirPath)
	cmt.A(err)
	file := newestLogFile(files)
	var (
		endfix = "."
		nr     = 1
	)
	if file != nil {
		parts := pyfuncs.Split(file.Name(), ".")
		if parts[len(parts)-1] == "log" {
			nr, err = strconv.Atoi(parts[len(parts)-2])
			if err != nil {
				nr = 1
			}
			endfix = name + endfix + stri(nr+1) + ".log"
			return endfix
		}
		return name + ".1.log"

	}
	return name + ".1.log"
}

func newestLogFile(logFiles []os.FileInfo) os.FileInfo {
	if len(logFiles) == 0 {
		return nil
	}
	// Get the youngest file
	var yngsFile os.FileInfo = logFiles[0]
	var f os.FileInfo
	for _, f = range logFiles {
		// s
		if f.ModTime().After(yngsFile.ModTime()) {
			yngsFile = f
		}
	}
	return yngsFile

}

func oldestLogFile(logFiles []os.FileInfo) os.FileInfo {
	if len(logFiles) == 0 {
		return nil
	} else if len(logFiles) == 1 {
		return logFiles[0]
	} else {
		var oldstFile os.FileInfo = logFiles[0]
		var f os.FileInfo
		for _, f = range logFiles {
			if f.ModTime().Before(oldstFile.ModTime()) {
				oldstFile = f
			}
		}
		return oldstFile
	}
}

///////////////////////////////////////////////

// findLogFiles d
func findLogFiles(dirpath string) ([]os.FileInfo, error) {
	var filesInfo []os.FileInfo
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		cmt.A(err)
		if !info.IsDir() && cmt.InStr(info.Name(), ".log") {
			parts := pyfuncs.Split(info.Name(), ".")
			_, err := strconv.Atoi(parts[len(parts)-2])
			if err == nil {
				filesInfo = append(filesInfo, info)
			}
		}
		return err
	})
	return filesInfo, err
}

func stri(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
