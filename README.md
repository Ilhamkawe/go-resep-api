<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h1>:)</h1>

<h3 align="center">GO-Rest-Api</h3>

  <p align="center">
   simple project to record and view food recipes
    <br />
    <a href="https://github.com/github_username/repo_name"><strong>Explore the docs »</strong></a>
  </p>
</div>





<!-- ABOUT THE PROJECT -->
## About The Project

<p align="center">
  <img src="https://i.ibb.co/F4jjTsg/microservice-diagram.jpg" alt="A4-1" border="0" />
</p>

In this project I created an API to record master data (ingredients and categories) along with recipe data

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* GIN
* GORM
* PotgreSQL
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started


### Prerequisites

you have to install postgreSQL and golang on your machine to run this project.


### Installation

1. Download and Install PostgreSQL at [https://postgresql.org](https://www.postgresql.org/download/)
2. Download and Install Go/Golang at [https://go.dev](https://go.dev/dl/)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Migration
on main.go change the configuration connection to the database, according to your device
```
dsn := "host=localhost user=postgres password=yourpassword dbname=go-resep port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
```
note:
you only need to create an empty database, because I have provided DB migration

## Usage

1. RUN golang project
   
   ```sh
   go run main.go
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## API Documentation

https://www.postman.com/flight-saganist-73393294/workspace/public-share/request/22067657-30097fd1-793f-4afb-8417-acdc407cd2b0

<!-- ROADMAP -->
## Roadmap

- [X] DB Migration
- [X] CRUD Master 
- [X] CRUD Resep
- [ ] Unit Test
- [ ] Dockerize

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Kawe - [@kawe123_](https://www.instagram.com/kawe123_/) - muhammad.ilham.kusumawardhana@gmail.com

Project Link: [https://github.com/Ilhamkawe/Simple-Microservice-project/](https://github.com/Ilhamkawe/Simple-Microservice-project/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
