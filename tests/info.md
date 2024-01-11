Yes... I do keep my tests in this folder. You might ask why. You might disagree. But I like it this way. 

By the Golang standard, tests should be in the same folder as the code they are testing. It makes sense, as the tests - being in the same module - have access to the private functions and variables of the code they are testing. However I like to keep my tests seperate from the code. I like to keep my code, and more importantly, my file structure clean and tidy, and I like to keep my tests clean and tidy. 

When I had the tests in the same folder as the code, I found that I was constantly opening the test files by mistake, and vice versa. I was constantly paying too much attention to navigating files, which is not what I want to be doing. I want to be writing code.

Therefore, I keep my tests in a seperate folder.