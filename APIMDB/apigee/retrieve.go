package apigee

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  //"encoding/json"
  //"sort"
  "encoding/json"
  "sort"
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

func GetApigeeToken() []byte {
	url := "https://login.apigee.com/oauth/token"
	method := "POST"
	payload := strings.NewReader("grant_type=password&username=chayakorn.a%40kbtg.tech&password=Apigee%401234")
  
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

  func GetBillpaymentKvm(token string) []byte {
	url := "https://api.enterprise.apigee.com/v1/organizations/kbank-nonprod/environments/uat/keyvaluemaps/billpayment_partner_url"
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

func QueryKvm() []kvm {
	apigee1 := apigee_token{}
	jsonErr := json.Unmarshal(GetApigeeToken(), &apigee1)
	if jsonErr != nil {
		fmt.Println(jsonErr)
  }
  
  fmt.Println("token :", apigee1.Access_token)

  // part 2 
  people1 := apimKvm{}
	err := json.Unmarshal(GetBillpaymentKvm(apigee1.Access_token), &people1)
	if err != nil {
		panic(err)
	}

  fmt.Println("Going sort")
  people := people1.Entry
  sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
  
  return people
  /*for _, k := range people {
		fmt.Println(k.Name,"\t", k.Value)
	}*/
}