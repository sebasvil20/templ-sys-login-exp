run:
	@go run main.go

generate_templ:
	@templ generate

all: generate_templ run
