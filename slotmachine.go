// Package to handle and route other web request as a server
package main

import (
	"fmt"
	//"github.com/skratchdot/open-golang/open"
	"math/rand"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	openURL("http://localhost:8888/")
	// Handlers
	http.HandleFunc("/", handler)
	// CSS
	//http.HandleFunc("/style.css", Default)
	http.ListenAndServe(":8888", nil)
	//open.Run("http://localhost:8888")
}

func handler(w http.ResponseWriter, r *http.Request) {

	this_pull := randSeq(10)
	fmt.Fprintf(w, "<html><head></head><body>")
	fmt.Fprintf(w, "<h1><code>")
	fmt.Fprintf(w, this_pull)
	fmt.Fprintf(w, "</code></h2>")
	fmt.Fprintf(w, `<p><button onclick="myFunction()">Pull Lever</button></p> <script> function myFunction() { location.reload(); } </script>`)
	fmt.Fprintln(w, "<code>")

	var m map[string]int
	m = make(map[string]int)

	for _, y := range this_pull {
		m[string(y)] = m[string(y)] + 1
		//	return string(m)
	}

	fmt.Fprintf(w, "<h2><code>")
	for x, y := range m {
		if y == 2 {
			fmt.Fprintln(w, "You got a Pair of: ", x, "'s<br>")
			fmt.Fprintln(w, `<p><iframe src="//giphy.com/embed/nXxOjZrbnbRxS" width="100" frameBorder="0" style="max-width: 100%" class="giphy-embed" webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe></p>`)
		} else if y == 3 {
			fmt.Fprintln(w, "You got a Trio of: ", x, "'s<br>")
			fmt.Fprintln(w, `<p><iframe src="//giphy.com/embed/sIIhZliB2McAo" width="200" frameBorder="0" style="max-width: 100%" class="giphy-embed" webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe></p>`)
		} else if y == 4 {
			fmt.Fprintln(w, "You got a Quartet of: ", x, "'s<br>")
			fmt.Fprintln(w, `<p><iframe src="//giphy.com/embed/ujUdrdpX7Ok5W" width="200" frameBorder="0" style="max-width: 100%" class="giphy-embed" webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe></p>`)
		} else if y > 4 {
			fmt.Fprintln(w, "You got a Yahtzee of: ", x, "'s<br>")
			fmt.Fprintln(w, `<p><iframe src="//giphy.com/embed/XUFPGrX5Zis6Y" width="480" height="269" frameBorder="0" style="max-width: 100%" class="giphy-embed" webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe></p>`)
		}
	}
	fmt.Fprintf(w, "</h2></code>")
	fmt.Fprintf(w, "</body></html>")
}

var letters = []rune("1234567890ABCDEF")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:4001/").Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}
