package main

import f "fmt"

func main(){
	websites := map[string]string{
		"Google": "google.com",
		"Amazon": "amazon.com",
	}


	f.Println(websites)
	f.Println(websites["Amazon"])
}