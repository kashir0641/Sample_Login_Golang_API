# Login Application

Login project contains four API's :
1. <b>Register API</b> is for first time user registration with four fields. All four fields are compulsory.
2. <b>Login API </b> is POST API first validate request fields and then check user available in database otherwise throughs error message. If user exists in DB then create one access-token and store that token into cookies with expiry time.
3. <b>Logout API </b> is POST API which delete the user cookies from there Browser/Mobile.
4. <b>Update API </b> validate user requests first and then update user details to database. And this is POST API.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Whats is JWT ?
<b>JSON Web Token (JWT)</b> is an open standard for securely transmitting information between parties as a JSON object.
JWT is commonly used for authorization. JWTs can be signed using a secret or a public/private key pair. Once a user is logged in, each subsequent request will require the JWT, allowing the user to access routes, services, and resources that are permitted with that token.

### Installation golang

<br/>Install Go with [homebrew](https://brew.sh/):

```Shell
sudo brew install go
```

with [apt](https://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

[install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)


<br/>Install packages from github to my gopath/
```Shell
go get -u github.com/gorilla/mux
```

### Installation MongoDB
<b>Installation Manual link</b> : https://docs.mongodb.com/manual/installation/

#####COMPONENTS

    mongod - The database server.
    mongos - Sharding router.
    mongo  - The database shell (uses interactive javascript).

#####UTILITIES

    install_compass   - Installs MongoDB Compass for your platform.

#####BUILDING

    See docs/building.md.

#####RUNNING

  For command line options invoke:

    $ ./mongod --help

  To run a single server database:

    $ sudo mkdir -p /data/db
    $ ./mongod
    $
    $ # The mongo javascript shell connects to localhost and test database by default:
    $ ./mongo
    > help
    
    
## Usage


####Login instruction
<b>Step One</b> — Get The users name and Password from postman or client
<br/><b>Step Two</b> — Generate access token with help of secret key. For best practice secret key stores into bash_profile files.

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b login-module`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin login-module`
5. Submit a pull request.

## go-testing

Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the “go test” command, which automates execution of any function of the form
        
    $func TestXxx(*testing.T)

The test directory contains tests of the Go tool chain and runtime.
It includes black box tests, regression tests, and error output tests.

A simple test function looks like this:

    func TestAbs(t *testing.T) {
        got := Abs(-1)
        if got != 1 {
            t.Errorf("Abs(-1) = %d; want 1", got)
        }
    }

To run just these tests, execute:

    $ go test -run NameOfTest
    $ go test -run Test_Login
    
   Some Commands for test files -
    
    go test -run ''      # Run all tests.
    go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
    go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
    go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
   

Standard library tests should be written as regular Go tests in the appropriate package.

The tool chain and runtime also have regular Go tests in their packages.
The main reasons to add a new test to this directory are:

* it is most naturally expressed using the test runner; or
* it is also applicable to `gccgo` and other Go tool chains.

