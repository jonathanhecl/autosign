# AutoSign
_AutoSignature Executable in Golang_

### Times on Windows:
* 300~ms create a signed file
* 4~ms verify a signed file
### Times on Linux:
* 40~ms create a signed file
* 4~ms verify a signed file

## Start using it
 1.  Download and install it:
 ```
    $ go get github.com/jonathanhecl/autosign
 ``` 
 2.  Import it in your code:
 ```
    import "github.com/jonathanhecl/autosign"
 ``` 
 3. Apply it:
 ```
    hash := autosign.Init()
 ``` 
Or you can add a Salt with:
 ```
    hash := autosign.Init([]byte("salt"))
 ``` 

Also you can use the hash to verification of the version on internet or you can use it like "tokenCode" in a package like my other project [controlrun](https://github.com/jonathanhecl/controlrun)

## Screenshots

### Running
![Running](https://i.imgur.com/MSjuj1I.png)

### AutoSign
![AutoSign](https://i.imgur.com/TjRdmnq.png)