# mp2
Implementation of a distributed file system. We use UDP to detect failures and TCP to execute the file commands.

## Installation

Build and deploy with

```bash
**Make all**
```

## Warning

Wirte your own makefile to deploy. The original content in makefile is only supported during the course.

## Usage

To run on local host:

```bash
** go run main.go debug **
```

To run across machines:
```bash
** go run main.go **
```

## Commands
To deploy a intro node
```bash
** join intro**
```

To deploy a ordinary node
```bash
** join [port_number]**
```

To leave the system
```bash
** leave **
```
To get node info
```bash
** id **
```
To list members of the system
```bash
** members **
```
To put a file
```bash
** put [local_filepath] [filename] **
```

To get a file
```bash
** get [filename] [local_filepath] **
```

To delete a file
```bash
** delete [filename] **
```

To get a list of servers with a  file
```bash
** ls [filename] **
```

To get a list of files on the server
```bash
** store  **
```

To run maple(map) command
```bash
** maple <exec_filename> <number_of_maple> <sdfs_intermediate_file_prefix> <input_filename> **
```

To run juice(reduce) command
To get a list of files on the server
```bash
** juice <exec_filename> <number_of_juice> <sdfs_intermediate_file_prefix> <output_filename> <delete_intermediate_file_or_not> <method_of_dividing_task(hash or not)> **
```
See example map_reduce command in "command_backup". See example executable file and source code file in "/mj_exe"