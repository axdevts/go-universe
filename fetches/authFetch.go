package fetches

import (
	"encoding/json"
	"fmt"
	fetch_models "github.com/axdevts/go-universe/fetches/models"
	"io/ioutil"
	"net/http"
)

func LoadData() {
	fmt.Println("Load data called.")
	SayHello()
}

func FetchData(ch chan string) {
	url := "http://localhost:8000/api/posts"
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk2Njk1OTQsInN1YiI6MX0.Na_DwqVE8abFTncHAqmYxTKnhl5fw-VofgsB--bG1OI"

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Authorization header with the Bearer token
	// req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Cookie", "Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk2NzIxMTgsInN1YiI6MX0.2RQGB1PxecKpB17LheCRKzxqymrdKKNlSk8QpDV4CoM")

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var resData fetch_models.PostResponse

	errRes := json.Unmarshal(body, &resData)

	if errRes != nil {
		fmt.Println("Error unmarshaling JSON:", errRes)
		return
	}

	fmt.Println("Response Data in JSON:", resData)

	jsonData, err := json.MarshalIndent(resData.Response.Data[0], "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Printf("%s\n", jsonData)

	ch <- string(body)

}
