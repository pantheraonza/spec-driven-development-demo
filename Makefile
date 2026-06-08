.PHONY: help setup backend frontend test clean

help:
	@echo "Spec-Driven Development Demo - Makefile"
	@echo "======================================="
	@echo ""
	@echo "Targets:"
	@echo "  setup          - Setup inicial del proyecto"
	@echo "  backend        - Ejecutar backend"
	@echo "  frontend       - Ejecutar frontend"
	@echo "  test           - Ejecutar tests"
	@echo "  clean          - Limpiar artefactos"

setup:
	@echo "Setup inicial..."
	@cd backend && go mod download
	@cd frontend && npm install

backend:
	@echo "Iniciando backend..."
	@cd backend && make run

frontend:
	@echo "Iniciando frontend..."
	@cd frontend && npm start

test:
	@echo "Ejecutando tests..."
	@cd backend && make test
	@cd frontend && npm test -- --watch=false

clean:
	@echo "Limpiando..."
	@cd backend && make clean
	@rm -rf frontend/node_modules frontend/build
