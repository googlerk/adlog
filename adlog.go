package adlog

import (
  "flag"
  "os"
  "log"
)

var (
  Log      *log.Logger
)


func init() {
    // set location of log file
    var logpath = "./logs/adlogs.log"

   flag.Parse()
   f, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
      log.Println(err)
    }
    defer f.Close()

   if err1 != nil {
      panic(err1)
   }
      Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
      Log.Println("LogFile : " + logpath)
}
