package gapilib


import (
	"fmt"
	"os"

  	yaml "github.com/goccy/go-yaml"
)

type credTyp struct {
	ClientId string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Scopes []string `yaml:"scopes"`
}

func ReadCredFile(filnam string)(cred *credTyp, err error){
	var credObj credTyp

	inB, err := os.ReadFile(filnam)
	if err != nil {return nil, fmt.Errorf("Read File: %v", err)}

//	fmt.Printf("in:\n%s\n", string(inB))

	err = yaml.Unmarshal(inB, &credObj)
	if err != nil {return nil, fmt.Errorf("Unmarshal: %v",err)}

	return &credObj, nil
}

func ReadCred(inB []byte)(cred *credTyp, err error){

	var credObj credTyp

//	fmt.Printf("in:\n%s\n", string(inB))

	err = yaml.Unmarshal(inB, &credObj)
	if err != nil {return nil, err}

	return &credObj, nil
}

func PrintCred(cred *credTyp) {
	fmt.Println("************ credObj **************")
	fmt.Printf("ClientId:     %s\n", cred.ClientId)
	fmt.Printf("ClientSecret: %s\n", cred.ClientSecret)
	if len(cred.Scopes) >0 {
		fmt.Printf("Scopes:\n")
	} else {
		fmt.Printf("no Scopes\n")
	}
	for i:=0; i< len(cred.Scopes); i++  {
		fmt.Printf(" -%d: %s\n", i+1, cred.Scopes[i])
	}
	fmt.Println("********** end credObj ************")
}
