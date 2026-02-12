platforms=(
"windows/amd64"
"linux/amd64"
"darwin/amd64"
"darwin/arm64"
)

for platform in "${platforms[@]}"
do
  GOOS=${platform%/*}
  GOARCH=${platform#*/}
  output="bin/flumint-$GOOS-$GOARCH"
  if [ $GOOS = "windows" ]; then
    output+='.exe'
  fi
  GOOS=$GOOS GOARCH=$GOARCH go build -o $output
done
