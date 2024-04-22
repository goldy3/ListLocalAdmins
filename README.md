# ListLocalAdmins

This repository contains a Go program designed to manage Windows server administrator groups. The program reads server names from a text file, executes a PowerShell command on each server to list user accounts in the local admin group, and outputs the results to a CSV file.

## Features

- Read server names from a text file.
- Execute PowerShell commands remotely.
- Export results to a CSV file for easy analysis.

## Getting Started

### Prerequisites

- Go (Golang) installed on your machine.
- PowerShell accessible on the target Windows servers.
- Network permissions to execute remote commands.

### Installation

Clone the repository to your local machine:

bash
git clone https://github.com/yourusername/ListLocalAdmins.git
cd ListLocalAdmins

## Usage
Ensure you have a text file containing the names of the servers you wish to query, separated by commas.
Run the program using:
bash
Copy code
go run main.go path_to_your_input_file.txt path_to_your_output_file.csv
Replace path_to_your_input_file.txt with the path to your input file, and path_to_your_output_file.csv with the desired path for the output file.
