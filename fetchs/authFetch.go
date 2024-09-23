package fetchs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoadData() {
	fmt.Println("Load data called.")
	SayHello()
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"`
	Posts    []Post
}

type Post struct {
	CategoryID uint   `gorm:"foreignkey:CategoryID" json:"categoryID"`
	Title      string `gorm:"not null" json:"title"`
	Body       string `gorm:"type:text" json:"body"`
	UserID     uint   `gorm:"foreignkey:UserID" json:"userID"`
	// Category   uint   `gorm:"foreignkey:CategoryID"`
	User User `gorm:"foreignkey:UserID"`
	// Comments   []Comment
}

type PostData struct {
	Data        []Post `json:"data"`
	CurrentPage uint   `json:"current_page"`
	From        uint   `json:"from"`
	To          uint   `json:"to"`
	LastPage    uint   `json:"last_page"`
	PerPage     uint   `json:"per_page"`
	Total       uint   `json:"total"`
}

type PostResponse struct {
	Response PostData `json:"response"`
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

	var resData PostResponse

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

	// fmt.Println("Len of Posts : ", len(resData.Response.Data))

	// fmt.Println("First Post:resData.Response.Data[0].UserID : ", resData.Response.Data[0].User.Name)

	// fmt.Println(string(body))
	ch <- string(body)

	// resp, err := http.Get("http://localhost:8000/api/posts")
	// if err != nil {
	// 	ch <- fmt.Sprintf("Error: %v", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// ch <- string(body)
}
