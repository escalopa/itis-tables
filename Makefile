deploy:
	echo "Building image..."
	docker build -f ./Dockerfile -t dekuyo/itis-table:${TAG} --target=production --no-cache .
	echo "Deploying image to Docker Hub..."
	docker image push dekuyo/itis-table:${TAG}

test:
	go test -coverprofile=coverage.txt  -covermode=count ./...

.PHONY: deploy test
