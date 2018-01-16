package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		timer := time.NewTimer(time.Second * 15)
		<-timer.C
		println("Таймер истек")

		responseP, err := http.Get("https://api.vk.com/method/wall.get?owner_id=-41774259&count=10&offset=0&access_token=13d3de0213d3de0213d3de029b13b3c292113d313d3de0249e43129fa9241a0bc1fa3e5&v=5.52")
		if err != nil {
			panic(err)
		}

		defer responseP.Body.Close()
		contentsP, err := ioutil.ReadAll(responseP.Body)
		if err != nil {
			panic(err)
		}

		type PresponseJson struct {
			Response struct {
				Count int `json:"count"`
				Items []struct {
					ID          int    `json:"id"`
					FromID      int    `json:"from_id"`
					OwnerID     int    `json:"owner_id"`
					Date        int    `json:"date"`
					MarkedASADS int    `json:"marked_as_ads"`
					PostType    string `json:"post_type"`
					Text        string `json:"text"`
				} `json:"items"`
			} `json:"response"`
		}

		var unmarshP = PresponseJson{}
		respP := json.Unmarshal(contentsP, &unmarshP)
		if respP != nil {
			panic(respP)
		}

		for _, rowC := range unmarshP.Response.Items {
			zapr := fmt.Sprintf("https://api.vk.com/method/wall.getComments?owner_id=-41774259&post_id=%v&count=4&access_token=13d3de0213d3de0213d3de029b13b3c292113d313d3de0249e43129fa9241a0bc1fa3e5&v=5.52",
				rowC.ID)                 //Запрос для комментариев
			fmt.Println(rowC.Text, "\n") //Текс постов
			//fmt.Println(zapr)
			responseC, err := http.Get(zapr)
			if err != nil {
				panic(err)
			}
			defer responseC.Body.Close()
			contentsC, err := ioutil.ReadAll(responseC.Body)
			if err != nil {
				panic(err)
			}
			type CresponseJson struct {
				Response struct {
					Count int `json:"count"`
					Items []struct {
						ID     int    `json:"id"`
						FromID int    `json:"from_id"`
						Date   int    `json:"date"`
						Text   string `json:"text"`
					} `json:"items"`
				} `json:"response"`
			}
			var unmarshC= CresponseJson{}
			respC := json.Unmarshal(contentsC, &unmarshC)
			if respC != nil {
				panic(respC)
			}
			for _, rowC := range unmarshC.Response.Items { //Текс комментариев
				fmt.Println(rowC.Text)
			}
		}
	}
}
