run: build
	( cd back && ./main )

build: 
	( cd back && go build )
	( cd front && npm run build )


