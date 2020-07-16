### To do:
- Tidy up the project files, for example:
	- Separate the request/response object from entities folder
	- Separate the API from main.go, so that it is testable
	- Find better file names
- How to assert nil value? For example take a look at AtCoderFilterParameter test.
- Refactor the handlers. There are a lot of repeating code in compare func.
- Find a better way to render CompareResponse.
- Add logs
- Add metrics
- Deploy to Heroku? Or App Engine
- Integration tests
- Is it okay to define a struct inside a func? For example look at codeforces repository
- Refactor the repo package.
- Is it really necessary to separate AtCoderFilterParameter and CodeforcesFilterParameter?
- Define error wrapper? For example look at util/request.go. Maybe it is better
to compare error object than error message.
- Find a better way to perform validation
- Find a better way to handle unmarshal error
- Gitlab CI. Automate the deployment.
