package joker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func random() {
	chack := &Chack{}

	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	error := json.Unmarshal(body, chack)
	if error != nil {
		errors.New("Something went wrong")
	}
	fmt.Println(chack.Value)
}
