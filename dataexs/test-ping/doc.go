package main

/*
docker run -itd --sysctl net.ipv4.ping_group_range="0 2147483647" \
--name "test-ping" --network custom \
-v /Users/luxiaotong/code/go_practice/dataexs/test-ping:/app \
-w /app alpine:latest /app/test-ping.app
*/
