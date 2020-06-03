build:
   @GOARCH=wasm GOOS=js go build -o app.wasm ./app
   @go build -o httpServer ./server

run: build
   PORT=8000 ./httpServer

deploy: build
   gcloud app deploy --project=attractionsTips
