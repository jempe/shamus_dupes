# shamus_dupes

This is a utility to find duplicate files in a directory tree. It is written in Bash and golang. The Bash script creates a script that needs to be run in the directory tree to be checked. it creates a script that generates a text file with all the hashes sorted by hash. 

The golang program reads 2 different files and compares the hashes to find duplicates. It then prints out the duplicates. or creates a script to delete the duplicates.

## Usage

Go to the directory that has the files you want to keep. Run the bash script.

```bash
create_shasums_script.sh > get_sums.sh
```

Then run the script that was created.

```bash
bash get_sums.sh
```

You will now have a file called `shasums.txt` in the directory you ran the script in.

Now go to the directory that you suspect has duplicates. Run the bash script.

```bash
create_shasums_script.sh > get_sums.sh
```

Then run the script that was created.

```bash
bash get_sums.sh
```

You will now have a file called `shasums.txt` in the directory you ran the script in.

Now run the golang program to find the duplicates.

```bash
shasums_duplicates remove FilesToRemove/shasums.txt FilesToKeep/shasums.txt > remove_duplicates.sh
```

Inspect the `remove_duplicates.sh` script to make sure it is doing what you want.

Then run the script that was created.

```bash
bash remove_duplicates.sh
```

This will remove the duplicates from the FilesToRemove directory.





