package config

import(
  "encoding/json"
  "os"
)

type WholeConfig struct {
  Dbusername string
  Dbpassword string
  Dbhostname string
  Dbschema   string
  Dedishost  string
  Serverbase string 
}

func Parse(path string) (ret *WholeConfig ,err error) {
  raw,err := os.ReadFile(path)
  if err != nil{
    return
  }
  err = json.Unmarshal(raw,&ret)
  return
}
