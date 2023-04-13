# gtfs-viewer-api

Use the output of [gtfs-viewer-processing repoqitory](https://github.com/my-ava-portfolio/gtfs-viewer-processing)

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
sudo docker build --rm -t gtfs-viewer-api .
sudo docker run -p 8080:7001 gtfs-viewer-api
```
