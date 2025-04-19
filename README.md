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
```
[Unit]
Description=My Go API Service

[Service]
ExecStart=/path/to/myapp
Restart=always
User=www-data

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