target triple = "x86_64-pc-windows-msvc19.39.33523"

@.str.literal.0 = constant [17 x i8] c"chcp 65001 > nul\00"
@.str.literal.1 = constant [11 x i8] c"Enter name\00"
@.str.literal.2 = constant [4 x i8] c"%s\0A\00"
@.str.literal.3 = constant [10 x i8] c"Enter age\00"
@.str.literal.4 = constant [4 x i8] c"%s\0A\00"
@.str.literal.5 = constant [10 x i8] c"Name: %s\0A\00"
@.str.literal.6 = constant [11 x i8] c"Age: %lld\0A\00"

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
	%3 = getelementptr [11 x i8], [11 x i8]* @.str.literal.1, i32 0, i32 0
	%4 = getelementptr [4 x i8], [4 x i8]* @.str.literal.2, i32 0, i32 0
	%5 = call i32 (i8*, ...) @printf(i8* %4, i8* %3)
	%6 = call i8* @readString()
	%7 = alloca i8*
	store i8* %6, i8** %7
	%8 = getelementptr [10 x i8], [10 x i8]* @.str.literal.3, i32 0, i32 0
	%9 = getelementptr [4 x i8], [4 x i8]* @.str.literal.4, i32 0, i32 0
	%10 = call i32 (i8*, ...) @printf(i8* %9, i8* %8)
	%11 = call i64 @readInt()
	%12 = alloca i64
	store i64 %11, i64* %12
	%13 = load i8*, i8** %7
	%14 = getelementptr [10 x i8], [10 x i8]* @.str.literal.5, i32 0, i32 0
	%15 = call i32 (i8*, ...) @printf(i8* %14, i8* %13)
	%16 = load i64, i64* %12
	%17 = getelementptr [11 x i8], [11 x i8]* @.str.literal.6, i32 0, i32 0
	%18 = call i32 (i8*, ...) @printf(i8* %17, i64 %16)
	%19 = load i32, i32* %0
	ret i32 %19
}
