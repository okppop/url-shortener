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
Executable file "url-shortener" was created.

#### Edit config file:
```
cp config.yaml.example config.yaml
```
Then, use any editor edit "config.yaml".

#### Import database table:

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

Add later.