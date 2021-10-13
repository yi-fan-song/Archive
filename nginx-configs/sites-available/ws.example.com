# This is an example of using nginx to route websocket connections

map $http_upgrade $connection_upgrade {
        default upgrade;
        '' close;
}

server {
        listen 3000 ws.example.com;
        server_name ws.example.com;
        
        location / {
                proxy_pass http://127.0.0.1:3001;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection $connection_upgrade;
                proxy_set_header Host $host;
        }
}
