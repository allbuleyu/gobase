package ch7

import (
	"time"
	"html/template"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

var Tracks = []*Track{
	{Title:"go", Artist:"dddd", Album:"fffff", Year:2012, Length:time.Second*3 + time.Minute * 3},
	{Title:"go2", Artist:"dddd2", Album:"fffff2", Year:2013, Length:time.Second*4 + time.Minute * 4},
	{Title:"go3", Artist:"dddd3", Album:"fffff3", Year:2014, Length:time.Second*5 + time.Minute * 5},
}
type MyStruct struct {
	Items []*Track
}

func PrintTracks(trackss []*Track) *template.Template {
	xxx := template.Must(template.New("xxx").Parse(table))

	//s := MyStruct{tracks}
	//xxx.Execute(os.Stdout, tracks)
	return xxx
}

var table string = `
<table>
    <tbody>
		<tr>
            <td onclick="xxx()">Title</td>
            <td>Artist</td>
            <td>Album</td>
            <td>Year</td>
            <td>Length</td>
        </tr>
        {{range .}}
		<tr>
            <td>{{.Title}}</td>
            <td>{{.Artist}}</td>
            <td>{{.Album}}</td>
            <td>{{.Year}}</td>
            <td>{{.Length}}</td>
        </tr>
		{{end}}
    </tbody>
</table>

<script>

function xxx() {
	location.href = location.href+"?sort=title";
}

</script>
`