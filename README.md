** ANTON **

Name inspired from famous `Silicon Valley Series -> Gilfoyle invention`. I decided to start this project to brush up my skills with Go.

**_ Prerequisites _**

- Please make sure you have Go installed `1.20+`
- Please make sure your GOROOT and GOPATH are set correctly.

** About the Application **

We are going to build an AI (authentication Inspector), that is responsible for all the authentication related work. As most of the software companies spend time to
build the same stuff again and again and waste useful resources on this. With this they'll be able to focus on some more importants tasks.

`OUr basic idea is AI will handle`

- Basic Auth
- Google Login Auth
- Microsoft Login Auth

This application will be able to run independently without much effort. We'll make it as a dockerize container so that it is easy to pull and run asap.
We are going to use `.env` file to keep our variables and you can find the reference in `.example_env` file.

The Project will support webhooks as well, So our initial idea is that we take care of basic activities related to Auth of the user and each activity of the user will trigger an event webhook and you can configure the event webhook handler anywhere and based on event type, you can keep latest detailed of the user in your system.

We are going to use `gin` as our framework, so that it is easy to setup server and manage api routes.
