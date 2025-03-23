package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

var dir, _ = os.Getwd()
var logDir = fmt.Sprintf("%s/logs", dir)
var pactDir = fmt.Sprintf("%s/pacts", getParentDirectory())
var Regex = matchers.Regex

func TestNewV2Pact(t *testing.T) {

	mockProvider, err := consumer.NewV2Pact(consumer.MockHTTPProviderConfig{
		Consumer: "SampleBookConsumer",
		Provider: "SampleBookProvider",
		LogDir:   logDir,
		PactDir:  pactDir,
	})
	assert.NoError(t, err)
	fmt.Printf("log dir: %s\n", logDir)
	fmt.Printf("pact dir: %s\n", pactDir)

	err = mockProvider.
		AddInteraction().
		Given("A book exists in system").
		UponReceiving("A request for book").
		WithRequest("GET", "/BookStore/v1/Book/ISBN/9781449331818").
		WillRespondWith(200, func(b *consumer.V2ResponseBuilder) {
			b.Header("Content-Type", Regex("application/json", "application\\/json"))
			b.BodyMatch(&Book{
				ISBN:        "9781449331818",
				Title:       "Learning JavaScript Design Patterns",
				SubTitle:    "A JavaScript and jQuery Developer's Guide",
				Author:      "Addy Osmani",
				PublishDate: "2020-06-04T09:11:40.000Z",
				Publisher:   "O'Reilly Media",
				Pages:       254,
				Description: "Description",
				Website:     "hihi",
			})
		}).
		ExecuteTest(t, func(config consumer.MockServerConfig) error {

			fmt.Printf("Pact Mock Server Running at: %s:%d\n", config.Host, config.Port)
			// Act: test our API client behaves correctly
			// Initialise the API client and point it at the Pact mock server
			client := bookClient(config.Host, config.Port)

			// Execute the API client
			book, err := client.GetBook("9781449331818")

			fmt.Printf("Received Book: %+v\n", book) // Debugging
			// Assert: check the result
			assert.NoError(t, err)
			assert.Equal(t, "9781449331818", book.ISBN)

			return err
		})

	assert.NoError(t, err)
}

type Book struct {
	ISBN        string `json:"isbn" pact:"example=9781449331818"`
	Title       string `json:"title" pact:"example=Learning JavaScript Design Patterns"`
	SubTitle    string `json:"subTitle" pact:"example=A JavaScript and jQuery Developer's Guide"`
	Author      string `json:"author" pact:"example=Addy Osmani"`
	PublishDate string `json:"publish_date" pact:"example=2020-06-04T09:11:40.000Z"`
	Publisher   string `json:"publisher" pact:"example=O'Reilly Media"`
	Pages       int    `json:"pages" pact:"example=254"`
	Description string `json:"description" pact:"example=With Learning JavaScript Design Patterns"`
	Website     string `json:"website" pact:"example=http://www.addyosmani.com.br"`
}

// Book API Client to test
type bookAPIClient struct {
	port int
	host string
}

func getParentDirectory() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parentDir := filepath.Dir(currentDir)
	return parentDir
}

func bookClient(host string, port int) *bookAPIClient {
	return &bookAPIClient{
		host: host,
		port: port,
	}
}

func (u *bookAPIClient) GetBook(isbn string) (*Book, error) {

	path := fmt.Sprintf("/BookStore/v1/Book/ISBN/%s", isbn)
	resp, err := http.Get(fmt.Sprintf("http://%s:%d%s", u.host, u.port, path))

	if err != nil {
		return nil, err
	}
	fmt.Println("Response Status Code:", resp.StatusCode)
	defer resp.Body.Close()

	// Debug: Print raw API response
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Raw Response:", string(respBody))

	// Reset body for decoding
	resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

	book := new(Book)
	err = json.NewDecoder(resp.Body).Decode(book)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Decoded Book: %+v\n", book)
	return book, nil
}
