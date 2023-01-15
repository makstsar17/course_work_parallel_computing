# CourseWork Parallel Computing
## Installation
1. Clone the repository.
```bash
git clone git@github.com:makstsar17/course_work_parallel_computing.git
```
2. Run server:

- Change the directory.
```bash
cd course_work_parallel_computing\server
```
- Create new directory
```bash
mkdir data
```
Add all documents to this directory, so that inverted index has data to work on.

- Run server with parallel or consecutively execution.
```bash
go run main.go -thr=8 -pe=true
```
Flags:

* thr - the number of goroutines(defaults to 10)
* pe - parallel build (default to true)

3. Run Client.

- Change the directory.
```bash
cd course_work_parallel_computing\client
```
- Run client.
```bash
go run main.go
```

