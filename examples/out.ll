target triple = "x86_64-"

@.str.literal.0 = constant [3 x i8] c"2+\00"
@.fmt.str.1 = constant [3 x i8] c"%s\00"
@.str.literal.2 = constant [4 x i8] c"2=?\00"
@.str.literal.3 = constant [2 x i8] c"\0A\00"
@.fmt.str.4 = constant [6 x i8] c"%s %s\00"
@.str.literal.5 = constant [10 x i8] c"Result: 4\00"
@.str.literal.6 = constant [2 x i8] c"\0A\00"
@.fmt.str.7 = constant [6 x i8] c"%s %s\00"
@.fmt.str.8 = constant [6 x i8] c"%f %d\00"

define i32 @main() {
entry:
	%0 = alloca i32
	%1 = getelementptr [3 x i8], [3 x i8]* @.str.literal.0, i32 0, i32 0
	%2 = getelementptr [3 x i8], [3 x i8]* @.fmt.str.1, i32 0, i32 0
	%3 = call i32 (i8*, ...) @printf(i8* %2, i8* %1)
	%4 = getelementptr [4 x i8], [4 x i8]* @.str.literal.2, i32 0, i32 0
	%5 = getelementptr [2 x i8], [2 x i8]* @.str.literal.3, i32 0, i32 0
	%6 = getelementptr [6 x i8], [6 x i8]* @.fmt.str.4, i32 0, i32 0
	%7 = call i32 (i8*, ...) @printf(i8* %6, i8* %4, i8* %5)
	%8 = getelementptr [10 x i8], [10 x i8]* @.str.literal.5, i32 0, i32 0
	%9 = getelementptr [2 x i8], [2 x i8]* @.str.literal.6, i32 0, i32 0
	%10 = getelementptr [6 x i8], [6 x i8]* @.fmt.str.7, i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i8* %8, i8* %9)
	%12 = getelementptr [6 x i8], [6 x i8]* @.fmt.str.8, i32 0, i32 0
	%13 = call i32 (i8*, ...) @printf(i8* %12, double 0x40091EB851EB851F, i32 10)
	%14 = load i32, i32* %0
	ret i32 %14
}

declare i32 @printf(i8* %format, ...)
