build:
	GOOS=linux go build -o mta.bin

zip:
	zip deployment.zip mta.bin