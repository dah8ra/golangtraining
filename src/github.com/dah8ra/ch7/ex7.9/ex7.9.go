package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
	"os"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type List struct {
	Tracks []*Track
}

var t = template.Must(template.New("").Parse(`
<h1>Track list</h1>
<table>
<tr style='text-align: left'>
  <th>Title</th>
  <th>Artist</th>
  <th>Album</th>
  <th>Year</th>
  <th>Length</th>
</tr>
{{range .Tracks}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.RawQuery
		fmt.Fprintf(os.Stdout, query)
		printTracksForHTML(query, w)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type HistorySort struct {
	Clicked string
	Tracks  []*Track
	less    func(x, y *Track) bool
}

func (x HistorySort) Len() int           { return len(x.Tracks) }
func (x HistorySort) Less(i, j int) bool { return x.less(x.Tracks[i], x.Tracks[j]) }
func (x HistorySort) Swap(i, j int) {
	x.Tracks[i], x.Tracks[j] = x.Tracks[j], x.Tracks[i]
}

func printTracksForHTML(clicked string, w http.ResponseWriter) {

	tracks := List{}
	tracks.Tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	sort.Sort(HistorySort{clicked, tracks.Tracks, func(x, y *Track) bool {
		if clicked == "Title" {
			if x.Title != y.Title {
				return x.Title < y.Title
			}
		} else if clicked == "Artist" {
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
		} else if clicked == "Album" {
			if x.Album != y.Album {
				return x.Album < y.Album
			}
		} else if clicked == "Year" {
			if x.Year != y.Year {
				return x.Year < y.Year
			}
		} else if clicked == "Length" {
			if x.Length != y.Length {
				return x.Length < y.Length
			}
		}
		return false
	}})

	if err := t.Execute(w, &tracks); err != nil {
		log.Fatal(err)
	}
}
