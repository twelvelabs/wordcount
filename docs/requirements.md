## API Outline

Using the language and stack of your choice, please create a secure web service that takes a block of text as input and returns the total count of words and a case-insensitive count of the occurrence of each word in the text. Counting punctuation is not required.

## Expected behavior

Given the input “Welcome to Austin!”
Produce this output: {“count”:3, “words”: {“welcome”: 1, “to”:1, “austin”:1}}

## Requirements

- The service is HTTPS only and answers on port 443.
- The service starts automatically on boot.
- The certificate and key are correctly installed.
- All API requests must be authenticated. (This service is not publicly available.)
- Requests and replies are JSON.
- All secrets are stored securely.
- The word counter can handle at least 2 megabytes of input text.
- Please provide secure access for us to log in and grade your server, including the ability to become root or "sudo su". This will not be held against you in terms of keeping the server secure.

## Not requirements

- Speed. A reasonable solution will be relatively fast, but does not need to be fully optimized.
- You do not need to trust our test certificate authority on your client machine.
