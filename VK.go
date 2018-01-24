package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"sort"
)
func main() {
	responseP, err := http.Get("https://api.vk.com/method/wall.get?owner_id=-41774259&count=5&offset=0&access_token=13d3de0213d3de0213d3de029b13b3c292113d313d3de0249e43129fa9241a0bc1fa3e5&v=5.52")
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
	var unmarshP= PresponseJson{}
	respP := json.Unmarshal(contentsP, &unmarshP)
	if respP != nil {
		panic(respP)
	}
	dataPost := []int{}
	dataComments := []int{}
	for _, rowC := range unmarshP.Response.Items {
		zapr := fmt.Sprintf("https://api.vk.com/method/wall.getComments?owner_id=-41774259&post_id=%v&count=4&sort=desc&access_token=13d3de0213d3de0213d3de029b13b3c292113d313d3de0249e43129fa9241a0bc1fa3e5&v=5.52",
			rowC.ID) //Запрос для комментариев
		dataPost = append(dataPost, rowC.Date)
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
		var unmarshC = CresponseJson{}
		respC := json.Unmarshal(contentsC, &unmarshC)
		if respC != nil {
			panic(respC)
		}
		for _, rowC := range unmarshC.Response.Items { //Текс комментариев
			fmt.Println(rowC.Text)
			dataComments = append(dataComments, rowC.Date)
		}
	}
	//fmt.Println(dataPost)
	for i := 1; i <= 5; i++ {
		timer := time.NewTimer(time.Second * 5)
		<-timer.C
		println("Таймер истек")

		OtvetP, err := http.Get("https://api.vk.com/method/wall.get?owner_id=-41774259&count=5&offset=0&access_token=13d3de0213d3de0213d3de029b13b3c292113d313d3de0249e43129fa9241a0bc1fa3e5&v=5.52")
		if err != nil {
			panic(err)
		}
		defer OtvetP.Body.Close()
		SoderjimoeP, err := ioutil.ReadAll(OtvetP.Body)
		if err != nil {
			panic(err)
		}
		type PostStructJson struct {
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
		var UnmarshP= PostStructJson{}
		respP := json.Unmarshal(SoderjimoeP, &UnmarshP)
		if respP != nil {
			panic(respP)
		}
		sort.Ints(dataComments)
		sort.Ints(dataPost)
		for _, rt := range UnmarshP.Response.Items{
			if rt.Date > dataPost[len(dataPost)-1] {
				fmt.Println(rt.Text)
				dataPost = append(dataPost, rt.Date)
			}else {
				continue
			}
		}
		for _, row := range UnmarshP.Response.Items {
			zapros := fmt.Sprintf("https://api.vk.com/method/wall.getComments?owner_id=-41774259&post_id=%v&count=4&sort=desc&access_token=13d3de0213d3de0213d3de029b13b3c292113d313d3de0249e43129fa9241a0bc1fa3e5&v=5.52",
				row.ID) //Запрос для комментариев
			//fmt.Println(row.Text)
			OtvetC, err := http.Get(zapros)
			if err != nil {
				panic(err)
			}
			defer OtvetC.Body.Close()
			SoderjimoeC, err := ioutil.ReadAll(OtvetC.Body)
			if err != nil {
				panic(err)
			}
			type CommentsStructJson struct {
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
			var UnmarshC= CommentsStructJson{}
			respCom := json.Unmarshal(SoderjimoeC, &UnmarshC)
			if respCom != nil {
				panic(respCom)
				}
			for _, rowC := range UnmarshC.Response.Items { //Текс комментариев
				if rowC.Date > dataComments[len(dataComments)-1]{
					fmt.Println(rowC.Text)
					dataComments = append(dataComments, rowC.Date)
				}else {
					continue
				}
			}
		}
	}
}

