# GoReplace - search & replace tool

Search and replace tool is file oriented tool, it accept command line arguments for replacing one string with another. The source file(s) and the target folder where the replaced files would be copied are also command line parameters.   

## 1	Command line arguments

GoReplace command support and require multiple parameters to operate as expected. 

```
> goreplace {-s | -source}  [{-t | -target}] {-f | -find} {-r | -replace} 
```

You can mix order in which parameters are given according to a specific needs.

| Argument              | Description                                                  |
| --------------------- | ------------------------------------------------------------ |
| {-s } | { -source }       | The source folder and source filename - wild characters are allowed to limit source files |
| [{-t }] | [{ -target }]    | The destination folder, if not set the same files are saved back to the original source (**optional**) |
| {-f } | { -find }         | Searched string to change from (source)                      |
| {-r } | { -replace }      | String to replace to.                                        |

### 1.1	Examples

A few examples how to call the tool:  

```bash
goreplace -s E:\bisaga-blog\Documents\*.md -t E:\bisaga-blog\Temp -f "(images/" -r "(/images/" 
```





















