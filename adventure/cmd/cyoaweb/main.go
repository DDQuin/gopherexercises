package main

import (
	"flag"
	"fmt"
	cyoa "gopher/adventure"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "The port for cyoa web")
	filename := flag.String("file", "gopher.json", "Json file")
	flag.Parse()
	fmt.Printf("Using the stroy in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(pathFn),
	)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {

	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"	
	}
	return path[len("/story/"):]

}

var storyTmpl = 	
`<!DOCTYPE html>
<html>
	<head>
		<title>choose your own adventure</title>
		<meta charset="utf-8">
		
	</head>
	<body>
	<style>
	body {
	font-family: helvetica, arial;
	}
	h1 {
	text-align:center;
	position:relative;
	}
	.page {
	width: 80%;
	max-width: 500px;
	margin: auto;
	margin-top: 40px;
	margin-bottom: 40px;
	padding: 80px;
	background: #FFFCF6;
	border: 1px solid #eee;
	box-shadow: 0 10px 6px -6px #777;
	}
	ul {
	border-top: 1px dotted #ccc;
	padding: 10px 0 0 0;
	-webkit-padding-start: 0;
	}
	li {
	padding-top: 10px;
	}
	a,
	a:visited {
	text-decoration: none;
	color: #6295b5;
	}
	a:active,
	a:hover {
	color: #7792a2;
	}
	p {
	text-indent: 1em;
	}
	</style>
	<section class="page">
	<h1>{{.Title}}</h1>
	{{range .Paragraphs}}
	<p>{{.}}</p>
	{{end}}

		<ul>
			{{range .Options}}
			<li>
				<a href="/story/{{.Chapter}}">{{.Text}}</a>
			</li>
			{{end}}
		</ul>	
	</section>
	</body>
</html>`


