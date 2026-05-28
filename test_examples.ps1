param()

$ErrorActionPreference = "Stop"

# Check prerequisites
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "ERROR: Go not found in PATH" -ForegroundColor Red
    exit 1
}
if (-not (Get-Command clang -ErrorAction SilentlyContinue)) {
    Write-Host "ERROR: clang not found in PATH" -ForegroundColor Red
    exit 1
}
Write-Host ""

function Get-ActualOutput($exePath) {
    $psi = New-Object System.Diagnostics.ProcessStartInfo
    $psi.FileName = $exePath
    $psi.RedirectStandardOutput = $true
    $psi.UseShellExecute = $false

    $proc = [System.Diagnostics.Process]::Start($psi)
    $output = $proc.StandardOutput.ReadToEnd()
    $proc.WaitForExit()

    return @{ Output = $output; ExitCode = $proc.ExitCode }
}

$tests = @(
    @{ Name = "hello.dart"; Expected = @"
Hello, World!
"@ 
    }
    @{ Name = "arithmetic.dart"; Expected = @"
Sum: 8
"@ 
    }
    @{ Name = "array.dart"; Expected = @"
Sum: 15
"@ 
    }
    @{ Name = "fibonacci.dart"; Expected = @"
Fibonacci(10) = 55
"@ 
    }
    @{ Name = "for_loop.dart"; Expected = @"
Count: 1
Count: 2
Count: 3
Count: 4
Count: 5
"@ 
    }
    @{ Name = "if_else.dart"; Expected = @"
Adult
"@ 
    }
    @{ Name = "reverse_array.dart"; Expected = @"
Array:
1
2
3
4
5
Reversed Array:
5
4
3
2
1
"@ 
    }
)

function Normalize-Output($text) {
    $lines = $text -split "`r`n|`n"
    $trimmed = $lines | ForEach-Object { $_.TrimEnd("`r") }
    while ($trimmed.Count -gt 0 -and $trimmed[-1] -eq "") {
        $trimmed = $trimmed[0..($trimmed.Count - 2)]
    }
    if ($trimmed.Count -eq 0) {
        return ""
    }
    return $trimmed -join "`n"
}

$passed = 0
$failed = 0

foreach ($test in $tests) {
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host "  $($test.Name)" -ForegroundColor Cyan
    Write-Host "========================================" -ForegroundColor Cyan

    $tempDir = Join-Path $env:TEMP "cc_test_$([System.IO.Path]::GetRandomFileName())"
    $null = New-Item -ItemType Directory -Path $tempDir -Force
    $output = Join-Path $tempDir "out"

    Write-Host "Compiling..." -ForegroundColor Yellow

    $compileOutput = go run .\cmd\compiler\main.go -i "examples\$($test.Name)" -o $output 2>&1
    if ($LASTEXITCODE -ne 0) {
        Write-Host "COMPILATION FAILED (exit code: $LASTEXITCODE)" -ForegroundColor Red
        Write-Host $compileOutput
        $failed++
        Remove-Item -Recurse -Force $tempDir -ErrorAction SilentlyContinue
        Write-Host ""
        continue
    }

    $exePath = "$output.exe"
    if (-not (Test-Path $exePath)) {
        Write-Host "Executable not found: $exePath" -ForegroundColor Red
        $failed++
        Remove-Item -Recurse -Force $tempDir -ErrorAction SilentlyContinue
        Write-Host ""
        continue
    }

    Write-Host "Running..." -ForegroundColor Yellow

    $result = Get-ActualOutput $exePath
    $actual = $result.Output
    $exitCode = $result.ExitCode

    Write-Host ""
    Write-Host "Output:" -ForegroundColor Yellow
    Write-Host $actual

    if ($exitCode -ne 0) {
        Write-Host "Warning: non-zero exit code ($exitCode)" -ForegroundColor DarkYellow
    }

    $normalizedExpected = Normalize-Output $test.Expected
    $normalizedActual = Normalize-Output $actual

    if ($normalizedActual -eq $normalizedExpected) {
        Write-Host "PASS" -ForegroundColor Green
        $passed++
    }
    else {
        Write-Host "FAIL" -ForegroundColor Red
        Write-Host "--- Expected:" -ForegroundColor Red
        Write-Host $test.Expected
        Write-Host "--- Actual:" -ForegroundColor Red
        Write-Host $actual
        $failed++
    }

    Remove-Item -Recurse -Force $tempDir -ErrorAction SilentlyContinue
    Write-Host ""
}

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Results: $passed passed, $failed failed" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan

if ($failed -gt 0) {
    exit 1
}
