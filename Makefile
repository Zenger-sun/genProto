proto:
  protoc --go_out=. proto/*.proto
  
help: Makefile
	@echo " Choose a command run:"
	@sed -n 's/^##://p' $< | column -t -s ':' | sed -e 's/^/ /'
