package IMA

import (
        "html/template"
        "net/http"
        "appengine"
		"strconv"
        "appengine/datastore"
)



func init() {
        http.HandleFunc("/select.html", root_select)
        http.HandleFunc("/sign_select", sign_select)
	
}

func assignmentKey_select(c appengine.Context) *datastore.Key {
        return datastore.NewKey(c, "IMA", "default_assignment", 0, nil)
}

func root_select(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("IMA").Ancestor(assignmentKey_select(c)).Order("-Date").Limit(10)
        greetings := make([]Assignment, 0, 10)
        if _, err := q.GetAll(c, &greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        if err := assignmentTemplate_select.Execute(w, greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}
	
var assignmentTemplate_select = template.Must(template.New("book").Parse(`
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
select.hidden {

visibility: hidden;
 width: 0%;
    padding: 0px 0px;
    margin: 0px 0;
    display: inline-block;
    border: 0px solid #ccc;
    border-radius: 0px;
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
  
    <title>Select</title>
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
   
    <form action="/sign_select" method="post">

    <label for="country">Select a previous game:</label>
    <select onclick="myFunction()" id="selectTitle" name="selectTitle">
		{{range .}}
      {{with .Title}}
      <option name="selectmain">{{.}}</option>
      {{else}}
      {{end}}
	        
    {{end}}
	</select>
	<label for="country">Developer:</label>
    <select id="infoDeveloper" name="infoDeveloper" disabled>
		{{range .}}
      {{with .Developer}}
      <option name="selectmain">{{.}}</option>
      {{else}}
      {{end}}
	        
    {{end}}
    </select>
	 <select class="hidden" id="infoDeveloper2" name="infoDeveloper2">
		{{range .}}
      {{with .Developer}}
      <option name="selectmain">{{.}}</option>
      {{else}}
      {{end}}
	        
    {{end}}
    </select>
	 <label for="country">Year:</label>
    <select id="infoYear" name="infoYear" disabled>
		{{range .}}
      {{with .Year}}
      <option name="selectmain">{{.}}</option>
      {{else}}
      {{end}}
	        
    {{end}}
	</select>
	 <select class="hidden" id="infoYear2" name="infoYear2">
		{{range .}}
      {{with .Year}}
      <option name="selectmain">{{.}}</option>
      {{else}}
      {{end}}
	        
    {{end}}
	</select>
	
	<label for="country">Previous Negative Count:</label>
    <select id="infoNegative" name="infoNegative" disabled>
		{{range .}}
      {{with .Negative}}
      <option name="selectmain">{{.}}</option>
      {{else}}
      {{end}}
	        
    {{end}}
    </select>
	<label for="country">Previous Neutral Count:</label>
    <select id="infoNeutral" name="infoNeutral" disabled>
		{{range .}}
      {{with .Neutral}}
      <option name="selectmain">{{.}}</option>
      {{else}}
      {{end}}
	        
    {{end}}
    </select>
	<label for="country">Previous Positive Count:</label>
    <select id="infoPositive" name="infoPositive" disabled>
		{{range .}}
      {{with .Positive}}
      <option name="selectmain">{{.}}</option>
      {{else}}
      {{end}}
	        
    {{end}}
    </select>
	<label for="country">Previous Score:</label>
    <select id="infoScore" name="infoScore" disabled>
		{{range .}}
      {{with .Score}}
      <option name="selectmain">{{.}}%</option>
      {{else}}
      {{end}}
	        
    {{end}}
    </select>
    <input type="submit" value="Update">
	<script>
	function myFunction() {
	document.getElementById("infoYear").selectedIndex = document.getElementById("selectTitle").selectedIndex
		document.getElementById("infoDeveloper").selectedIndex = document.getElementById("selectTitle").selectedIndex
		document.getElementById("infoYear2").selectedIndex = document.getElementById("selectTitle").selectedIndex
		document.getElementById("infoDeveloper2").selectedIndex = document.getElementById("selectTitle").selectedIndex
		document.getElementById("infoNegative").selectedIndex = document.getElementById("selectTitle").selectedIndex
		document.getElementById("infoNeutral").selectedIndex = document.getElementById("selectTitle").selectedIndex
		document.getElementById("infoPositive").selectedIndex = document.getElementById("selectTitle").selectedIndex
		document.getElementById("infoScore").selectedIndex = document.getElementById("selectTitle").selectedIndex
}
	</script>
</div>
	
  </body>
</html>
`))
func sign_select(w http.ResponseWriter, r *http.Request) {
        
		
        title = r.FormValue("selectTitle")
						developer = r.FormValue("infoDeveloper2")
						year,_ = strconv.Atoi(r.FormValue("infoYear2"))
						
      
       
        http.Redirect(w, r, "/basic.html", http.StatusFound)
}

