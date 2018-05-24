package IMA

import (
        "html/template"
        "net/http"
        "time"
		"strconv"
        "appengine"
        "appengine/datastore"
)



func init() {
        http.HandleFunc("/basic.html", root_basic)
        http.HandleFunc("/sign_basic", sign_basic)
	
}


func assignmentKey_basic(c appengine.Context) *datastore.Key {
        return datastore.NewKey(c, "IMA", "default_assignment", 0, nil)
}

func root_basic(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("IMA").Ancestor(assignmentKey_basic(c)).Order("-Date").Limit(1)
        greetings := make([]Assignment, 0, 1)
        if _, err := q.GetAll(c, &greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        if err := assignmentTemplate_basic.Execute(w, title); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}
	
var assignmentTemplate_basic = template.Must(template.New("book").Parse(`
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
display: block;
text-align: center;
color: LightGrey;
font-size: 26px;
}
basich {
font-family: verdana;
font-size: 14px;
display: fixed;
float: left;
padding: 1%;
font-size: 26px;
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
	visibility: hidden;
    width: 0%;
    padding: 0px 0px;
    margin: 8px 0;
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


<script>
var count = 0;

id = [""];
like = [""]
dislike = [""]

if (negative == null)
{
var negative = 0;
}
if (neutral == null)
{
var neutral = 0;
}
if (positive == null)
{
var positive = 0;
}
if (score == null)
{
var score = 0;
}
negativekeyword = ["Bad", "Crap", "Terrible","Mediocre", "Shit", "Worst", "Overrated", "Hate", "Dislike", "Sucks", "Buggy", "Mess", "Garbage", "Trash", "Overhyped", "Fail"]
neutralkeyword = ["Gameplay", "Walkthrough","Let's Play", "Lets Play", "Trailer", "Playthrough", "Review", "First Impression", "Thoughts", "Opinion"]
positivekeyword = ["Good", "Fun","Best", "Cool", "Awesome", "Perfect", "Amazing", "Love", "Like", "Excellent", "Better", "Praise", "Worthy", "Great"]

// Your use of the YouTube API must comply with the Terms of Service:
// https://developers.google.com/youtube/terms
// Called automatically when JavaScript client library is loaded.
function onClientLoad() {
    gapi.client.load('youtube', 'v3', onYouTubeApiLoad);
}
// Called automatically when YouTube API interface is loaded (see line 9).
function onYouTubeApiLoad() {
    gapi.client.setApiKey('AIzaSyBxdFYoKdVqwqWPpIDkoWicGB6zBSMEnUs');
	search();
}
					
					
					
				</script>
				<script src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
<script>
function home() {
window.location = "index.html";
}
</script>
<main onclick="home()">Influencer Marketing Analysis</main></div>
<form action="/sign_basic" name="main" method="post">
 
   
	<input type="text" id="negative_out" name="negative_out" value="">
	<input type="text" id="neutral_out" name="neutral_out" value="">
	<input type="text" id="positive_out" name="positive_out" value="">
		<input type="text" id="score_out" name="score_out" value="">

    <input type="submit" value="Save Results">
  </form>
<div>
  <label style="color: red">Negative: </label>
<br>
    <label id="negative" name="negative" style="color: red">Calculating...</label>
	<br>
</div>
<div>
  <label style="color: lightgrey">Neutral: </label>
<br>
    <label style="color: lightgrey" id="neutral" name="neutral">Calculating...</label>
	<br>
</div>
<div>
  <label style="color: green">Positive: </label>
<br>
    <label id="positive" style="color: green" name="positive">Calculating...</label>
	<br>
</div>
<div>
  <label style="color: orange">Score: </label>
<br>
    <label id="score" name="score" style="color: orange">Calculating...</label>
<br>
    <label style="color: orange" id="desc">Calculating...</label>
	</div>
<script>	
// Called when the search button is clicked in the html code
function search() {
								document.getElementById("negative").innerHTML = "Calculating...";
								document.getElementById("positive").innerHTML = "Calculating...";
								document.getElementById("neutral").innerHTML = "Calculating...";
								document.getElementById("score").innerHTML = "Calculating...";
positive = 0;
negative = 0;
neutral = 0;
    // Use the JavaScript client library to create a search.list() API call.
    var request = gapi.client.youtube.search.list({
        part: 'snippet',
		type: 'video',
        q:document.getElementById("title").innerHTML + " good",
		maxResults: 50
    });
request.execute(function(response) {
                    var len = response.items.length
					negativeobj = document.getElementById("negative");
					neutralobj = document.getElementById("neutral");
					positiveobj = document.getElementById("positive");
					scoreobj = document.getElementById("score");
					descobj = document.getElementById("desc");
					titleobj = document.getElementById("title");
                    for (var i = 0; i < len; i++) {
					id = response.items[i].id.videoId
					if (response.items[i].snippet.title.indexOf(titleobj.innerHTML) !== -1 || response.items[i].snippet.title.indexOf(titleobj.innerHTML.toLowerCase()) !== -1 || response.items[i].snippet.title.indexOf(titleobj.innerHTML.toUpperCase()) !== -1) {
															if (response.items[i].snippet.title.indexOf("Mod") == -1 && response.items[i].snippet.title.indexOf("MOD") == -1 && response.items[i].snippet.title.indexOf("mod") == -1) {
	
					               for (var a = 0; a < negativekeyword.length; a++) {


								                           if (response.items[i].snippet.title.indexOf(negativekeyword[a]) !== -1 || response.items[i].snippet.title.indexOf(negativekeyword[a].toLowerCase()) !== -1 || response.items[i].snippet.title.indexOf(negativekeyword[a].toUpperCase()) !== -1 || response.items[i].snippet.description.indexOf(negativekeyword[a]) !== -1 || response.items[i].snippet.description.indexOf(negativekeyword[a].toLowerCase()) !== -1 || response.items[i].snippet.description.indexOf(negativekeyword[a].toUpperCase()) !== -1) {
						negative+=2;
						score = Math.round((positive/(negative+positive))*100);
					
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																															document.getElementById("score_out").value=score;

											score = score+"%";
					scoreobj.innerHTML = score;

						 $.getJSON("https://www.googleapis.com/youtube/v3/videos", {
					key: "AIzaSyBxdFYoKdVqwqWPpIDkoWicGB6zBSMEnUs",
					part: "snippet,statistics",
					id: id
				}, function(data) {
				    if (data.items[0].statistics.dislikeCount >= data.items[0].statistics.likeCount) {
							positive+=2;
															positiveobj.innerHTML = positive;
															document.getElementById("positive_out").value=document.getElementById("positive").innerHTML;
score = Math.round((positive/(negative+positive))*100);
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																															document.getElementById("score_out").value=score;

																score = score+"%";
					scoreobj.innerHTML = score;

					}
					else
					{
					negative+=2;
								negativeobj.innerHTML = negative;
								   document.getElementById("negative_out").value=document.getElementById("negative").innerHTML;

								score = Math.round((positive/(negative+positive))*100);
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																															document.getElementById("score_out").value=score;

																score = score+"%";
					scoreobj.innerHTML = score;

					}
					
				}
)						
}		
								negativeobj.innerHTML = negative;
								   document.getElementById("negative_out").value=document.getElementById("negative").innerHTML;

						}
						}
						 for (var a = 0; a < neutralkeyword.length; a++) {


								                           if (response.items[i].snippet.title.indexOf(neutralkeyword[a]) !== -1 || response.items[i].snippet.title.indexOf(neutralkeyword[a].toLowerCase()) !== -1 || response.items[i].snippet.title.indexOf(neutralkeyword[a].toUpperCase()) !== -1 || response.items[i].snippet.description.indexOf(neutralkeyword[a]) !== -1 || response.items[i].snippet.description.indexOf(neutralkeyword[a].toLowerCase()) !== -1 || response.items[i].snippet.description.indexOf(neutralkeyword[a].toUpperCase()) !== -1) {
						neutral++;
						

						 $.getJSON("https://www.googleapis.com/youtube/v3/videos", {
					key: "AIzaSyBxdFYoKdVqwqWPpIDkoWicGB6zBSMEnUs",
					part: "snippet,statistics",
					id: id
				}, function(data) {
				    if (data.items[0].statistics.likeCount >= data.items[0].statistics.dislikeCount) {
							positive+=2;
															positiveobj.innerHTML = positive;
																							   document.getElementById("positive_out").value=document.getElementById("positive").innerHTML;

score = Math.round((positive/(negative+positive))*100);
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																										document.getElementById("score_out").value=score;

																score = score+"%";
					scoreobj.innerHTML = score;

					}
					else
					{
					negative+=2;
								negativeobj.innerHTML = negative;
								   document.getElementById("negative_out").value=document.getElementById("negative").innerHTML;

								score = Math.round((positive/(negative+positive))*100);
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																															document.getElementById("score_out").value=score;

																score = score+"%";
					scoreobj.innerHTML = score;

					}
					
				}
)
								neutralobj.innerHTML = neutral;
								document.getElementById("neutral_out").value=document.getElementById("neutral").innerHTML;
						}
						}
						 for (var a = 0; a < positivekeyword.length; a++) {


								                           if (response.items[i].snippet.title.indexOf(positivekeyword[a]) !== -1 || response.items[i].snippet.title.indexOf(positivekeyword[a].toLowerCase()) !== -1 || response.items[i].snippet.title.indexOf(positivekeyword[a].toUpperCase()) !== -1 || response.items[i].snippet.description.indexOf(positivekeyword[a]) !== -1 || response.items[i].snippet.description.indexOf(positivekeyword[a].toLowerCase()) !== -1 || response.items[i].snippet.description.indexOf(positivekeyword[a].toUpperCase()) !== -1) {
						positive+=2;
						score = Math.round((positive/(negative+positive))*100);
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																															document.getElementById("score_out").value=score;

																score = score+"%";
																					scoreobj.innerHTML = score;



						 $.getJSON("https://www.googleapis.com/youtube/v3/videos", {
					key: "AIzaSyBxdFYoKdVqwqWPpIDkoWicGB6zBSMEnUs",
					part: "snippet,statistics",
					id: id
				}, function(data) {
				    if (data.items[0].statistics.likeCount >= data.items[0].statistics.dislikeCount) {
							positive+=2;
															positiveobj.innerHTML = positive;
						document.getElementById("positive_out").value=document.getElementById("positive").innerHTML;

score = Math.round((positive/(negative+positive))*100);
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																															document.getElementById("score_out").value=score;

																score = score+"%";
					scoreobj.innerHTML = score;

				

					}
					else
					{
					negative+=2;
							negativeobj.innerHTML = negative;
							   document.getElementById("negative_out").value=document.getElementById("negative").innerHTML;

								score = Math.round((positive/(negative+positive))*100);
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																															document.getElementById("score_out").value=score;

																score = score+"%";
																					scoreobj.innerHTML = score;


					}
					
				}
)
								positive+=2;
								positiveobj.innerHTML = positive;
								document.getElementById("positive_out").value=document.getElementById("positive").innerHTML;
								score = Math.round((positive/(negative+positive))*100);
					if (score <= 100 && score >= 90) {
					descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very positive";
					}
					else if (score < 90 && score >= 70) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are positive";
					}
					else if (score < 70 && score >= 50) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are mixed";
					}
					else if (score < 50 && score >= 25) {
					
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are negative";
					}
					else if (score < 25) {
										descobj.innerHTML = "Opinions on "+titleobj.innerHTML+" are very negative";
					}
																															document.getElementById("score_out").value=score;

																score = score+"%";
																					scoreobj.innerHTML = score;


						}
						}
                    }}
					

					})
    // Send the request to the API server, call the onSearchResponse function when the data is returned

}


// Triggered by this line: request.execute(onSearchResponse);

 </script>
<head>
        <script src="search.js" type="text/javascript"></script>        
        <script src="https://apis.google.com/js/client.js?onload=onClientLoad" type="text/javascript"></script>
         
    <form action="/sign_basic" method="post" name="main">
    
  </form>
  
  </head>
  
  
  <script>
 
  
  </script>
  <body>


	<div>
	
  <title id="title" style="left: 50%">{{.}}</title>
  
      
</div>









	<script>

	</script>

  </body>
</html>
`))
func sign_basic(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        g := Assignment{
				
                
        }
						g.Title = title
						g.Developer = developer
						g.Year = year
						g.Negative,_ = strconv.Atoi(r.FormValue("negative_out"))
						g.Neutral,_ = strconv.Atoi(r.FormValue("neutral_out"))
						g.Positive,_ = strconv.Atoi(r.FormValue("positive_out"))
						g.Score,_ = strconv.Atoi(r.FormValue("score_out"))

	
	g.Date = time.Now()
	
       
        key := datastore.NewIncompleteKey(c, "IMA", assignmentKey_basic(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
				        http.Redirect(w, r, "/index.html", http.StatusFound)

}




