echo "Entering code folder"
cd code
echo "Building for linux"
go build
echo "Building for windows"
GOOS=windows go build
echo "Preprocessing webassembly filesystem"
STR1=$(base64 -w 0 ../images/player.png)
STR2=$(base64 -w 0 ../images/tileset1.png)
STR3=$(base64 -w 0 ../maps/jsons/map1.json)
STR4=$(base64 -w 0 ../images/zombie.png)
STR5=$(base64 -w 0 ../images/fireball.png)
STR5=$(base64 -w 0 ../images/bow.png)
STR5=$(base64 -w 0 ../images/sword.png)
sed -i "s@IMAGES_PLAYER_DATA@$STR1@g" fs/fs_wasm.go
sed -i "s@IMAGES_TILESET1_DATA@$STR2@g" fs/fs_wasm.go
sed -i "s@MAPS_JSONS_MAP1_DATA@$STR3@g" fs/fs_wasm.go
sed -i "s@IMAGES_ZOMBIE_DATA@$STR4@g" fs/fs_wasm.go
sed -i "s@IMAGES_FIREBALL_DATA@$STR5@g" fs/fs_wasm.go
sed -i "s@IMAGES_SWORD_DATA@$STR5@g" fs/fs_wasm.go
sed -i "s@IMAGES_BOW_DATA@$STR5@g" fs/fs_wasm.go
echo "Building for webassembly"
GOOS=js GOARCH=wasm go build -o DownToTheCenter.wasm
echo "Moving everything to the build folder"
mv DownToTheCenter ../build/DownToTheCenter
mv DownToTheCenter.exe ../build/DownToTheCenter.exe
mv DownToTheCenter.wasm ../build/DownToTheCenter.wasm
echo "If you want, this script can move your web build to the /var/www/html/DownToTheCenter folder for you play it just running apache2ctl start and acessing localhost/DonwToTheCenter. It's optional, and you can exit the program now by using cntrl C 3 times. You msut create that folder before using this (LINUX ONLY). On Win use WSL or get your way out. On mac, I don't have an small idea."
sudo cp ../build/DownToTheCenter.wasm /var/www/html/DownToTheCenter/
sudo cp ../build/index.html /var/www/html/DownToTheCenter/
sudo cp ../build/wasm_exec.js /var/www/html/DownToTheCenter/
cd ..