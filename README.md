# 🚀 Quasar-Fire

## 🏁 How To Start

1. Install Golang
2. Clone this repository: `git clone https://github.com/franciscoruizar/quasar-fire`.
3. Read [Docs](https://github.com/franciscoruizar/quasar-fire/blob/main/docs/docs.md) & [API Documentation](https://documenter.getpostman.com/view/12160106/TzXtHKur)
4. Run the project: `make run`.
5. Start developing

## Deploy

### Tools & architecture
- Golang
- Gin-Gonic Server
- Hexagonal architecture
- AWS

### Endpoints
- Health-check       `GET` `http://localhost:8000/health`
- Top Secret         `POST` `http://localhost:8000/topsecret`
- Top Secret Split   `POST` `http://localhost:8000/topsecret_split/{satelite_name}`


### Commands

- Listing dependencies: `make dependencies`
- Update dependencies: `make update`
- Run tests: `make tests`
