# Premier League API

Create an API that serves the latest scores of fixtures of matches in a “**Mock Premier League**”

## User Types

There should be:

- **Admin accounts** which are used to
  - [ ] signup
  - [ ] login
  - [ ] Add teams
  - [ ] Remove teams
  - [ ] Edit Teams
  - [ ] View Teams 
    - [ ] Pagination
    - [ ] Sorting
    - [ ] Searching
  - [ ] create fixtures 
  - [ ] edit fixture
  - [ ] Remove fixture
  - [ ] View Fixture
    - [ ] Pagination
    - [ ] Sorting
    - [ ] Searching
  - [ ] Generate unique links for fixture
- **Users accounts** who can
  - [ ] signup
  - [ ] login
  - [ ] view teams
    - [ ] Pagination
    - [ ] Sorting
    - [ ] Searching
  - [ ] view completed fixtures
    - [ ] Pagination
    - [ ] Sorting
    - [ ] Searching
  - [ ] view pending fixtures
    - [ ] Pagination
    - [ ] Sorting
    - [ ] Searching
  - [ ] robustly search fixtures/teams
- [ ] Only the search API should be availble to the public.

## Authentication and Session Management
1. Use redis as your session store.
    - [ ] When a user logs in a session is created on mongoDB
    - [ ] The session is cached on Redis and first read there
    - [ ] if a session (JWT) is still valid and not on Redis, we check the DB
3. Authentication and Authorization for admin and user accounts should be done using `Bearer token` and `JWT`.
```js
{
    "role": "ADMIN | USER",
    "session": "uuid-uuid-uuid",
    "user": "uuid-uuid-uuid"
}
```

## Tools/Stack

- [ ] Golang
- [ ] MongoDB
- [ ] Redis
- [ ] Docker
- [ ] POSTMAN

## Tests

Unit tests are a must, submissions without tests will be ignored.

## Submission

1. Your API endpoints should be well documented in POSTMAN.
2. Seed the db with lots of data before final submission.
3. Code should be hosted on a git repository, Github preferably.
4. The API should be hosted on a live server (e.g. https://heroku.com)
5. Your app should be `containerized` using `docker`.
7. Write e2e tests for all user stories
8. Implement `web caching` API endpoints using `Redis`.
9. Implement `rate limiting` for user account API access.