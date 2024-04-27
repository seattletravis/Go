package maps

import f "fmt"

func main(){
	websites := map[string]string{
		"Google": "google.com",
		"Amazon": "amazon.com",
	}


	f.Println(websites)
	f.Println(websites["Amazon"])
	websites["LinkedIn"] = "linkedin.com"
	f.Println(websites)

	delete(websites, "Google")
	f.Println(websites)
}