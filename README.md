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
    "message": "", 
    "src_file": "",
    "src_line": "",
    "context": {
      "a": "abc",
      "c": "123",
    },
    "level": "", 
    "time": "",
    "backtrace": ""
}
```