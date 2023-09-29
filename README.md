# Social Network API &middot; ![MIT License](https://img.shields.io/badge/license-MIT-green) ![Go Version](https://img.shields.io/badge/go-%3E=1.17-blue) ![Potatowski Tag](https://img.shields.io/badge/potatowski-social%20network%20api-blue) 

This is a API (Application Programming Interface) to a social media or microblogging platform that allows users to post short messages with title based on Twitter


## Contributing

The main purpose of this repository is to continue evolving, making it faster and easier to use. Development of Social Network API happens in the open on GitHub, and I are grateful to the community for contributing bugfixes and improvements.

Has a issue list [here](https://github.com/potatowski/social-network-api/issues), you can add some need, functionality, or develop some of them, even refactor, just follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) in your commits

Follow these steps:
1. **Fork this Repository:** Click the "Fork" button in the top right corner of the page. This will create a copy of the repository in your GitHub account.
2. **Clone the Project Repository:** Clone the project repository to your local machine using Git. Navigate to the desired directory and run the following command
   ```bash
   git clone https://github.com/<your-username>/social-network-api.git
   ```
3. **Navigate to the Project Directory:** Change your current directory to the project folder using the cd command
   ```bash
   cd /social-network-api
   ```
4. **Create a New Branch:** create a new branch for your changes
   ```bash
   git checkout -b <change-type>/<what-will-be-worked-on>
   ```
   > Example:
   ```bash
   git checkout -b feat/post
   ```
5. **Submit a Pull Request:** push your branch to your repository on GitHub to dev branch

## Setting Up the Project for Development
To get started with the project and begin coding contributions, follow these steps:

1. **Install Docker:** Ensure Docker is installed on your system. If not, download and install it from the official [Docker website](https://www.docker.com/).
2. **Install API Testing Tools:** Choose and install a suitable API testing tool such as [Insomnia](https://insomnia.rest/download) or [Postman](https://www.postman.com/downloads/). These tools will help you test the project's APIs.
3. **Access the project folder:** Navigate to the directory where the project is
4. **Create env file:** Create .env.local file in folder **/app** and added all enviroment variables to run the application, has a exemple [here](/app/.env.example)
5. **Run Docker Composer:** Run the docker-compose up command to start the necessary containers and services specified in the docker-compose.yml file, the Docker Compose will handle the setup and orchestration of the project's dependencies and environment.
   ```bash
   docker-compose up
   ```

Now, you're all set to start working on the project and contribute to its development. Open the API testing tool (Insomnia or Postman) to interact with the project's APIs and begin coding!

## Changelog

Here [Changelog](CHANGELOG.md) to see latest updates and changes.

## License

This project is [MIT License](LICENSE)
