#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🧪 Construction System API - Test Suite${NC}"
echo -e "${BLUE}=======================================${NC}"

# Function to run a specific test type
run_test() {
    local test_type=$1
    local description=$2
    
    echo -e "\n${YELLOW}🚀 Running $description...${NC}"
    
    if docker compose -f docker-compose.test.yml run --rm $test_type; then
        echo -e "${GREEN}✅ $description completed successfully!${NC}"
        return 0
    else
        echo -e "${RED}❌ $description failed!${NC}"
        return 1
    fi
}

# Function to cleanup
cleanup() {
    echo -e "\n${YELLOW}🧹 Cleaning up test environment...${NC}"
    docker compose -f docker-compose.test.yml down -v
    docker system prune -f
}

# Trap to ensure cleanup on exit
trap cleanup EXIT

# Parse command line arguments
case "$1" in
    "unit")
        echo -e "${BLUE}Running Unit Tests Only${NC}"
        run_test "unit-tests" "Unit Tests"
        ;;
    "integration")
        echo -e "${BLUE}Running Integration Tests Only${NC}"
        run_test "integration-tests" "Integration Tests"
        ;;
    "api")
        echo -e "${BLUE}Running API Tests Only${NC}"
        # Start test infrastructure
        docker compose -f docker-compose.test.yml up -d postgres-test test-api
        run_test "api-tests" "API Tests"
        ;;
    "all"|"")
        echo -e "${BLUE}Running All Tests${NC}"
        
        # Start test infrastructure
        echo -e "\n${YELLOW}🏗️  Setting up test environment...${NC}"
        docker compose -f docker-compose.test.yml up -d postgres-test
        
        # Run unit tests
        if ! run_test "unit-tests" "Unit Tests"; then
            exit 1
        fi
        
        # Run integration tests
        if ! run_test "integration-tests" "Integration Tests"; then
            exit 1
        fi
        
        # Start API for API tests
        docker compose -f docker-compose.test.yml up -d test-api
        
        # Run API tests
        if ! run_test "api-tests" "API Tests"; then
            exit 1
        fi
        
        echo -e "\n${GREEN}🎉 All tests completed successfully!${NC}"
        ;;
    "coverage")
        echo -e "${BLUE}Running Tests with Coverage Report${NC}"
        docker compose -f docker-compose.test.yml up -d postgres-test

        # Run tests with coverage
        docker compose -f docker-compose.test.yml run --rm unit-tests sh -c "
            go test -v -race -coverprofile=coverage.out -covermode=atomic ./... &&
            go tool cover -html=coverage.out -o coverage.html &&
            go tool cover -func=coverage.out
        "
        
        echo -e "${GREEN}📊 Coverage report generated: coverage.html${NC}"
        ;;
    "help"|"-h"|"--help")
        echo -e "${BLUE}Usage: $0 [OPTION]${NC}"
        echo -e "  unit        Run unit tests only"
        echo -e "  integration Run integration tests only"
        echo -e "  api         Run API tests only"
        echo -e "  all         Run all tests (default)"
        echo -e "  coverage    Run tests with coverage report"
        echo -e "  help        Show this help message"
        ;;
    *)
        echo -e "${RED}❌ Unknown option: $1${NC}"
        echo -e "Use '$0 help' for usage information"
        exit 1
        ;;
esac
