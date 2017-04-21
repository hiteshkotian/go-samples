To compile the source code first set GOROOT variable.
     export GOROOT=<<Replace_With_Path_of_the_project_on_Filesystem>>

First we build the library which is in package hitesh/stringutil.
     go install hitesh/stringutil.
This will create a pkg folder with the file.

Now we execute "go install" on the main file that was created.
     go install hitesh/hello

After the compilation, a bin directory is created with the executable called "hello".
Running the executable using the following command :
     ./bin/hello
