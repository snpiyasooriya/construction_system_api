#!/bin/bash

# Simple test script for Docker Compose V2
# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🧪 Simple Schedule API Test${NC}"
echo -e "${BLUE}===========================${NC}"

# Function to cleanup
cleanup() {
    echo -e "\n${YELLOW}🧹 Cleaning up...${NC}"
    docker compose -f docker-compose.test.yml down -v 2>/dev/null || true
}

# Trap to ensure cleanup on exit
trap cleanup EXIT

echo -e "\n${YELLOW}🏗️  Starting test database...${NC}"
if docker compose -f docker-compose.test.yml up -d postgres-test; then
    echo -e "${GREEN}✅ Test database started${NC}"
else
    echo -e "${RED}❌ Failed to start test database${NC}"
    exit 1
fi

echo -e "\n${YELLOW}⏳ Waiting for database to be ready...${NC}"
sleep 10

echo -e "\n${YELLOW}🚀 Running unit tests...${NC}"
if docker compose -f docker-compose.test.yml run --rm unit-tests; then
    echo -e "${GREEN}✅ Unit tests completed successfully!${NC}"
else
    echo -e "${RED}❌ Unit tests failed!${NC}"
    exit 1
fi

echo -e "\n${GREEN}🎉 Tests completed successfully!${NC}"
