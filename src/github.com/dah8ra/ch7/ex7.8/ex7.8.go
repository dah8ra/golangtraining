package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	queue := [...]string{"Title", "Artist", "Album", "Year", "Length"}
	// for loop shows clicked history.
	// e.g. First loop is first click.
	count := 10000 // loop count
	showflag := false
	start := time.Now()
	for i := 0; i < count; i++ {
		f64 := float64(i)
		modulo := math.Mod(f64, 5)
		clicked := queue[int(modulo)]
		if showflag {
			fmt.Println("\nClicked " + clicked + ":")
		}
		sort.Sort(historySort{clicked, tracks, func(x, y *Track) bool {
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
		if showflag {
			printTracks(tracks)
		}
	}
	end := time.Now()
	history :=  (end.Sub(start)).Seconds()
	fmt.Printf("historySort -> %f seconds\n", history)
	
	startStable := time.Now()
	for i := 0; i < count; i++ {
		f64 := float64(i)
		modulo := math.Mod(f64, 5)
		clicked := queue[int(modulo)]
		if showflag {
			fmt.Println("\nClicked " + clicked + ":")
		}
		sort.Stable(historySort{clicked, tracks, func(x, y *Track) bool {
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
		if showflag {
			printTracks(tracks)
		}
	}
	endStable := time.Now()
	stable := (endStable.Sub(startStable)).Seconds()
	fmt.Printf("stableSort -> %f seconds\n", stable)
	fmt.Printf("diff -> %f seconds\n",  history - stable)
}

type historySort struct {
	c    string
	t    []*Track
	less func(x, y *Track) bool
}

func (x historySort) Len() int           { return len(x.t) }
func (x historySort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x historySort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}
