set -x

go get -v
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ms-todo || exit $?

if [ ! -d "./bin" ]; then
  mkdir bin
else
  rm -rf bin/*
fi

mv ms-todo ./bin
cp -r conf/ ./bin

echo "Build ok"
exit 0
