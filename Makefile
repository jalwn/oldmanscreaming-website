hello:
	echo "helloworld"

tailwind-watch: ## compile tailwindcss and watch for changes
	pnpm tailwindcss -i ./static/css/custom.css -o ./static/css/style.css --watch