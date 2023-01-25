cd code
go build
GOOS=windows go build
GOOS=js GOARCH=wasm go build -o DownToTheCenter.wasm
mv DownToTheCenter ../build/DownToTheCenter
mv DownToTheCenter.exe ../build/DownToTheCenter.exe
mv DownToTheCenter.wasm ../build/DownToTheCenter.wasm
sudo cp ../build/DownToTheCenter.wasm /var/www/html/DownToTheCenter/
sudo cp ../build/index.html /var/www/html/DownToTheCenter/
sudo cp ../build/wasm_exec.js /var/www/html/DownToTheCenter/
cd ..