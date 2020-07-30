test:
	go test . -coverprofile=cov.out

report:
	go tool cover -func=cov.out
