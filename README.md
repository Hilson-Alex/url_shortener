![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![JS](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)

# SQLite URL Shortener

This project is a low config url shortener made for hobby. It uses a SQLite database hosted at an auto created .db file, so the only thing you need to do to run is... run it.


## Getting Started

By following these instructions you will get the server running on your local machine. 


### Installing

You can Take the Windows Bynary on the [latest tag][latest-tag] and execute, or run the following commands:

```shell
go install github.com/Hilson-Alex/url_shortener@latest
url_shortener
```

### Using

When running, the program will be serving on a port (usually localhost:8080).
It supplies some API endpoints for shortening and a small and simple frontend as well.

The server routes are the following:

#### GUI

- ***{host}/app***:
   
  Allows the User to create a short URL with an expire date between 1 and 30 days and shows the resulting shortened URL.

- ***{host}/app/list***:

  List all shortened URLs

- ***{host}/to/:key***:
  
  The actual short URL. Redirects the user for an URL based on the passed key

#### API

- ***{host}/short/create***

  Create a new short URL, it doesn't duplicate URLs in the database.
  In case of conflict, the longest expire date is kept on the database, but the user expire date is returned to not confuse the user.
  - Receives a JSON with:
    ```typescript
    {
      originalUrl: string, // The URL to be shortened
      expireDate: integer  // The number of DAYS to keep the URL active
    }
    ```
  - Returns:
    ```typescript
    {
      key: string,         //The resulting key for the short url
      originalUrl: string, // The URL to be shortened
      expireDate: integer, // The number of UNIX timestamp for the expire date
      shortUrl: string     // The shortened URL. {host}/to/:key
    }
    ```

- ***{host}/short/:key***
  
  Get the shortened URL without redirecting.
  - Returns: 
    ```typescript
    {
      key: string,         //The resulting key for the short url
      originalUrl: string, // The URL to be shortened
      expireDate: integer, // The number of UNIX timestamp for the expire date
      shortUrl: string     // The shortened URL. {host}/to/:key
    }
    ```

- ***{host}/short/list***
  
  Get All the Shortened URL.
  - Returns: 
    ```typescript
    [
      {
        key: string,         //The resulting key for the short url
        originalUrl: string, // The URL to be shortened
        expireDate: integer, // The number of UNIX timestamp for the expire date
        shortUrl: string     // The shortened URL. {host}/to/:key
      },
      // more
    ]
    ```

## Cloning and Building

To get the repository running locally on your machine you will need to install the [Go compiler](https://go.dev/dl/)

You can clone the project by running

```shell
git clone git@github.com:Hilson-Alex/url_shortener.git
```

or, if you can't use ssh

```shell
git clone https://github.com/Hilson-Alex/url_shortener.git
```

Next, open the project folder and run:

```shell
go run .
```

Or, you can build the binary instead:

```shell
go build
```

And then you will have a `url_shortener.exe` on the root folder. 

## Built With

- [Gin](https://gin-gonic.com/en/) was used to handle the server endpoints and build the API
- [ncrues/go-sqlite3](https://github.com/ncruces/go-sqlite3) was used as a GCO-free alternative for the sqlite database driver 

## Versioning

[Semantic Versioning](http://semver.org/) was used for versioning. For the versions
available, see the [tags on this repository](https://github.com/Hilson-Alex/url_shortener/releases/tags/).

## Author

- [Hilson A. W. Junior][Hilson-Alex] - *Initial work*


## License

This project is under the [GNU AGPLv3](https://choosealicense.com/licenses/agpl-3.0/) - for more info read the [LICENSE](LICENSE) file.

[Hilson-Alex]: https://github.com/Hilson-Alex
[latest-tag]: https://github.com/Hilson-Alex/url_shortener/releases/tag/v1.0.0
