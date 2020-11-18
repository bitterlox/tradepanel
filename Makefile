W_PID=0
.PHONY: dev frontend wails

dev:
	@stty -echoctl

	cd client/frontend && npm run serve > /dev/null 2>&1 & N_PID=$$!

	cd client && wails serve > /dev/null 2>&1 & W_PID=$$!

	@trap "pkill $$N_PID $$W_PID && echo 'killed $$N_PID $$W_PID'" INT

	@echo "\nStarted wails bridge; serving frontend on port 8080"
	@echo "send interrupt to kill dev process"

	@sleep infinity
