# Advent-of-Code-2024
Advent of Code is a yearly returning series of `25` (`50`) tasks created by `Eric Wastl`. **Credits see below**.

Every day in December, till 25., a new task is released. If you manage to solve this task you receive **one star** and a part two of this task, that adds a little extra question to the first part, is unlocked. If you manage to solve the additional part two aswell you earn a **second star**. So you can earn a **total number of 50 stars** and complete 25 tasks with an extra question to each. Each day the tasks get a little bit harder.

You can solve the tasks with whatever language you like or like to learn. To me this project offers a great opportunity to take a look into other programming languages and learn new skills.

This repository shows my solutions to the tasks. I also added the tasks / questions to each day as `README.md`. So everybody can follow the code made and check it against the task. Given data by the task is always stored inside the day directory. Default `.txt` except any other format is given by the task.

## Language
This year I looked into Go and tried to solve the tasks using this language. I was pretty new to Go so let's see where this ride will go.

## Credits
All credits of the tasks and questions go to `Eric Wastl` (Twitter: `@ericwastl`)
See his Website: [Advent of Code](https://adventofcode.com/) for further information.

## Structure
Every day gets a new directory with an App file inside. This file contains the solutions to both tasks of the day in separate functions. The Main function would run both tasks of this day and print out the solutions. If data is given by the task it will also be stored inside this directory. The tasks of the day are also given as README.md in the day directory. Tests are provided in every directory and can be executed with the go test command. In the Utils directory some helper functions are stored that can be used by every day.

You can execute all days using the App given in the main directory and entering the number of the day you want to execute.