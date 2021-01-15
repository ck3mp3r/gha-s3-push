VERSION=0.1.0

build:
	docker build -t ckemper/gha-s3push:${VERSION} .

publish:
	docker push \
		ckemper/gha-s3push:${VERSION}

publish-latest:
	docker tag \
		ckemper/gha-s3push:${VERSION} \
		ckemper/gha-s3push:latest
	docker push \
		ckemper/gha-s3push:latest
