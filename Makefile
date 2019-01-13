update:
	if [ -d "linuxsuren.github.io" ]; then \
	    cd linuxsuren.github.io && git pull; \
	else \
		git clone git@github.com:LinuxSuRen/linuxsuren.github.io.git; \
	fi

live:
	make update
	hugo server

deploy:
	make update
	hugo
	cp -r public/* linuxsuren.github.io
	cd linuxsuren.github.io && git add . && git commit -m 'deploy' && git push