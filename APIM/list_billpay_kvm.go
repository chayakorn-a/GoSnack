package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "sort"
)
type apigee struct {
    Access_token string `json:"access_token"`
}

type apimKvm struct {
	Number string `json:"name"`
	Entry []kvm `json:"entry"`
}

type kvm struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

func retrieveData() []byte {
  url := "https://kasikorn-bank.login.apigee.com/oauth/token"
  method := "POST"
  payload := strings.NewReader("grant_type=password&username=machineuserkbtg%40kbtg.tech&password=P%40ssw0rd")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Authorization", "Basic ZWRnZWNsaTplZGdlY2xpc2VjcmV0")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
  }
  return body
}

func retrieveDataKvm(token string) []byte {
  url := "https://api.enterprise.apigee.com/v1/organizations/kbank-prod/environments/prod/keyvaluemaps/billpayment_partner_url"
  method := "GET"

  payload := strings.NewReader(``)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Authorization", "Bearer "+token)

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
  }
  return body
}

func main() {
  apigee1 := apigee{}
	jsonErr := json.Unmarshal(retrieveData(), &apigee1)
	if jsonErr != nil {
		fmt.Println(jsonErr)
  }
  
  fmt.Println("token :", apigee1.Access_token)

  // part 2 
  people1 := apimKvm{}
	err := json.Unmarshal(retrieveDataKvm(apigee1.Access_token), &people1)
	if err != nil {
		fmt.Println(err)
		return
	}

  fmt.Println("Going sort")
  people := people1.Entry
  sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
  
  for _, k := range people {
		fmt.Println(k.Name,"\t", k.Value)
	}
}