{{define "navbar"}}
<header class="navbar navbar-static-top">
    <div class="container">
        <div class="navbar-header">
          <button class="navbar-toggle collapsed" type="button" data-toggle="collapse" data-target="#bs-navbar" aria-controls="bs-navbar" aria-expanded="false">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a href="/" class="navbar-brand">HackerZ Blog</a>
        </div>

        <nav id="bs-navbar" class="collapse navbar-collapse">
          <ul class="nav navbar-nav">
            <li {{if .IsHome}}class="active"{{end}}><a href="/">Home</a></li>
            <li {{if .IsCategory}}class="active"{{end}}><a href="/category">Category</a></li>
            <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">Topic</a></li>
          </ul>
          <ul class="nav navbar-nav navbar-right">
                {{if .IsLogin}}
                <li><a href="/login?exit=rilegou">Exit</a></li>
                {{else}}
                <li><a href="/login">Login in</a></li>
                {{end}}
            </ul>
        </nav>
    </div>
</header>
{{end}}