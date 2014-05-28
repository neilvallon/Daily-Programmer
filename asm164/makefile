helloasm.o:
	nasm -f macho -o helloasm.o hello.asm

helloc.o:
	gcc -g -c -Wall -m32 -o helloc.o hello.c

hello: helloc.o helloasm.o
	gcc -v -m32 -g -Wall -o hello helloasm.o helloc.o
	rm helloc.o helloasm.o
