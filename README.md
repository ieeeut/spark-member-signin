# spark-member-signin
This project will allow members to signin at IEEE meetings, check how many meetings they have attended, export data for 302 student attendance records, and allow easy email blasts to members / spark members.

# steps to install
First install Go from https://golang.org and install revel from https://revel.github.io, clone this project into your go/src folder, and run 'revel run spark-member-signin'

# note for me
Change the app secret before deploying to prod

**Slack Integration**

This repo will comment in the #github-projects branch whenever anything is committed or changed.


# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).