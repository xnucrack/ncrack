all:
	@gcc -o libcool_lib.dylib -dynamiclib library.c
	@gcc main.c -o cool_bin -L. -lcool_lib
