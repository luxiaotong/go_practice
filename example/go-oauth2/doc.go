package gooauth2

/*

## references

https://github.com/go-oauth2/oauth2
https://juejin.cn/post/7211805251216949305#heading-14
https://medium.com/@cyantarek/build-your-own-oauth2-server-in-go-7d0f660732c3
https://itnext.io/an-oauth-2-0-introduction-for-beginners-6e386b19f7a9


## client - get token by code

curl -X POST -b 'go_session_id=MjZiNjAxZDgtMTJkZS00YmExLTlmNmQtMWMzZmY2NjhlZjUy.9a30bdd94a922878db8101a4513cac30a1d095c2' \
 -d 'client_id=test&&grant_type=authorization_code&redirect_uri=http%3A%2F%2F127.0.0.1%3A9094%2Foauth2%2Fcallback' \
 -d 'code=NJFLOTQXMWUTN2YYZS0ZMTZMLTG0NWETMJNKOTKYMDLJZGJH' \
 http://127.0.0.1:9096/oauth2/token

*/
