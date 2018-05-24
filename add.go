package IMA

import (
        "html/template"
        "net/http"
		"strconv"
        "appengine"
        "appengine/datastore"
)


func init() {
        http.HandleFunc("/add.html", root_add)
        http.HandleFunc("/sign_add", sign_add)
	
}


func assignmentKey_add(c appengine.Context) *datastore.Key {
        return datastore.NewKey(c, "IMA", "default_assignment", 0, nil)
}

func root_add(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("IMA").Ancestor(assignmentKey_add(c)).Order("-Date").Limit(10)
        greetings := make([]Assignment, 0, 10)
        if _, err := q.GetAll(c, &greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        if err := assignmentTemplate_add.Execute(w, greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}
	
var assignmentTemplate_add = template.Must(template.New("book").Parse(`
<html>
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
    width: 100%;
    padding: 12px 20px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-sizing: border-box;
}

input[type=submit] {
    width: 100%;
    background-color: orange;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
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

  <head>
  
    <title>Add</title>
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
    <form action="/sign_add" method="post">
    <label for="fname">Title</label>
    <input type="text" name="txtTitle" placeholder="Enter game title...">

    <label for="lname">Developer</label>
    <input type="text" name="txtDeveloper" placeholder="Enter developer name...">

    <label for="lname">Year</label>
    <input type="text" name="txtYear" placeholder="Enter release year...">
  
    <input type="submit" value="Analyse">
  </form>
</div>
	
  </body>
</html>
`))

func sign_add(w http.ResponseWriter, r *http.Request) {
         
						title = r.FormValue("txtTitle")
						developer = r.FormValue("txtDeveloper")
						year,_ = strconv.Atoi(r.FormValue("txtYear"))
	
	
		        http.Redirect(w, r, "/basic.html", http.StatusFound)

}
