target triple = "x86_64-pc-windows-msvc19.39.33523"

@.str.literal.0 = constant [17 x i8] c"chcp 65001 > nul\00"
@.str.literal.1 = constant [11 x i8] c"Sum: %lld\0A\00"

declare i32 @printf(i8* %format, ...)

declare i32 @system(i8* %command)

declare i8* @__acrt_iob_func(i32 %fd)

declare i8* @fgets(i8* %buf, i32 %n, i8* %stream)

declare i8* @malloc(i64 %size)

declare i64 @strlen(i8* %s)

declare i64 @atoll(i8* %s)

define i8* @readString() {
0:
	%1 = call i8* @malloc(i64 256)
	%2 = call i8* @__acrt_iob_func(i32 0)
	%3 = call i8* @fgets(i8* %1, i32 256, i8* %2)
	%4 = call i64 @strlen(i8* %1)
	%5 = icmp sgt i64 %4, 0
	br i1 %5, label %6, label %11

6:
	%7 = sub i64 %4, 1
	%8 = getelementptr i8, i8* %1, i64 %7
	%9 = load i8, i8* %8
	%10 = icmp eq i8 %9, 10
	br i1 %10, label %12, label %11

11:
	ret i8* %1

12:
	store i8 0, i8* %8
	br label %11
}

define i64 @readInt() {
0:
	%1 = alloca [64 x i8]
	%2 = getelementptr [64 x i8], [64 x i8]* %1, i64 0, i64 0
	%3 = call i8* @__acrt_iob_func(i32 0)
	%4 = call i8* @fgets(i8* %2, i32 64, i8* %3)
	%5 = call i64 @atoll(i8* %2)
	ret i64 %5
}

define i32 @main() {
entry:
	%0 = alloca i32
	store i32 0, i32* %0
	%1 = getelementptr [17 x i8], [17 x i8]* @.str.literal.0, i32 0, i32 0
	%2 = call i32 @system(i8* %1)
	%3 = alloca [5 x i64]
	%4 = getelementptr [5 x i64], [5 x i64]* %3, i64 0, i64 0
	store i64 1, i64* %4
	%5 = getelementptr [5 x i64], [5 x i64]* %3, i64 0, i64 1
	store i64 2, i64* %5
	%6 = getelementptr [5 x i64], [5 x i64]* %3, i64 0, i64 2
	store i64 3, i64* %6
	%7 = getelementptr [5 x i64], [5 x i64]* %3, i64 0, i64 3
	store i64 4, i64* %7
	%8 = getelementptr [5 x i64], [5 x i64]* %3, i64 0, i64 4
	store i64 5, i64* %8
	%9 = getelementptr [5 x i64], [5 x i64]* %3, i64 0, i64 0
	%10 = alloca { i64, i64* }
	%11 = insertvalue { i64, i64* } zeroinitializer, i64 5, 0
	%12 = insertvalue { i64, i64* } %11, i64* %9, 1
	store { i64, i64* } %12, { i64, i64* }* %10
	%13 = load { i64, i64* }, { i64, i64* }* %10
	%14 = alloca { i64, i64* }
	store { i64, i64* } %13, { i64, i64* }* %14
	%15 = alloca i64
	store i64 0, i64* %15
	%16 = load { i64, i64* }, { i64, i64* }* %14
	%17 = extractvalue { i64, i64* } %16, 0
	%18 = extractvalue { i64, i64* } %16, 1
	%19 = alloca i64
	%20 = alloca i64
	store i64 0, i64* %20
	br label %forin.cond.1

forin.cond.1:
	%21 = load i64, i64* %20
	%22 = icmp slt i64 %21, %17
	br i1 %22, label %forin.body.2, label %forin.exit.4

forin.body.2:
	%23 = load i64, i64* %20
	%24 = getelementptr i64, i64* %18, i64 %23
	%25 = load i64, i64* %24
	store i64 %25, i64* %19
	%26 = load i64, i64* %15
	%27 = load i64, i64* %19
	%28 = add i64 %26, %27
	store i64 %28, i64* %15
	br label %forin.step.3

forin.step.3:
	%29 = load i64, i64* %20
	%30 = add i64 1, %29
	store i64 %30, i64* %20
	br label %forin.cond.1

forin.exit.4:
	%31 = load i64, i64* %15
	%32 = getelementptr [11 x i8], [11 x i8]* @.str.literal.1, i32 0, i32 0
	%33 = call i32 (i8*, ...) @printf(i8* %32, i64 %31)
	%34 = load i32, i32* %0
	ret i32 %34
}
