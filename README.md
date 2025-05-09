# gdrive2

libary that assists creating access the google gdrive api.

## types

### credTyp  

type credTyp struct {
    ClientId string `yaml:"client_id"`
    ClientSecret string `yaml:"client_secret"`
    Scopes []string `yaml:"scopes"`
}

#### example file  

---
client_id: abc
client_secret: secret
scopes:
  - scopeA
  - scopeB


## functions

### ReadCredFile

ReadCredFile(filnam string)(cred *credTyp, err error)

### ReadCred

ReadCred(inB []byte)(cred *credTyp, err error)
