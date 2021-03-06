## Days Between

### What is it?
This app counts days between two dates given.

Built within the following constraints:

This is required to be a CLI application implemented in golang using only primitives, built-in language functions and no external npm dependencies. You should build the application so it accepts input from stdin. Please do guide the user through the steps required to use your application, and do validate your input.

You need to calculate the distance in whole days between two dates, counting only the days in between those dates, i.e. 01/01/2001 to 03/01/2001 yields “1”. The valid date range is between 01/01/1900 and 31/12/2999, all other dates should be rejected.
When testing your solution, use the following sample data to demonstrate your code works:
- a) 2/6/1983 to 22/6/1983 19 days
- b) 4/7/1984 to 25/12/1984 173 days
- c) 1/3/1989 to 3/8/1983 2036 days
  
### How to run:
The Easiest way is to download the binary that matches your system architecture and run from the command line.

__Usage:__

Linux/MacOS:
```bash
./days_between 2/6/1983 to 22/6/1983
```
Windows:
```bash
days_between.exe 2/6/1983 to 22/6/1983
```

Run code on your local machine:
```bash
git clone git@github.com:petelah/days_between.git
cd days_between
go run cmd/app/main.go 2/6/1983 to 22/6/1983
```
Replace the two dates above with dates of your choosing.

### Run tests:
Run from the project directory
```bash
go test ./cmd/app -v
```

### CI testing
Continuous integration tests are run upon pushing to master branch.
The test suite will run tests on:
- Date constraints 
- - Year between 1900 & 2999
- - Months between 1 & 12
- - Days between 1 & 31
- Leap years

Tests will run on every push to master.

Tests will also run and before releases and releases being created will depend on test suite passing.
  

### Improvements/todo
- ~~Add validation to check if the first date is before the second date.~~
- ~~Error handling.~~
- Multistage docker build so it can be run independent of OS (I mean even though golang is cross platform, but why not?).
- ~~Reduce complexity of overall application.~~
- ~~Unit tests.~~
- ~~Create releases.~~
- Add linting checks to ci pipeline
