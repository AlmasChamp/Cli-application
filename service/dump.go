package joker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func dump(n int) {

	wg := new(sync.WaitGroup)
	const categoryUrl = "https://api.chucknorris.io/jokes/random?category="
	category := []string{}

	resp, err := http.Get("https://api.chucknorris.io/jokes/categories")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &category)
	if err != nil {
		errors.New("Something went wrong")
	}

	for i := 0; i < len(category); i++ {
		wg.Add(1)

		file, err := os.OpenFile(category[i]+".txt", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		go addJokes(n, categoryUrl, category[i], wg)

	}

	wg.Wait()
	fmt.Println("End")

}

func addJokes(n int, categoryUrl string, category string, wg *sync.WaitGroup) {

	type Categories struct {
		Id    string `json:"id"`
		Value string `json:"value"`
	}
	categories := &Categories{}
	uniqJokes := make(map[string]bool)

	for i := 0; i < n; i++ {
		resp, err := http.Get(categoryUrl + category)
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		error := json.Unmarshal(body, categories)
		if error != nil {
			errors.New("Something went wrong")
		}

		if uniqJokes[categories.Id] == false {
			uniqJokes[categories.Id] = true
			fmt.Println(categories.Id)
			file, err := os.OpenFile(category+".txt", os.O_RDWR, 0666)
			if err != nil {
				log.Fatal(err)
			}

			a, err := file.Stat()
			if err != nil {
				log.Println(err)
			}
			file.WriteAt([]byte(categories.Value+"\n"), a.Size())
		}
	}
	wg.Done()
}
