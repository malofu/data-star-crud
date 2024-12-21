# Starting a new Golang + D* project

## Install

go install github.com/air-verse/air@latest
go get -u github.com/go-chi/chi/v5
go install github.com/a-h/templ/cmd/templ@latest
go get github.com/a-h/templ
go get github.com/starfederation/datastar/sdk/go

## Configuring Templ with Air

cmd = "templ generate && go build -o ./tmp/main.exe ." // execute templ generation after build
exclude_file = ["*_templ.go"] // avoid infinite loop
include_ext = ["go", "tpl", "tmpl", "html", "templ"] // add templ files to be watched

## Live server

air
