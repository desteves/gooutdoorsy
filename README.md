# Coding Challenge

Thanks for applying for a backend role at Outdoorsy. We've put together this code challenge, which should take around 2-3 hours to complete.

## Functionality
The task is to develop a rentals JSON API that returns a list of campervans that can be filtered, sorted, and paginated. We have included files to create a database of rentals.

Your application should support the following endpoints and query parameters in any combination or order.

- `rentals`
    - `rentals?price[min]=9000&price[max]=75000`
    - `rentals?page[limit]=3&page[offset]=6`
    - `rentals?ids=3,4,5`
    - `rentals?near=33.64,-117.93` // within 100 miles
    - `rentals?sort=price`
    - `rentals?near=33.64,-117.93&price[min]=9000&price[max]=75000&page[limit]=3&page[offset]=6&sort=price`
- `rentals/<RENTAL_ID>`


## Notes
- Running `docker-compose up` will automatically generate a database and some data to work with. Connect and use this database.
- Write production ready code.
- Please make frequent, and descriptive git commits.
- Use third-party libraries or not; your choice.
- Our preference would be that you would use Golang to complete this task.
- Feel free to add functionality as you have time, but the feature described above is the priority.
- Please add tests

## What we're looking for
- The functionality of the project matches the description above
- An ability to think through all potential states
- In the README of the project, describe exactly how to run the application

When complete, please push your code to Github and send the link to the project or zip the project (including the `.git` directory) and send it back.

Thank you and please ask if you have any questions!