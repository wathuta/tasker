A cli task manager that is able to add ,list and delete tasks.

it uses the following comands
    Tasker "[command]"
            list
            do
            add

the application uses boltdb

add command takes arguments which are then stored as tasks in the database

list command takes no argument 
it lists all the tasks in the database

do takes arguments that are the ids of the tasks.The command is responsible for deleting the tasks in the positions stated in the arguments  

to run the app change the current directory on the terminal to the project directory and type this command
            " $ make install"