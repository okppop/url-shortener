# url-shortener
url-shortener is a short URL service.

Using postgresql, redis, offer fast user experience.

## Deployment
There are two ways to deploy this project:

- Compile deployment
- Container deployment

### Compile deployment
Environment requirement:

- Linux environment
    - systemd(optional)
    - root permission
- Go development environment
- Postgresql
- Redis

#### Clone repository:
```
git clone https://github.com/okppop/url-shortener.git

cd url-shortener/
```

#### Compile:
```
go mod tidy

go build .
```

#### Edit config file:
```
cp config.yaml.example config.yaml
```
Then, use any editor edit "config.yaml".

#### Import database schema:

Change host, username, database base your setting.
```
psql -h host -U username database < schema.sql
```

#### Deployment:
```
sudo mkdir -p /usr/local/url-shortener
sudo mv url-shortener config.yaml /usr/local/url-shortener/
sudo cp url-shortener.service /etc/systemd/system/
sudo systemctl daemon-reload
```

#### Start service:
```
sudo systemctl start url-shortener.service
```

### Container Deployment

Environment requirement:

- Linux environment
- Docker or other container manager
- Go development environment
- Postgresql
- Redis

#### Clone repository:
```
git clone https://github.com/okppop/url-shortener.git

cd url-shortener/
```

#### Import database schema:

Change host, username, database base your setting.
```
psql -h host -U username database < schema.sql
```

#### Compile:
```
go mod tidy

CGO_ENABLED=0 go build .
```

#### Edit config file:
```
cp config.yaml.example config.yaml
```
Then, use any editor edit "config.yaml".

#### Build container image:

```
docker build . -t url-shortener:1.0
```

#### Run container:

```
docker run -d --network host --name url url-shortener:1.0
```