
# To Run
- In one terminal run: `docker-compose up`
- In another termian: Use `curl` or Postman or something similar to hit the endpoints :
  `localhost:8081/rentals/{id}`
  `localhost:8081/rentals[?OptionalFilters]`
  - Examples with curl and jq
    ```bash
    curl "localhost:8081/rentals/3" | jq .
    curl "localhost:8081/rentals?price_min=9000&price_max=75000" | jq .
    curl "localhost:8081/rentals?limit=3&offset=6" | jq .
    curl "localhost:8081/rentals?ids=3,4,5" | jq .
    curl "localhost:8081/rentals?near=33.64,-117.93" | jq .
    curl "localhost:8081/rentals?sort=price" | jq .
    curl "localhost:8081/rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price" | jq .
    ```

# Test

The tests make heavy use of the `ginkgo` package.

- Set up postgres db with the provided data
  ```bash
  cd test/data
  docker-compose up
  ```
- Run ALL tests in another terminal
  ```bash
  # from the root project directory run:
  go test ./....
  ```
- To run package-specific tests:
  - `cd` to package and run `ginkgo test .`
  - Example:
    ```bash
    cd database 
    ginkgo test .
    ```



# Future features/changes

- Reach 100% code coverage. Currently:
  ```bash
  gooutdoorsy git:(master) go test --cover ./...
  ?       github.com/desteves/gooutdoorsy [no test files]
  ok      api       coverage: 76.6% of statements
  ok      database  coverage: 78.6% of statements
  ok      rental    coverage: 38.9% of statements
  ```
- Add versioning to the api `/v0/`
- Security (ssl, pwd files)
- Logging everywhere and generate docs
- Test ALL the filter parameters
- Fix bug with the `near` filter


----


# Coding Challenge

Thanks for applying for a backend role at Outdoorsy. We've put together this code challenge, which should take around 3-4 hours to complete.

## Functionality
The task is to develop a rentals JSON API that returns a list of rentals that can be filtered, sorted, and paginated. We have included files to create a database of rentals.

Your application should support the following endpoints.

- `/rentals/<RENTAL_ID>` Read one rental endpoint
- `/rentals` Read many (list) rentals endpoint
    - Supported query parameters
        - price_min (number)
        - price_max (number)
        - limit (number)
        - offset (number)
        - ids (comma separated list of rental ids)
        - near (comma separated pair [lat,lng])
        - sort (string)
    - Examples:
        - `rentals?price_min=9000&price_max=75000`
        - `rentals?limit=3&offset=6`
        - `rentals?ids=3,4,5`
        - `rentals?near=33.64,-117.93` // within 100 miles
        - `rentals?sort=price`
        - `rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price`

The rental object JSON in the response should have the following structure:
```json
{
  "id": "int",
  "name": "string",
  "description": "string",
  "type": "string",
  "make": "string",
  "model": "string",
  "year": "int",
  "length": "decimal",
  "sleeps": "int",
  "primary_image_url": "string",
  "price": {
    "day": "int"
  },
  "location": {
    "city": "string",
    "state": "string",
    "zip": "string",
    "country": "string",
    "lat": "decimal",
    "lng": "decimal"
  },
  "user": {
    "id": "int",
    "first_name": "string",
    "last_name": "string"
  }
}
```

## Notes
- Running `docker-compose up` will automatically generate a postgres database and some data to work with. Connect and use this database.
- Write production ready code.
- Please make frequent, and descriptive git commits.
- Use third-party libraries or not; your choice.
- Please use Golang to complete this task.
- Feel free to add functionality as you have time, but the feature described above is the priority.
- Please add tests

## What we're looking for
- The functionality of the project matches the description above
- An ability to think through all potential states
- In the README of the project, describe exactly how to run the application and execute the tests

When complete, please push your code to Github to your own account and send the link to the project or zip the project (including the `.git` directory) and send it back.

Thank you and please ask if you have any questions!
