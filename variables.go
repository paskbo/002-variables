package main

import "fmt"

func main() {
  var test int
   test = 1
   fmt.Println(test)
   test = 2
   fmt.Println(test)
   ExampleLol := "Chvfefe"
   // fmt.Println(ExampleLol)
   /*In line 11, it replaces `var YourName string` to something like `x := "y"`
   And since := assigns, for example in line 11, ExampleLol as a "Chvfefe",
   don't have to re-type :=, since the string has been already assigned.

   If we want to change it's value, we simply type =. For example:
   Oh, and don't forget to print it again, like in line 20. I forgot it already.
   Edit 2: And remove line 12, so it doesn't print the value two times.*/
   ExampleLol = "Chvfefe 2.0"
   fmt.Println(ExampleLol)

   /* Give it a run on cmd (go run variables.go), and see it for yourself!
   lol :)

   Another reminder, since ExampleLol has been declared as a string,
   if you make any mistakes while typing, the cmd will tell you that
   ExampleLol is a stringtype. Because, you declared it as a string,
   in line 11. Good to know, lol!

   Short reminder to myself:
   Always change the variable itself, like ExampleLol2, and give it a new string,
   if you simply want a new string, with a new variable.

   Only use line 7, 9 or 20 as an example, if you want to change the variable value.
   It can be confusing at times, but it's very basic.

   Now, if we want to declare the variable from line 9 and line 20, here's an example:
   */

   ExampleLol, test = "Chvfefe is always", 100
   fmt.Println(ExampleLol, test)

   NewVariableHaha := "Amazing new variable wow!"
   ExampleLol, NewVariableHaha = NewVariableHaha, ExampleLol
   fmt.Println(ExampleLol, NewVariableHaha, test)
   /* Lol, I think I get it now from time to time. It's kinda cool!*/
   TheFourthExample := "This should comes at the end, wow it's getting confusing"
   fmt.Println(test, NewVariableHaha, ExampleLol, TheFourthExample)

  VarIntroduction := "So, I basically changed ExampleLol to NewVariable, and NewVariable to ExampleLol."
  TheFourthExample = "So, I could basically add more and more variables, but decided that this will be the last text. Oh, I btw just changed the value from this one aswell, gosh I'm so cool. Vote for chvfefe"

  fmt.Println(test, ExampleLol, NewVariableHaha, test, VarIntroduction, TheFourthExample)
 }
/*Oh, and something to add before closing this tab and forgetting everything,
you need to use everything you type in. If you declare something without using it,
go will try to throw out an error, because they want you to remove it.
It's good and bad, but I think we should take advantatge of that.
Specially because in JavaScript, I was kinda lost.*/
