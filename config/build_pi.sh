build_sensorbee --only-generate-source
env GOOS=linux GOARCH=arm GOARM=7 go build -o sensorbee_pi sensorbee_main.go
