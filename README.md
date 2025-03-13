# Build dev image
```
docker build -t xenexporterdev .
```
# Run dev container
```
docker run -idt -v $(pwd):/go/src -p 5000:5000 --name xenexporter xenexporterdev /bin/bash
```
# Attach to dev container
```
docker exec -it xenexporter /bin/bash
```

# Initialize the project
```
go mod init goxenexporter
go mod tidy
```

# For XenApi
Download SDK from https://www.xenserver.com/downloads, extract the zip file and copy the `XenServer-SDK/XenServerGo/src` folder into the project. Rename `src` folder to `xenapi`.
add the following line at the end of the go.mod file.
```
replace xenapi => ./xenapi
```