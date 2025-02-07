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
- Go Development environment
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
executable file "url-shortener" was created.

#### Edit config file:
```
cp config.yaml.example config.yaml
```
then, use any editor edit "config.yaml".

#### Import database table:

Change host, username, database base your setting.
```
psql -h host -U username database < schema.sql
```

#### Deployment:
```
sudo mkdir -p /usr/local/url
sudo mv url-shortener config.yaml /usr/local/url/
sudo -i cp url.service /etc/systemd/system/
sudo systemctl daemon-reload
```

#### Start service:
```
sudo systemctl start url.service
```

### Container Deployment

Add later.