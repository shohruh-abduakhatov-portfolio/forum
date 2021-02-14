# [Forum](https://github.com/01-edu/public/tree/master/subjects/forum)

Forum project


# Milestone

### Template:
- [x] set template
- [ ] 404 page

### SQLite:
- [x] Entity relationship diagram
- [x] queries
- [x] create index for double PKs
- [x] delete on cascade

### Authentication:
- [x] register
- [x] via unique email
- [x] username
- [x] password(encrypted)

- [x] session
- [x] cookies to allow each user to have only one opened session
- [ ] forgot password

<b>[Bonus:](https://github.com/01-edu/public/tree/master/subjects/forum/authentication)</b>

- [x] Github 
- [ ] Google or/and FB auth


### Communication:
- [x] registered users create posts and comments
- [x] associate one or more categories to post created
- [x] View only: unregistered users


### Likes and Dislikes:
- [x] number of likes/dislikes
- [x] only registered users
- [x] put like and dislike

### Filter posts by:
- [x] categories
- [x] created posts by user // api ready
- [x] liked posts by user // api ready

### Docker:
- [ ] golang
- [ ] sqlite3



## [forum-advanced-features:](https://github.com/01-edu/public/tree/master/subjects/forum/advanced-features)
- [ ] notify users when posts are:
    - liked/disliked
    - commented

-  page that
- [ ] Shows the user created posts
- [ ] Shows where the user left a like or a dislike
- [ ] Shows where and what the user has been commenting (For this, the comment will have to be shown, as well as the post commented
)
- [ ] You have to create a section where you will be able to Edit/Remove posts and comments


## [forum-image-upload:](https://github.com/01-edu/public/tree/master/subjects/forum/image-upload)
- [ ] supports at least: JPEG, PNG and GIF (max 20 mb)


## [forum-moderation](https://github.com/01-edu/public/tree/master/subjects/forum/moderation)

- [ ] 


## [forum-security:](https://github.com/01-edu/public/tree/master/subjects/forum/security)

- [ ] implement a Hypertext Transfer Protocol Secure (HTTPS) protocol
    - generate an SSL certificate or use "Certificate Authorities"(CA's)
- [ ] Rate Limiting
- bonus:
- [ ] test ddos


check others




## Links
- [ERD](https://online.visual-paradigm.com/w/dkhhvphc/diagrams/#proj=0&type=ERDiagram)
- [Sessions](https://astaxie.gitbooks.io/build-web-application-with-golang/en/06.2.html)
- [source code](https://github.com/astaxie/build-web-application-with-golang/tree/master/en/code/src/apps/ch.5.2)
- [routing golang](https://github.com/golang-samples/template/blob/master/extends/main.go)

## [Templates](https://www.creative-tim.com/templates/free?page=5):

- [toolkit](https://demos.creative-tim.com/material-kit/index.html?_ga=2.136642038.797548910.1593860822-1386350353.1593860822)

- [aragon](https://www.creative-tim.com/product/argon-design-system) | [docs](https://demos.creative-tim.com/argon-design-system/docs/components/navbar.html#) | [demo](https://demos.creative-tim.com/argon-design-system/index.html)

- [lib](https://demos.creative-tim.com/material-kit/docs/2.1/getting-started/introduction.html)

- [sign in/up](https://www.creative-tim.com/product/login-and-register-modal)

- [sqlite example](https://github.com/fevziomurtekin/go-sqlite/blob/master/sqlite.go)

- [oauth golang](https://github.com/andrewtian/golang-github-oauth-example) | [2](https://github.com/sohamkamani/go-oauth-example) | [blog](https://www.sohamkamani.com/golang/2018-06-24-oauth-with-golang/) | [lib](https://github.com/gosimple/oauth2)

- [Rate Limiting](https://www.alexedwards.net/blog/how-to-rate-limit-http-requests) | 

- [SSL server example](https://gist.github.com/samthor/5ff8cfac1f80b03dfe5a9be62b29d7f2)