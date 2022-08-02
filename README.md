# Password Generator
A simple web API designed to generate passwords

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ret0rn/password-generator?style=plastic)](https://github.com/ret0rn/password-generator/blob/main/go.mod) [![GitHub](https://img.shields.io/github/license/ret0rn/password-generator?style=plastic)](https://github.com/ret0rn/password-generator/blob/main/LICENCE)

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


