$platforms = @(
    "windows/amd64",
    "linux/amd64",
    "darwin/amd64",
    "darwin/arm64"
)

foreach ($platform in $platforms) {
    $parts = $platform.Split("/")
    $GOOS = $parts[0]
    $GOARCH = $parts[1]

    $output = "bin/flumint-$GOOS-$GOARCH"

    if ($GOOS -eq "windows") {
        $output += ".exe"
    }

    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH

    go build -o $output
}

Write-Host "Build finished."
