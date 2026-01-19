.PHONY: help install-tools install generate build run dev clean test fmt lint all css-build css-watch

# Default target
.DEFAULT_GOAL := help

# Variables
BINARY_NAME=resume-app
MAIN_PATH=.
GO=go
TEMPL=templ
AIR=air
NPM=npm

# Colors for output
COLOR_RESET=\033[0m
COLOR_BOLD=\033[1m
COLOR_GREEN=\033[32m
COLOR_YELLOW=\033[33m
COLOR_BLUE=\033[34m

help: ## Show this help message
	@echo "$(COLOR_BOLD)$(COLOR_BLUE)Go Resume Builder - Makefile Commands$(COLOR_RESET)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "$(COLOR_GREEN)%-20s$(COLOR_RESET) %s\n", $1, $2}'
	@echo ""

install-tools: ## Install required tools (templ, air, tailwindcss)
	@echo "$(COLOR_YELLOW)Installing Templ...$(COLOR_RESET)"
	@$(GO) install github.com/a-h/templ/cmd/templ@latest
	@echo "$(COLOR_YELLOW)Installing Air...$(COLOR_RESET)"
	@$(GO) install github.com/cosmtrek/air@latest
	@echo "$(COLOR_YELLOW)Installing TailwindCSS...$(COLOR_RESET)"
	@$(NPM) install
	@echo "$(COLOR_GREEN)✓ Tools installed successfully!$(COLOR_RESET)"

install: ## Install Go dependencies
	@echo "$(COLOR_YELLOW)Installing dependencies...$(COLOR_RESET)"
	@$(GO) mod download
	@$(GO) mod tidy
	@echo "$(COLOR_GREEN)✓ Dependencies installed!$(COLOR_RESET)"

css-build: ## Build TailwindCSS for production
	@echo "$(COLOR_YELLOW)Building TailwindCSS...$(COLOR_RESET)"
	@$(NPM) run build:css
	@echo "$(COLOR_GREEN)✓ CSS built!$(COLOR_RESET)"

css-watch: ## Watch and rebuild TailwindCSS
	@echo "$(COLOR_YELLOW)Watching TailwindCSS...$(COLOR_RESET)"
	@$(NPM) run watch:css

generate: ## Generate Templ templates
	@echo "$(COLOR_YELLOW)Generating Templ templates...$(COLOR_RESET)"
	@$(TEMPL) generate
	@echo "$(COLOR_GREEN)✓ Templates generated!$(COLOR_RESET)"

build: css-build generate ## Build the application
	@echo "$(COLOR_YELLOW)Building $(BINARY_NAME)...$(COLOR_RESET)"
	@$(GO) build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "$(COLOR_GREEN)✓ Build complete: ./$(BINARY_NAME)$(COLOR_RESET)"

run: css-build generate ## Run the application
	@echo "$(COLOR_YELLOW)Running application...$(COLOR_RESET)"
	@$(GO) run $(MAIN_PATH)

dev: css-build generate ## Run with Air (hot reload)
	@echo "$(COLOR_YELLOW)Starting development server with hot reload...$(COLOR_RESET)"
	@if command -v $(AIR) > /dev/null; then \
		$(AIR); \
	else \
		echo "$(COLOR_YELLOW)Air not found. Installing...$(COLOR_RESET)"; \
		$(GO) install github.com/cosmtrek/air@latest; \
		$(AIR); \
	fi

clean: ## Clean build artifacts
	@echo "$(COLOR_YELLOW)Cleaning build artifacts...$(COLOR_RESET)"
	@rm -f $(BINARY_NAME)
	@rm -rf tmp/
	@rm -f static/output.css
	@find . -name "*_templ.go" -type f -delete
	@echo "$(COLOR_GREEN)✓ Clean complete!$(COLOR_RESET)"

test: ## Run tests
	@echo "$(COLOR_YELLOW)Running tests...$(COLOR_RESET)"
	@$(GO) test -v ./...

fmt: ## Format code
	@echo "$(COLOR_YELLOW)Formatting code...$(COLOR_RESET)"
	@$(GO) fmt ./...
	@$(TEMPL) fmt .
	@echo "$(COLOR_GREEN)✓ Code formatted!$(COLOR_RESET)"

lint: ## Run linter
	@echo "$(COLOR_YELLOW)Running linter...$(COLOR_RESET)"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "$(COLOR_YELLOW)golangci-lint not found. Install it from https://golangci-lint.run/$(COLOR_RESET)"; \
	fi

watch: dev ## Alias for dev (watch mode)

setup: install-tools install css-build generate ## Complete setup (install everything)
	@echo "$(COLOR_GREEN)✓ Setup complete! Run 'make dev' to start developing.$(COLOR_RESET)"

docker-build: ## Build Docker image
	@echo "$(COLOR_YELLOW)Building Docker image...$(COLOR_RESET)"
	@docker build -t $(BINARY_NAME) .
	@echo "$(COLOR_GREEN)✓ Docker image built!$(COLOR_RESET)"

docker-run: ## Run Docker container
	@echo "$(COLOR_YELLOW)Running Docker container...$(COLOR_RESET)"
	@docker run -p 8080:8080 $(BINARY_NAME)

all: clean install css-build generate build ## Clean, install, generate, and build
	@echo "$(COLOR_GREEN)✓ All tasks complete!$(COLOR_RESET)"