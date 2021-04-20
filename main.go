package main

import (
	"fmt"
	"io/ioutil"

	"github.com/donadonandre/go-httpclient/gohttp"
)

var (
	// Singleton here
	gitHubHttpClient = getGitHubClient()
)

func getGitHubClient() gohttp.HttpClient {
	client := gohttp.New()

	/*commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client.SetHeaders(commonHeaders)*/
	return client
}

func main() {
	getUrls()
}

func getUrls() {
	//headers := make(http.Header)
	// headers.Set("Authorization", "Bearer ABC-123")

	response, err := gitHubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
