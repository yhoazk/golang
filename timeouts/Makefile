all: clean comp run

clean:
	rm -rf ./timeouts.exe 
	rm -rf ./*~ 

comp:
	go build timeouts.go

run:
	./timeouts.exe
