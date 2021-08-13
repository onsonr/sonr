[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/sonr-io/core">
    <img src="https://uploads-ssl.webflow.com/60e4b57e5960f8d0456720e7/60fbc0e3fcdf204c7ed9946b_Github%20-%20Core.png" alt="Logo" height="275">
  </a>



  <p align="center">
  Core Framework that manages the Sonr Libp2p node in Go, Handles File Management, Connection to Peer, and Pub-Sub for Lobby.
    <a href="https://github.com/sonr-io/core"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/sonr-io/core">View Demo</a>
    ·
    <a href="https://github.com/sonr-io/core/issues">Report Bug</a>
    ·
    <a href="https://github.com/sonr-io/core/issues">Request Feature</a>
  </p>
</p>
<br />


[![CodeFactor](https://www.codefactor.io/repository/github/sonr-io/core/badge/release?s=ee02a1b599502678b3d583aa5b6d1f55d2137ded)](https://www.codefactor.io/repository/github/sonr-io/core/overview/release)      [![Bind Workflow](https://github.com/sonr-io/core/actions/workflows/bind-lib.yml/badge.svg)](https://github.com/sonr-io/core/actions/workflows/bind-lib.yml)      [![Build Workflow](https://github.com/sonr-io/core/actions/workflows/build-lib.yml/badge.svg)](https://github.com/sonr-io/core/actions/workflows/build-lib.yml)


*By [Sonr](https://www.sonr.io), creators of [The Sonr App](https://www.twitter.com/TheSonrApp)*

---

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)



### Built With

* [Golang]()




<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/sonr-io/core.git
   ```
2. Install NPM packages
   ```sh
   npm install
   ```



<!-- USAGE EXAMPLES -->
## Usage

This project contains a `makefile` with the following commands:
```bash
# Binds Android and iOS for Plugin Path
make bind

# Binds iOS Framework ONLY
make bind.ios

# Binds AAR for Android ONLY
make bind.android

# Compiles Protobuf models for Core Library and Plugin
make proto

# Binds Binary, Creates Protobufs, and Updates App
make upgrade

# Reinitializes Gomobile and Removes Framworks from Plugin
make clean
```

_For more examples, please refer to the [Documentation](https://example.com)_



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/sonr-io/core/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Prad Nukala - [TheSonrApp](https://twitter.com/TheSonrApp) - pradn@sonr.io

Project Link: [Github](https://github.com/sonr-io/core) - [Discord](https://github.com/sonr-io/core) - [Website](https://github.com/sonr-io/core)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Libp2p]()
* [Textile]()
* [HNS]()
* [Flutter]()





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/sonr-io/core.svg?style=for-the-badge
[contributors-url]: https://github.com/sonr-io/core/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/sonr-io/core.svg?style=for-the-badge
[forks-url]: https://github.com/sonr-io/core/network/members
[stars-shield]: https://img.shields.io/github/stars/sonr-io/core.svg?style=for-the-badge
[stars-url]: https://github.com/sonr-io/core/stargazers
[issues-shield]: https://img.shields.io/github/issues/sonr-io/core.svg?style=for-the-badge
[issues-url]: https://github.com/sonr-io/core/issues
[license-shield]: https://img.shields.io/github/license/sonr-io/core.svg?style=for-the-badge
[license-url]: https://github.com/sonr-io/core/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/sonr-io
