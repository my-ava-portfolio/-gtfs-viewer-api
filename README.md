# gtfs_viewer



# Install go

```bash
sudo apt-get update && sudo apt-get install golang-go
```

# set env var
```
export PATH="$HOME/go/bin:$PATH"
```

## Docker run

```bash
sudo docker build --rm -t gtfs-viewer-app .
sudo docker run -p 8080:7001 gtfs-viewer-app
```