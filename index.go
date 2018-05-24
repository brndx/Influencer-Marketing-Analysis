package IMA

import (
        "html/template"
        "net/http"
        "time"
		"strconv"
        "appengine"
        "appengine/datastore"
        "appengine/user"
)

type Assignment struct {
		Title string
        Developer  string
		Year	int
		Negative	int
		Positive	int
		Neutral		int
		Score	int
        Date    time.Time
}

var title string = "";
var developer string = "";
var year int = 0;
func init() {
        http.HandleFunc("/index.html", root)
        http.HandleFunc("/sign", sign)
	
}


func assignmentKey(c appengine.Context) *datastore.Key {
        return datastore.NewKey(c, "IMA", "default_assignment", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("IMA").Ancestor(assignmentKey(c)).Order("-Date").Limit(10)
        greetings := make([]Assignment, 0, 10)
        if _, err := q.GetAll(c, &greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        if err := assignmentTemplate.Execute(w, greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}
	 func main() {
		
	 }
var assignmentTemplate = template.Must(template.New("book").Parse(`
<html>
<head>
<style>
body {
background-image: url("https://image.ibb.co/dT5a66/image.png");
}
label, option, form {
font-family: verdana;
font-size: 14px;
}
label {
color: LightGrey;
}

main {

padding-top: 80px;
  text-align: center;
  font-family: verdana;
  display: fixed;
    font-size: 30px;
color: LightGrey;
    cursor: pointer;

}

input[type=text], select {
    width: 30%;
    padding: 12px 20px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-sizing: border-box;
}

input[type=submit] {
    width: 30%;
    background-color: orange;
    color: white;
    padding: 14px 20px;
	margin: 4% 35%;
    position: inline-block;
    border: none;
    border-radius: 4px;
    cursor: pointer;
	font-family: verdana;
font-size: 16px;
}

input[type=submit]:hover {
    background-color: #ee7600;
}

div {
background-image: url("https://image.ibb.co/dT5a66/image.png");

    border-radius: 5px;
    background-color: #f2f2f2;
    padding: 20px;
}
</style>
</head>
  <head>
  
    <title>Home</title>
  </head>
  <body>
<script>
function home() {
window.location = "index.html";
}
</script>
<div>
<main onclick="home()">Influencer Marketing Analysis</main></div>
</div>

		  		     <div>
      <input onclick="myFunction()" type="submit" value="Analyse a new game">

</div>

		  		      <div>
        <input onclick="myFunction2()" type="submit" value="View / Update a previous analysis">

</div>
	<script>
function myFunction() {
    window.location = "add.html";
}
function myFunction2() {
    window.location = "select.html";
}
</script>
  </body>
</html>
`))

func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        g := Assignment{
				
                
        }
						g.Title = r.FormValue("textTitle")
						g.Developer = r.FormValue("textDeveloper")
						g.Year,_ = strconv.Atoi(r.FormValue("textYear"))
						
	
	g.Date = time.Now()
        if u := user.Current(c); u != nil {
                g.Developer = u.String()
        }
        key := datastore.NewIncompleteKey(c, "IMA", assignmentKey(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/", http.StatusFound)
}
