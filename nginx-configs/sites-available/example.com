##
# Example for serving static files
##

server {
	listen 443 ssl;

	server_name example.com www.example.com;

	root /var/www/example.com;

	location / {
		# First attempt to serve request as file, then
		# as directory, then fall back to displaying a 404.
		try_files $uri $uri/ =404;
	}
}