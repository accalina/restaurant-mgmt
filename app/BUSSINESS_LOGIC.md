# ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/entity` folder for describe business models
- `./app/model` folder for describe struct payload contract request and response
- `./app/repository` folder for describe queries for models
- `./app/service` folder for describe bussiness logic
- `./app.go` init app
