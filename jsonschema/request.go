package jsonschema

var RequestSchema = `
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "message": {
      "type": "string",
      "pattern": "^\\S+"
    },
    "src_file": {
      "type": "string",
      "pattern": "^\\S+"
    },
    "src_line": {
      "type": "string",
      "pattern": "^\\d+$"
    },
    "context": {
      "type": "object"
    },
    "level": {
      "type": "string",
      "enum": [
        "debug",
        "info",
        "warning",
        "error",
        "fatal"
      ]
    },
    "time": {
      "type": "string",
      "pattern": "^\\d{4}-[01]\\d-[0-3]\\dT[0-2]\\d:[0-5]\\d:[0-5]\\d(.\\d+)?Z$"
    },
    "backtrace": {
      "type": "string",
      "pattern": "^\\S+"
    }
  },
  "required": [
    "message",
    "src_file",
    "src_line",
    "level",
    "time"
  ],
  "additionalProperties": false
}
`
