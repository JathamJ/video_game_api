Name: video_game_api-api
Host: 0.0.0.0
Port: 8101

Timeout: 3000 # 3s
MaxConns: 1000 # 最大并发连接数，默认值为 1000 qps
MaxBytes: 3145728 #3M


#Log
Log:
ServiceName: video_game_api-api
Level: debug
Mode: console
Encoding: plain

#DB
DB:
DataSource: root:root@tcp(127.0.0.1:3306)/aigc?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai