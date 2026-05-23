target triple = "x86_64-pc-linux-gnu"

@.str.literal.0 = constant [13 x i8] c"Count: %lld\0A\00"

declare i32 @printf(i8* %format, ...)

define i32 @main() {
entry:
	%0 = alloca i32
	%1 = alloca i64
	store i64 1, i64* %1
	br label %for.cond.1

for.cond.1:
	%2 = load i64, i64* %1
	%3 = icmp sle i64 %2, 5
	br i1 %3, label %for.body.2, label %for.exit.4

for.body.2:
	%4 = load i64, i64* %1
	%5 = getelementptr [13 x i8], [13 x i8]* @.str.literal.0, i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i64 %4)
	br label %for.step.3

for.step.3:
	%7 = load i64, i64* %1
	%8 = add i64 %7, 1
	store i64 %8, i64* %1
	br label %for.cond.1

for.exit.4:
	%9 = load i32, i32* %0
	ret i32 %9
}
