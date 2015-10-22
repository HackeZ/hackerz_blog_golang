{{define "header"}}
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
  <link rel="shortcut icon" href="/static/img/favicon.png">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
  <style>
  body{font-size: 63%;}
.site{width:700px;display: block; text-align: justify; margin: 2em auto 1em auto;}
.title{font-size: 3em;text-align: center; font-family: Georgia;}
.content{ line-height: 1.6em; font-size: 1.6em; font-family:STXihei,Microsoft Yahei,Arial;word-wrap: break-word;word-break:break-all;}
hr{margin:2em;filter: alpha(opacity=50);opacity:0.5;}
img{width: 100%; height: auto;}
blockquote {
background:#f9f9f9;
border-left:10px solid #ccc;
margin:1.5em 10px;
padding:.5em 10px;
quotes:"\201C""\201D""\2018""\2019";
}
blockquote:before {
color:#ccc;
content:open-quote;
font-size:4em;
line-height:.1em;
margin-right:.25em;
vertical-align:-.4em;
}
blockquote p {
font-style: italic;
display:inline;
}
code {
  font-family: Consolas;
  font-size: 16px;
}
    ::selection{
        background-color:#99BBFF;
        color:#fff;
    }
    </style>
{{end}}