package main

import "net/http"

func main() {

}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	buf := make([]byte, 1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result = result + string(buf)
	}

	return
}
