package main

import (
  "fmt"
  //"strings"
  //"net/http"
  //"io/ioutil"
  "encoding/json"
  //"apigee/retrieve"
  "sort"
  "APIM/apigee"
)
type apigee_token struct {
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

func main() {
  apigee1 := apigee_token{}
	jsonErr := json.Unmarshal(apigee.GetApigeeToken(), &apigee1)
	if jsonErr != nil {
		fmt.Println(jsonErr)
  }
  
  fmt.Println("token :", apigee1.Access_token)

  // part 2 
  people1 := apimKvm{}
	err := json.Unmarshal(apigee.GetBillpaymentKvm(apigee1.Access_token), &people1)
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