# SumUp challenge.

## What I'd do different (on my implementation):

1. Return ResponseJSONError (with error description encoded in JSON struct) rather than plain HTTP Method error (no
   body)
2. Maybe more detailed responses back to user to highlight what exact value is incorrect
3. Also utilize logging library for application audit during runtime
4. There's little sense to continue verify ticks once we found a fruit, but provided TCs (e.g., "
   TestValidate_InvalidMoveBackwards") seems expects to verify all ticks despite found fruit
5. I was in doubt do I need to create my own unit tests with this solution or rely on the existing ones (and thus
   duplicate them). It would be good to clearly specify it in the task description. For local development/testing
   purposes I used my IDE and curl and they worked fine to me.

## Notes

There could be a bug in Test Case "TestValidate_5by5Success". Probably TC used pointer to old (source) snake struct to
compare fields of returned JSON. Below in TC outcome "expected" is what my handler returned and "actual" is what Test
Case expected.

```
   === RUN   TestValidate_5by5Success
   main_test.go:263:
   Error Trace:	main_test.go:263
   Error:      	Not equal:
   expected: main.state{GameID:"1", Width:5, Height:5, Score:1, Fruit:main.fruit{X:4, Y:0}, Snake:main.snake{X:0, Y:0, VelX:1, VelY:0}}
   actual  : main.state{GameID:"1", Width:5, Height:5, Score:1, Fruit:main.fruit{X:4, Y:0}, Snake:main.snake{X:4, Y:4, VelX:0, VelY:1}}

        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -10,6 +10,6 @@
        	            	  Snake: (main.snake) {
        	            	-  X: (int) 0,
        	            	-  Y: (int) 0,
        	            	-  VelX: (int) 1,
        	            	-  VelY: (int) 0
        	            	+  X: (int) 4,
        	            	+  Y: (int) 4,
        	            	+  VelX: (int) 0,
        	            	+  VelY: (int) 1
        	            	  }
        	Test:       	TestValidate_5by5Success
--- FAIL: TestValidate_5by5Success (0.00s)
```

This could be verified by running POST request with the same data, so backend will return correct expected result and
thus the issue with TC:

```
curl -i --header "Content-Type: application/json" --request POST --data '{"gameId":"1","width":5,"height":5,"score":0,"fruit":{"x":4,"y":4},"snake":{"x":0,"y":0,"velX":1,"velY":0},"ticks":[{"velX":1,"velY":0},{"velX":1,"velY":0},{"velX":1,"velY":0},{"velX":0,"velY":1},{"velX":0,"velY":1},{"velX":0,"velY":1},{"velX":1,"velY":0},{"velX":0,"velY":1}]}' http://localhost:8080/validate`

HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 04 Mar 2021 14:00:51 GMT
Content-Length: 107

{"gameId":"1","width":5,"height":5,"score":1,"fruit":{"x":3,"y":0},"snake":{"x":0,"y":0,"velX":1,"velY":0}}%
```

So following the logic of this test case here is what TC actually expected for Snake struct (with this set up in JSON
response test case is Passed):
```snake{X: gs.Fruit.X, Y: gs.Fruit.Y, VelX: 0, VelY: 1}```
This is pretty confusing as from task description "The snake will always start at position (0, 0), with a velocity of (
1, 0)" and my initial set up was:
```Snake: snake{VelX: 1} // other fields are initialized by their default values```

Here is the commit for this fix 02a1ee2c16e3d236b233ce49a87fbe5f68714ed1

# Test runner

Test runner has plenty of hidden options, why not highlight them on task description? For example:

* run in verbose mode (to see which TCs are passed and which are not)
  ` ./run-test_linux_amd64 -test.v`
* display all available keys (there are many of them!)
  ` ./run-test_linux_amd64 --help`
* List failed tests:
  ` ./run-test_linux_amd64 | grep "FAIL"`
* List passed tests:
  ` ./run-test_linux_amd64 -test.v | grep "PASS"`
* Run failed test (specify the TC name after -test.run)
  `  ./run-test_linux_amd64 -test.v -test.run "TestValidate_InvalidMoveDiagonal"`
* Or even like this :-)
  `./run-test_linux_amd64 -test.list *`

``` 
== PLEASE INCLUDE THIS WITH YOUR SUBMISSION. ==
Well done! Your code is: [hidden]
== PLEASE INCLUDE THIS WITH YOUR SUBMISSION. ==
```