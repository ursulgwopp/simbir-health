.PHONY: run stop

# File to store PIDs
PID_FILE = pids.txt

run:
	@echo "Starting account microservice..."
	@go run cmd/account_microservice/main.go & echo $$! > $(PID_FILE).account
	@echo "Starting hospital microservice..."
	@go run cmd/hospital_microservice/main.go & echo $$! > $(PID_FILE).hospital
	@cat $(PID_FILE).account $(PID_FILE).hospital > $(PID_FILE)
	@rm -f $(PID_FILE).account $(PID_FILE).hospital
	@echo "Services started with PIDs:"
	@cat $(PID_FILE)
	@wait

stop:
	@echo "Stopping services..."
	@if [ -f $(PID_FILE) ]; then \
		while read pid; do \
			echo "Stopping process $$pid..."; \
			kill $$pid || echo "Failed to stop process $$pid"; \
		done < $(PID_FILE); \
		rm -f $(PID_FILE); \
	else \
		echo "No PID file found. No services to stop."; \
	fi