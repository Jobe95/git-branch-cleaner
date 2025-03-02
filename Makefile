build:
	go build -o git-branch-cleaner

install: build
	sudo mv git-branch-cleaner /usr/local/bin/

uninstall:
	sudo rm /usr/local/bin/git-branch-cleaner