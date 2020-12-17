# Katheesh's Public Files

electron persmission
	sudo chown root:root chrome-sandbox
	sudo chmod 4755 chrome-sandbox


BlueTooth Connection New Device
	
	sudo systemctl restart bluetooth
	cd //bin
	sudo ./bluetoothctl
	connect 41:42:F3:73:62:1A
	
composer run plateform not reqs
	
	composer install --ignore-platform-reqs

	
Laravel new project command not working run these steps
	
	nano ~/.bash_profile 
	And paste
	export PATH=~/.config/composer/vendor/bin:$PATH
	do source ~/.bash_profile and enjoy ;)

Create zip file

	zip -r output_file.zip file1 folder1

Run deb package 

	sudo dpkg -i filename.deb

Flutter App create 

	export PATH=`pwd`/flutter/bin:$PATH
	echo $PATH

Provide Port Permission

	sudo chmod a+rw /dev/ttyACM0

Provide folder permission

	sudo chmod -R 777 assets/

Electron sandbox issue solve

	sudo chown root:root chrome-sandbox
	sudo chmod 4755 chrome-sandbox

SQLmap

	sqlmap -u http://doamin.com/page.php?id=10 --dbs
	sqlmap -u http://doamin.com/page.php?id=10 -D dbname --tables
	sqlmap -u http://doamin.com/page.php?id=10 -D dbname -T tablename --columns
	sqlmap -u http://doamin.com/page.php?id=10 -D dbname -T tablename -C columnname,columnname --dump

Git global setup

	git config --global user.name "Katheeskumar Sithamparapillai"
	git config --global user.email "katheesh2016@outlook.com"
	
	git config --global credential.helper store

Create a new repository

	git clone https://gitlab.com/katheesh/reposity.git
	cd zonal-edu-web
	touch README.md
	git add README.md
	git commit -m "add README"
	git push -u origin master

Push an existing folder

	cd existing_folder
	git init
	git remote add origin https://gitlab.com/katheesh/reposity.git
	git add .
	git commit -m "Initial commit"
	git push -u origin master

Push an existing Git repository

	cd existing_repo
	git remote rename origin old-origin
	git remote add origin https://gitlab.com/katheesh/reposity.git
	git push -u origin --all
	git push -u origin --tags


cd foldername   	:change directy

mv filename targetname  :move file or folder and rename

rmdir foldername remove :remove folder

rm filename		:remove file

rm -rf filename   	:remove file force

vi filename ==========> editfile ===> press i key to edit   ====> to save press ctrl + Esc  ==> :x press enter 


https://forum.arduino.cc/index.php?topic=476291.0
































