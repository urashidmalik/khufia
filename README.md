#Purpose
Utility to quickly encrypt and decrpt files without OpenSSL

## Usage

###### Prereq
- Make sure you have make installed and go installed

```sh
# Install go on mac
brew install go
# Install make on mac
brew install make
```

###### Build

```sh
make build-mac
```

###### Encrypt

```sh
./bin/khufia-darwin-amd64 encrypt --password MYPASS --in /path/to/encrypted/file --out /path/to/decrypted/file
```

###### Decrypt

```sh
./bin/khufia-darwin-amd64 decrypt --password MYPASS --in /path/to/unencrypted/file --out /path/to/encrypted/file
```
