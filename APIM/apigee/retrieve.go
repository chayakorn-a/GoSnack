package apigee

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  //"encoding/json"
  //"sort"
)

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