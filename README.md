# Objective
Pipe the log to Lalamove's Kibana system

# Key generation
All the key will located at config/config.yaml
```sh
pwgen -s -n 16 3
```

# Format
```
{
  "level": "info"/* required */,
  "message": "I am an error message"/* required */,
  "time": "2018-09-19T07:45:46.215081910Z"/* required RFC3339 */,
  "src_file": "log-entry/main.go:19"/* required */,
  "src_line": "45"/* required */,
  "context": {
    "release": "1.1.1"/* required */,
    "locale": "zh_HK"/* required */,
    "location": "HK_HKG"/* required */,
    "lat": "123.100"/* required for mobile */,
    "lng": "-999.999"/* required for mobile */,
    "environment": "test"/* required e.g. ["dev", "test", "sandbox", "production"] */,
    "platform": "android"/* required e.g ["ios","android", "webapp"] */,
    "os": "android p"/* required for mobile */,
    "device": "s9"/* required for mobile */,
    "agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36"/* required for webapp */,
    "endpoint": "http://mobile-api/vanlogin"/* optional */,
    "footprint": "xxx-xxx-xxx"/* optional */,
    "client-id": "1234"/* optional */,
    "category": "backend"/* optional e.g. ["backend", "timeout", "bad gate way", "invalid format"] */
  }/* It is an optional field for extra information*/,
  "backtrace": "github.com/AlphaWong/log-entry/utilshttp.LogHandler\n\t/go/src/github.com/AlphaWong/log-entry/utilshttp/handler.go:45\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1964\nnet/http.(*ServeMux).ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2361\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2741\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1847"/* required log level >= error */
}
```