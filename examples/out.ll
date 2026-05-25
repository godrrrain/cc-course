target triple = "x86_64-pc-windows-msvc19.39.33523"

@.str.literal.0 = constant [11 x i8] c"Sum: %lld\0A\00"

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = alloca i32
	%1 = alloca [5 x i64]
	%2 = getelementptr [5 x i64], [5 x i64]* %1, i64 0, i64 0
	store i64 1, i64* %2
	%3 = getelementptr [5 x i64], [5 x i64]* %1, i64 0, i64 1
	store i64 2, i64* %3
	%4 = getelementptr [5 x i64], [5 x i64]* %1, i64 0, i64 2
	store i64 3, i64* %4
	%5 = getelementptr [5 x i64], [5 x i64]* %1, i64 0, i64 3
	store i64 4, i64* %5
	%6 = getelementptr [5 x i64], [5 x i64]* %1, i64 0, i64 4
	store i64 5, i64* %6
	%7 = getelementptr [5 x i64], [5 x i64]* %1, i64 0, i64 0
	%8 = alloca { i64, i64* }
	%9 = insertvalue { i64, i64* } zeroinitializer, i64 5, 0
	%10 = insertvalue { i64, i64* } %9, i64* %7, 1
	store { i64, i64* } %10, { i64, i64* }* %8
	%11 = load { i64, i64* }, { i64, i64* }* %8
	%12 = alloca { i64, i64* }
	store { i64, i64* } %11, { i64, i64* }* %12
	%13 = alloca i64
	store i64 0, i64* %13
	%14 = load { i64, i64* }, { i64, i64* }* %12
	%15 = extractvalue { i64, i64* } %14, 0
	%16 = extractvalue { i64, i64* } %14, 1
	%17 = alloca i64
	%18 = alloca i64
	store i64 0, i64* %18
	br label %forin.cond.1

forin.cond.1:
	%19 = load i64, i64* %18
	%20 = icmp slt i64 %19, %15
	br i1 %20, label %forin.body.2, label %forin.exit.4

forin.body.2:
	%21 = load i64, i64* %18
	%22 = getelementptr i64, i64* %16, i64 %21
	%23 = load i64, i64* %22
	store i64 %23, i64* %17
	%24 = load i64, i64* %13
	%25 = load i64, i64* %17
	%26 = add i64 %24, %25
	store i64 %26, i64* %13
	br label %forin.step.3

forin.step.3:
	%27 = load i64, i64* %18
	%28 = add i64 1, %27
	store i64 %28, i64* %18
	br label %forin.cond.1

forin.exit.4:
	%29 = load i64, i64* %13
	%30 = getelementptr [11 x i8], [11 x i8]* @.str.literal.0, i32 0, i32 0
	%31 = call i32 (i8*, ...) @printf(i8* %30, i64 %29)
	%32 = load i32, i32* %0
	ret i32 %32
}
