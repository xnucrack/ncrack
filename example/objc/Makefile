all:
	@gcc library.m -dynamiclib -framework Foundation -o libcool_lib.dylib
	@clang main.m -o cool_bin -framework Foundation -L. -lcool_lib
