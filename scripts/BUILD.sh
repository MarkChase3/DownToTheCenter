cd code
go build
GOOS=windows go build
STR1=$(base64 -w 0 ../images/player.png)
STR2=$(base64 -w 0 ../images/tileset1.png)
STR3=$(base64 -w 0 ../maps/jsons/map1.json)
sed -i "s@IMAGES_PLAYER_DATA@$STR1@g" fs/fs_wasm.go
sed -i "s@IMAGES_TILESET1_DATA@$STR2@g" fs/fs_wasm.go
sed -i "s@MAPS_JSONS_MAP1_DATA@$STR3@g" fs/fs_wasm.go
GOOS=js GOARCH=wasm go build -o DownToTheCenter.wasm
mv DownToTheCenter ../build/DownToTheCenter
mv DownToTheCenter.exe ../build/DownToTheCenter.exe
mv DownToTheCenter.wasm ../build/DownToTheCenter.wasm
sudo cp ../build/DownToTheCenter.wasm /var/www/html/DownToTheCenter/
sudo cp ../build/index.html /var/www/html/DownToTheCenter/
sudo cp ../build/wasm_exec.js /var/www/html/DownToTheCenter/
cd ..