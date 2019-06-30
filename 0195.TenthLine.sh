# Read from the file file.txt and output the tenth line to stdout.
#awk 'NR == 10' file.txt
#sed -n 10p file.txt
#sed -n '10p' file.txt
tail -n+10 file.txt|head -1