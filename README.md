# DND Beyond Proxy
This is a simple proxy designed for use with https://github.com/PanoramicPanda/ddb-party-view

## Usage
By default, the only requirement is to define `PORT` which is the port that the proxy will listen on.

## Security
There is limited security to help prevent your proxy from being abused. Specifically, you can specify a list
of character ID's in comma delimited format as the `ALLOWED_CHARACTER_IDS` variable which will only allow 
those characters to be retrieved.

```ALLOWED_CHARACTER_IDS=12343124,145901435,1345981029```
