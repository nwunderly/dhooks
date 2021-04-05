package main

const ExampleConfig = `
wh-url: webhook url
wh-username: default username for messages
wh-avatar-url: default avatar url for messages

messages:
  # message 1
  - username: username
    avatar-url: avatar url
    content: message content
    file: /path/to/file.png

  # message 2
  - content: message content
    embed:
      color: 00ffaf
      title: embed title
      description: embed description
      fields:
        - name: name
          value: value
        - name: name
          value: value
`
