go_run:
	@go run main.go

generate_templ:
	@templ generate

run: generate_templ go_run
