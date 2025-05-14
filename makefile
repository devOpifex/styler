properties:
	go run scrape/main.go && cp properties.json cmd/

site:
	cd docs && GOOS=js GOARCH=wasm go build -o static/main.wasm ./main.go
