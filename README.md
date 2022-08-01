# Password Generator
A simple web API designed to generate passwords

### Insatll

#### Docker
```cmd
docker build -t password-generator .
docker run -p 80:80 password-generator
```



### Usage 

|Endpoint | Method | Params | Description|
|---------|--------|------------|--------|
|```/generate```| ```POST```| ```length: <uint>``` <br> ```symbols: <str>```| Generates a password of the specified ```length``` and including the specified ```symbols```


