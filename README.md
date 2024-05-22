# Password Generator
A simple web API designed to generate passwords

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ret0rn/password-generator?style=plastic)](https://github.com/ret0rn/password-generator/blob/main/go.mod) [![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fret0rn%2Fpassword-generator%2Fbadge%3Fref%3Dmain&style=plastic)](https://actions-badge.atrox.dev/ret0rn/password-generator/goto?ref=main) [![Dependabot](https://img.shields.io/badge/dependabot-eneble-brightgreen?style=plastic&logo=dependabot)](https://dependabot.com/) [![GitHub](https://img.shields.io/github/license/ret0rn/password-generator?style=plastic&color=brightgreen)](https://github.com/ret0rn/password-generator/blob/main/LICENCE)

### Setup

```
git clone https://github.com/ret0rn/password-generator.git
```


#### Docker
```
docker build --no-cache --rm -t password-generator .
docker run --rm -p 80:80 password-generator 
```

#### Makefile
```
make run
```

### Usage 

|Endpoint | Method | Params | Description|
|---------|--------|------------|--------|
|```/generate```| ```POST```| ```length: <uint>``` <br> ```symbols: <str>```| Generates a password of the specified ```length``` and including the specified ```symbols```


Result **with** symbols parameter


<img src="https://i.imgur.com/33ssQuy.jpg" alt="with symbols">


Result **without** symbols parameter

<img src="https://i.imgur.com/kNBeYPS.jpg" alt="with symbols">
