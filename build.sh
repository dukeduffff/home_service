export GOOS=linux
export GOARCH=amd64
#export CGO_ENABLED=1
rm -rf ./output

mkdir output
go build -o ./output/home
cp -r ./release/* ./output

zip ./output/home-${GOOS}-${GOARCH}.zip ./output/*