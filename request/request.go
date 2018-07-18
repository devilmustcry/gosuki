package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get : HTTP Get Request
func Get() {
	url := "https://swapi.co/api/people/1/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var body struct {
		Name      string
		Height    string
		Mass      string
		HairColor string `json:"hair_color"`
		SkinColor string `json:"skin_color"`
		EyeColor  string `json:"eye_color"`
	}
	blob, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(blob, &body)

	fmt.Println(body.Name)
	fmt.Println(body.Height)
	fmt.Println(body.Mass)
	fmt.Println(body.HairColor)
	fmt.Println(body.SkinColor)
	fmt.Println(body.HairColor)
	blob2, _ := json.Marshal(&body)
	fmt.Println(string(blob2))
}
