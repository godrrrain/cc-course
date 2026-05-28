target triple = "x86_64-pc-windows-msvc19.39.33523"

@.str.literal.0 = constant [17 x i8] c"chcp 65001 > nul\00"
@.str.literal.1 = constant [22 x i8] c"Fibonacci(10) = %lld\0A\00"

declare i32 @printf(i8* %format, ...)

declare i32 @system(i8* %command)

define i64 @fibonacci(i64 %n) {
entry:
	%0 = alloca i64
	%1 = alloca i64
	store i64 %n, i64* %1
	%2 = load i64, i64* %1
	%3 = icmp sle i64 %2, 1
	br i1 %3, label %if.then.1, label %if.end.2

if.then.1:
	%4 = load i64, i64* %1
	ret i64 %4

if.end.2:
	%5 = load i64, i64* %1
	%6 = sub i64 %5, 1
	%7 = call i64 @fibonacci(i64 %6)
	%8 = load i64, i64* %1
	%9 = sub i64 %8, 2
	%10 = call i64 @fibonacci(i64 %9)
	%11 = add i64 %7, %10
	ret i64 %11

unreachable.3:
	br label %if.end.2

unreachable.4:
	%12 = load i64, i64* %0
	ret i64 %12
}

define i32 @main() {
entry:
	%0 = alloca i32
	%1 = getelementptr [17 x i8], [17 x i8]* @.str.literal.0, i32 0, i32 0
	%2 = call i32 @system(i8* %1)
	%3 = call i64 @fibonacci(i64 10)
	%4 = alloca i64
	store i64 %3, i64* %4
	%5 = load i64, i64* %4
	%6 = getelementptr [22 x i8], [22 x i8]* @.str.literal.1, i32 0, i32 0
	%7 = call i32 (i8*, ...) @printf(i8* %6, i64 %5)
	%8 = load i32, i32* %0
	ret i32 %8
}
