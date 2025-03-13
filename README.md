# sakuga-go-sdk

![test-status](https://github.com/ctiller15/sakuga-go-sdk/actions/workflows/tests.yml/badge.svg)

An unofficial go SDK for the [Sakugabooru API](https://sakugabooru.com/help/api)

## Installation
```bash
go get github.com/ctiller15/sakuga-go-sdk@latest
```

## Usage
```go
	api := sakugaapi.NewAPI()

	opts := sakugamodels.PostsListOptions{
		Limit: 10,
	}
	apiResult, err := api.Posts.List(&opts)

	if err != nil {
		log.Fatal(err)
	}
    // Do something with the data
```

Currently supported api routes:

1. Posts - List
1. Tags - List
1. Tags - Related
1. Artists - List
1. Comments - Show
1. Wiki - List
1. Notes - List
1. Notes - Search
1. Notes - History
1. Users - Search
1. Forum - List
1. Pools - List Pools
1. Pools - List Posts
1. Favorites - List Users

Sakugabooru is a fantastic resource for animation. If you like their site, this SDK, (which relies entirely on their good graces), or animation in general it would be awesome if you [supported their patreon](https://www.patreon.com/Sakugabooru).

Future Plans:
- Add intentional support for [meta-tags](https://sakugabooru.com/help/cheatsheet)