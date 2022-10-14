echo "building shim"
GOOS=linux GOARCH=arm64 go build -o apt-get-arm64
GOOS=linux GOARCH=amd64 go build -o apt-get-amd64

ARCH=$(go env GOARCH)
FILE=apt-get-$ARCH

echo "Current arch: $ARCH"
echo "Copying shim from $FILE to /usr/bin/apt-get"
sudo cp $FILE /usr/bin/apt-get
