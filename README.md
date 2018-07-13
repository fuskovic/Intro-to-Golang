# Golang fundamentals
by Faris Huskovic


This repo is intended to be used for educational purposes 

Below is a table of contents of the different subjects and examples covered.

This repo came about as a result of me watching [this](https://www.youtube.com/playlist?list=PLq9Ra239pNZC0MgMN4j6ZiPHv_c0UPnBX) playlist. I compiled all my notes and examples together so I could always come back to it as a point of reference. Maybe you can too one day.


# Table of Contents

## variables

Example 1-Variables- declaring and converting them

## Primitives

Example 1- Booleans, integar types and basic arithmetic

Example 2- Arithmetic between different int types and bit operators

Example 3- Complex Integer types, strings, and runes

## Constants

Example 1- Constant fundamentals

Example 2- Using consts to check for equality

Example 3- Using iota and common use case example

## Arrays and Slices

Example 1- Arrays in Go

Example 2- Syntax for Arrays of Arrays/ multi-dimensional arrays

Example 3- Creating a copy Vs pointing to address of

Example 4- Slices- their indices and bounds

Example 5- Appending and concatenating slices

Example 6- Slices-removing inner indices by appending

Example 7- Creating slices with make function

Example 8- Differences between Arrays and Slices

## If and Switch

Example 1- Logical operators and short-circuiting

Example 2- Switch statements- tag and tagless syntax and fallthrough

Example 3- Type switching and type switching w/ Arrays

## Looping

Example 1- Multi-variable looping

Example 2- Leaving out the counter variable

Example 3- leaving out the incrementer statement (do while effect), infinite loop, and continue

Example 4- Breaking from nested loops using labels and "for ranging over collections"

## Funcs and Methods

Example 1- Scope and passing multiple parameters

Example 2- Return values and passing slices

Example 3- Anonymous funcs and funcs as types

Example 4- Methods

## Defer,panic, and recover

Example 1- Understanding defer and LIFO

Example 2- Closing response resources when making an API call

Example 3- How values are passed into deferred functions

Example 4- Order of execution

Example 5- Handling your own errors

Example 6- Using recover to save application execution

## Pointers

Example 1- Basics of pointers, address of, and dereferencing

Example 2- Using pointers, &a, and dereferencers with structs

Example 3- Underlying value relationships of arrays and slices

## Maps and Structs

Example 1- Making maps and returning values by key

Example 2- Appending and deleting key/value pairs from Maps

Example 3- Using bools to check if keys are present in a map

Example 4- Building structs and returning their values

Example 5- Anonymous structs and using pointers with structs

Example 6- Composition and embedding plus adding tags to structs

## Interfaces

Example 1- Basic Interface

Example 2- Pointer Receiver VS Value Receiver with Interfaces

Example 3- Composing Interfaces

Example 4- Any custom types with methods can use interfaces

Example 5- Type switching empty interfaces to test types

## Goroutines

Example 1- Using a go routine to invoke a named function

Example 2- Goroutines and anonymous funcs

Example 3- Passing data into a go routine

Example 4- Synchronizing goroutines with WaitGroups

Example 5- Synchronizing using mutex locks to run in parallel

Example 6- Manipulating runtime threads


## Channels

Example 1- Using sending and receiving channels

Example 2- Understanding Channel deadlock

Example 3- Making Channel data flows more clear

Example 4- Using buffered channels-in case your sender or receiver
needs more time to process but you don't want to block the other side.

Example 5- Ranging over channels

Example 6- For looping through channels to use bools to check if they're open first

Example 7- Passing structs into a channel and logging channel status


