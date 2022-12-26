file=$2
contest=$1

cd codeforces_contest_$contest/$file

if g++ $file.cpp -o $file;then
    a=0
    count=`ls -l input_*.txt | wc -l`
    while [ $a -lt $count ]
    do
        (./$file < input_$a.txt) > output_test_$a.txt
        if diff -q output_$a.txt output_test_$a.txt; then
            echo "Test $a Passed"
        else
            echo "Test $a Failed"
        fi
        a=`expr $a + 1`
    done
else
    echo "Failed to Compile $file.cpp"
fi
