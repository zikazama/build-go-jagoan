# build-go-jagoan

## How to run

```
go run main.go
```

## How to build

```
go build
```

## Curl Script

Create
```
curl -X POST http://localhost:8080/product -d '{"id": "3", "name": "Tablet", "price": 400.00}' -H "Content-Type: application/json"
```

Read
```
curl http://localhost:8080/products
```

Delete
```
curl -X DELETE "http://localhost:8080/product/delete?id=3"
```

## Example .service

Simpan file ini sebagai go-api.service dan salin ke direktori /etc/systemd/system/ di VPS.

```
[Unit]
Description=Go API App

[Service]
ExecStart=./build-go-jagoan/main.go
Restart=always
User=root

[Install]
WantedBy=multi-user.target
```

## Cara install nginx

1. Update System

```
sudo apt update
sudo apt upgrade -y
```

2. Install nginx

```
sudo apt install nginx -y
```

3. Check status nginx

```
sudo systemctl status nginx
```

4. Start

```
sudo systemctl start nginx
```

5. Enable nginx

```
sudo systemctl enable nginx
```

6. Firewall

```
sudo ufw allow 'Nginx Full'
```

## Basic Config nginx

```
server {
    listen 80;
    
    # Gunakan alamat IP VPS sebagai server_name
    server_name 123.45.67.89;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```