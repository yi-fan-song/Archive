##
# Example for a reverse proxy to nodejs app with a backend
##

server {
	listen 443 ssl;

	server_name app.example.com www.app.example.com;

	location /api {
		return 302 /api/;
	}

	location /api/ {
		proxy_pass https://127.0.0.1:8081/; # trailing slash tells nginx to replace the matched part of uri with '/'
	}

	location / {
		proxy_pass https://127.0.0.1:8080;
	}
}