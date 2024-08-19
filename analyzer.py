import os
import subprocess

# Folder path
input_path = ".\\generated"

# Shell command to execute for each file
# shell_command = ".go-lisa\go-lisa-0.1\bin\go-lisa -i {} -o .\output\test -a {}"

# The analysis to perform
analysis = ['parity', 'intervals', 'prefix', 'suffix'] # no taint now since requires additional args

blacklist = []

# Iterate over each file in the folder
for file_name in os.listdir(input_path):
    # Construct the full path to the file

    file_path = os.path.join(input_path, file_name)
    # Check if it's a file (not a directory)
    if os.path.isfile(file_path) and file_name not in blacklist:
        print(f"PROCESSING {file_path}...")
        for an in analysis:
            # full_command = shell_command.format(file_path, an)
            # res = os.system(full_command)
            res = subprocess.run(['.\\go-lisa\\go-lisa-0.1\\bin\\go-lisa', '-i', file_path, '-o', f'.\\output\\{file_name[0:-3]}-{an}', '-a', an], stdout=subprocess.PIPE, shell=True).stdout.decode('utf-8')
            lines = res.split('\n')
            for line in lines:
                if line.startswith("Generated Warnings:"):
                    print(f"{file_name} with analysis {an} | {line}")