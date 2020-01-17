# Test Health route
curl http://localhost:9999

# Test Invalid data
curl -X POST http://localhost:9999/calc -H "Content-Type: application/txt" -d '{operation": "sum", "operand1": 1, "operand2": 3}'

# Test Sum operation
curl -X POST http://localhost:9999/calc -H "Content-Type: application/json" -d '{"operation": "sum", "operand1": 1, "operand2": 3}'

# Test Subtract Endpoint
curl -X POST http://localhost:9999/calc -H "Content-Type: application/json" -d '{"operation": "subtract", "operand1": 5, "operand2": 3}'
