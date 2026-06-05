target triple = "x86_64-pc-windows-msvc19.39.33523"

@.str.literal.0 = constant [17 x i8] c"chcp 65001 > nul\00"
@.str.literal.1 = constant [7 x i8] c"Array:\00"
@.str.literal.2 = constant [4 x i8] c"%s\0A\00"
@.str.literal.3 = constant [6 x i8] c"%lld\0A\00"
@.str.literal.4 = constant [16 x i8] c"Reversed Array:\00"
@.str.literal.5 = constant [4 x i8] c"%s\0A\00"
@.str.literal.6 = constant [6 x i8] c"%lld\0A\00"

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
	store i64 5, i64* %15
	%16 = alloca i64
	store i64 0, i64* %16
	%17 = getelementptr [7 x i8], [7 x i8]* @.str.literal.1, i32 0, i32 0
	%18 = getelementptr [4 x i8], [4 x i8]* @.str.literal.2, i32 0, i32 0
	%19 = call i32 (i8*, ...) @printf(i8* %18, i8* %17)
	%20 = alloca i64
	store i64 0, i64* %20
	br label %for.cond.1

for.cond.1:
	%21 = load i64, i64* %20
	%22 = load i64, i64* %15
	%23 = icmp slt i64 %21, %22
	br i1 %23, label %for.body.2, label %for.exit.4

for.body.2:
	%24 = load { i64, i64* }, { i64, i64* }* %14
	%25 = load i64, i64* %20
	%26 = extractvalue { i64, i64* } %24, 1
	%27 = getelementptr i64, i64* %26, i64 %25
	%28 = load i64, i64* %27
	%29 = getelementptr [6 x i8], [6 x i8]* @.str.literal.3, i32 0, i32 0
	%30 = call i32 (i8*, ...) @printf(i8* %29, i64 %28)
	br label %for.step.3

for.step.3:
	%31 = load i64, i64* %20
	%32 = add i64 %31, 1
	store i64 %32, i64* %20
	br label %for.cond.1

for.exit.4:
	%33 = alloca i64
	store i64 0, i64* %33
	br label %for.cond.5

for.cond.5:
	%34 = load i64, i64* %33
	%35 = load i64, i64* %15
	%36 = sdiv i64 %35, 2
	%37 = icmp slt i64 %34, %36
	br i1 %37, label %for.body.6, label %for.exit.8

for.body.6:
	%38 = load { i64, i64* }, { i64, i64* }* %14
	%39 = load i64, i64* %33
	%40 = extractvalue { i64, i64* } %38, 1
	%41 = getelementptr i64, i64* %40, i64 %39
	%42 = load i64, i64* %41
	store i64 %42, i64* %16
	%43 = load { i64, i64* }, { i64, i64* }* %14
	%44 = load i64, i64* %15
	%45 = sub i64 %44, 1
	%46 = load i64, i64* %33
	%47 = sub i64 %45, %46
	%48 = extractvalue { i64, i64* } %43, 1
	%49 = getelementptr i64, i64* %48, i64 %47
	%50 = load i64, i64* %49
	%51 = load { i64, i64* }, { i64, i64* }* %14
	%52 = extractvalue { i64, i64* } %51, 1
	%53 = load i64, i64* %33
	%54 = getelementptr i64, i64* %52, i64 %53
	store i64 %50, i64* %54
	%55 = load i64, i64* %16
	%56 = load { i64, i64* }, { i64, i64* }* %14
	%57 = extractvalue { i64, i64* } %56, 1
	%58 = load i64, i64* %15
	%59 = sub i64 %58, 1
	%60 = load i64, i64* %33
	%61 = sub i64 %59, %60
	%62 = getelementptr i64, i64* %57, i64 %61
	store i64 %55, i64* %62
	br label %for.step.7

for.step.7:
	%63 = load i64, i64* %33
	%64 = add i64 %63, 1
	store i64 %64, i64* %33
	br label %for.cond.5

for.exit.8:
	%65 = getelementptr [16 x i8], [16 x i8]* @.str.literal.4, i32 0, i32 0
	%66 = getelementptr [4 x i8], [4 x i8]* @.str.literal.5, i32 0, i32 0
	%67 = call i32 (i8*, ...) @printf(i8* %66, i8* %65)
	%68 = alloca i64
	store i64 0, i64* %68
	br label %for.cond.9

for.cond.9:
	%69 = load i64, i64* %68
	%70 = load i64, i64* %15
	%71 = icmp slt i64 %69, %70
	br i1 %71, label %for.body.10, label %for.exit.12

for.body.10:
	%72 = load { i64, i64* }, { i64, i64* }* %14
	%73 = load i64, i64* %68
	%74 = extractvalue { i64, i64* } %72, 1
	%75 = getelementptr i64, i64* %74, i64 %73
	%76 = load i64, i64* %75
	%77 = getelementptr [6 x i8], [6 x i8]* @.str.literal.6, i32 0, i32 0
	%78 = call i32 (i8*, ...) @printf(i8* %77, i64 %76)
	br label %for.step.11

for.step.11:
	%79 = load i64, i64* %68
	%80 = add i64 %79, 1
	store i64 %80, i64* %68
	br label %for.cond.9

for.exit.12:
	%81 = load i32, i32* %0
	ret i32 %81
}
