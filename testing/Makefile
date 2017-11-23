APP=go_testing
all: clean comp run

clean:
	rm -rf ./${APP}.exe 
	rm -rf ./*~ 

comp:
	go build ${APP}.go

run:
	./${APP}.exe
