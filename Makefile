build:
	 gox -osarch="darwin/amd64" -output="resources/alfred-esa-workflow" github.com/kyokomi/alfred-esa-workflow/cmd/alfred-esa-workflow
	 alfreder -i resources/info.plist

