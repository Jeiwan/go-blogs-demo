# Blogs
An example multi-tenancy application in Golang.

Usage:
```
docker-compose up
go run main.go ui
```

## Version: 1.0.0

### Security
**create_post_auth**  

|basic|*Basic*|
|---|---|

### /blog

#### POST
##### Summary:

Create a new blog

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| blog | body | Blog name and password. | No | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | success |  |
| 401 | bad request | [Error](#error) |

### /blog/{blog_name}

#### GET
##### Summary:

View posts in a blog

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| blog_name | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | success | [Posts](#posts) |
| 404 | not found | [Error](#error) |

### /blog/{blog_name}/posts

#### POST
##### Summary:

Publish a post to a blog

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| blog_name | path |  | Yes | string |
| post | body |  | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | success | object |
| 401 | bad request | [Error](#error) |
| 404 | not found | [Error](#error) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| create_post_auth | |

### Models


#### Post

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| date | string |  | No |
| post | string |  | No |

#### Posts

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| posts | [ [Post](#post) ] |  | No |

#### Error

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |