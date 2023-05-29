# Online-Learning-Platform

<a name="readme-top"></a>



[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<div align="center">
  <a href="https://github.com/mdanys/Online-Learning-Platform">
    <img src="utils/logo.png" alt="Logo" width="200px">
  </a>

<h3 align="center">Let's Learning!</h3>

  <p align="center">
    This project is a API for Online-Learning-Platform.
    <br />
    <br />
    <br />
    <a href="https://app.swaggerhub.com/apis-docs/mdanys/online-learning-platform/1.0.0">View Demo</a>
    ·
    <a href="https://github.com/mdanys/Online-Learning-Platform/issues">Report Bug</a>
    ·
    <a href="https://github.com/mdanys/Online-Learning-Platform/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#features">Features</a>
      <ul>
        <li><a href="#entity-relationship-diagram">Entity Relationship Diagram</a></li>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## Features

- Users
    - Users can register
    - Users can log in
    - Users can get categories
    - Users can get courses
    - Users can make transaction
- Admin
    - Admin can delete users
    - Admin can add and edit category
    - Admin can add, edit, and delete course
    - Admin can get simple statistics

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Entity Relationship Diagram



```erd
[category] {bgcolor: "#d0f0d0"}
  *id {label: "bigint(20), auto_increment"}
  *name {label: "varchar(255), not null"}
  *dtm_crt {label: "datetime, not null"}
  *dtm_upd {label: "datetime, not null"}

[course] {bgcolor: "#f0b0b0"}
  *id {label: "bigint(20), auto_increment"}
  *category_id {label: "bigint(20), not null"}
  *name {label: "varchar(255), not null"}
  *detail {label: "text"}
  *price {label: "decimal(10,2), not null"}
  *picture {label: "varchar(1024)"}
  *dtm_crt {label: "datetime, not null"}
  *dtm_upd {label: "datetime, not null"}

[transaction] {bgcolor: "#b0d0f0"}
  *id {label: "bigint(20), auto_increment"}
  *user_id {label: "bigint(20), not null"}
  *category_id {label: "bigint(20), not null"}
  *price {label: "decimal(10,2), not null"}
  *dtm_crt {label: "datetime, not null"}
  *dtm_upd {label: "datetime, not null"}

[user] {bgcolor: "#b0b0f0"}
  *id {label: "bigint(20), auto_increment"}
  *email {label: "varchar(255), not null"}
  *password {label: "varchar(255), not null"}
  *name {label: "varchar(255), not null"}
  *picture {label: "varchar(1024)"}
  *role {label: "varchar(255), not null"}
  *deleted {label: "bool, not null"}
  *dtm_crt {label: "datetime, not null"}
  *dtm_upd {label: "datetime, not null"}

course *--1 category {label: "category_id <--- id"}
transaction *--1 user {label: "user_id <--- id"}
transaction *--1 course {label: "course_id <--- id"}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation local

1. Clone the repo
   ```bash
   git clone https://github.com/mdanys/Online-Learning-Platform
   ```
2. Get env at [config.env](https://drive.google.com/file/d/13wLy-4LO1EPOmMTaaCZFWr7fsc2_uNYz/view?usp=sharing)
3. Enter your config in `config.env`
   ```bash
   AWS_REGION = "ENTER YOUR AWS REGION"
   AWS_ACCESS_KEY_ID = "ENTER YOUR AWS ACCESS KEY ID"
   AWS_SECRET_ACCESS_KEY = "ENTER YOUR AWS SECRET KEY ID"
   ```
4. Run project
   ```bash
   cd cd Online-Learning-Platform
   paste your config.env
   go run .
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Documentation

_For more examples, please refer to the [OPEN API](https://app.swaggerhub.com/apis-docs/mdanys/online-learning-platform/1.0.0)_

[![GunTour-API][product-screenshot]](https://github.com/GunTour/Back-End/blob/main/utils/online-learning-platform



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/GunTour/Back-End.svg?style=for-the-badge
[contributors-url]: https://github.com/GunTour/Back-End/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/GunTour/Back-End.svg?style=for-the-badge
[forks-url]: https://github.com/GunTour/Back-End/network/members
[stars-shield]: https://img.shields.io/github/stars/GunTour/Back-End.svg?style=for-the-badge
[stars-url]: https://github.com/GunTour/Back-End/stargazers
[issues-shield]: https://img.shields.io/github/issues/GunTour/Back-End.svg?style=for-the-badge
[issues-url]: https://github.com/GunTour/Back-End/issues
[license-shield]: https://img.shields.io/github/license/GunTour/Back-End.svg?style=for-the-badge
[license-url]: https://github.com/GunTour/Back-End/blob/main/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url-1]: https://linkedin.com/in/khalidrianda
[linkedin-url-2]: https://linkedin.com/in/mochammaddany
[product-screenshot]: readme/GunTour.gif
[erd-screenshot]: readme/erd.jpg
[Go]: https://img.shields.io/github/go-mod/go-version/GunTour/Back-End
[go-url]: https://go.dev/
[Echo]: https://img.shields.io/badge/Echo-v4-9cf
[echo-url]: https://echo.labstack.com/
[Oauth]: https://img.shields.io/badge/OAuth-Google-informational
[oauth-url]: https://developers.google.com/identity/protocols/oauth2
[Gmail]: https://img.shields.io/badge/Gmail-Google-informational
[mail-url]: https://github.com/googleapis/google-api-go-client
[Calendar]: https://img.shields.io/badge/Calender-Google-informational
[calendar-url]: https://github.com/googleapis/google-api-go-client
[AWS]: https://img.shields.io/badge/AWS-EC2-orange
[aws-url]: https://aws.amazon.com/
[khalid]: https://img.shields.io/badge/-Khalid-black.svg?style=for-the-badge&logo=Khalid&colorB=555
[dany]: https://img.shields.io/badge/-Dany-black.svg?style=for-the-badge&logo=Dany&colorB=555
[khalid-url]: https://github.com/khalidrianda
[dany-url]: https://github.com/mdanys
[email-shield]: https://img.shields.io/badge/gmail-DD0031?style=for-the-badge&logo=gmail&logoColor=white
[email-1]: khalidrianda12@gmail.com
[email-2]: mochammaddany@gmail.com
