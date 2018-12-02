
all : day1.c day1b.c
	emcc day1.c -o day1.html --preload-file data.txt
	emcc day1b.c -o day1b.html --preload-file data.txt

clean :
	rm -rf *.js
	rm -rf *.html
	rm -rf *.wasm


