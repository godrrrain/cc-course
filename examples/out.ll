target triple = "x86_64-pc-linux-gnu"

@.str.literal.0 = constant [14 x i8] c"Hello, World!\00"
@.str.literal.1 = constant [4 x i8] c"%s\0A\00"
@.str.literal.2 = constant [16 x i8] c"VOVOOOOOOOOOON!\00"
@.str.literal.3 = constant [4 x i8] c"%s\0A\00"

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = alloca i32
	%1 = getelementptr [14 x i8], [14 x i8]* @.str.literal.0, i32 0, i32 0
	%2 = getelementptr [4 x i8], [4 x i8]* @.str.literal.1, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i8* %1)
	%4 = getelementptr [16 x i8], [16 x i8]* @.str.literal.2, i32 0, i32 0
	%5 = getelementptr [4 x i8], [4 x i8]* @.str.literal.3, i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i8* %4)
	%7 = load i32, i32* %0
	ret i32 %7
}
