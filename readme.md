# Building a REST API Backend with Golang by Echo Framework and caching with Redis 
[course](https://www.youtube.com/watch?v=6Dc0riyUYMQ&list=PLC4c48H3oDRw1827KV6GY8g887UC8usn-)

## Requirements
- Database: MySQL
- Caching server: Redis

## Configurations
0. Server
- Host: localhost
- Port: 8088
1. MySQL
- Read the *constant/constant.go* comments, can be customized
2. Redis
- Read the *constant/constant.go* commentsm can be customized

## How to run the code
1. Start the MySQL service (with above specifications)
2. Start the Redis service (with above specifications)
3. Run the code with the below command:
`go run user_manager/main.go`
