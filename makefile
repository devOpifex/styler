properties:
	go run scrape/main.go && cp properties.json cmd/

site:
	cd docs && styler && GOOS=js GOARCH=wasm go build -o static/main.wasm ./main.go
