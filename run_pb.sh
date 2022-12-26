file=$2
contest=$1
color="#FEC260"
color2="#E94560"

cd codeforces_contest_$contest/$file

if g++ $file.cpp -o $file;then
    a=0
    count=`ls -l input_*.txt | wc -l`
    while [ $a -lt $count ]
    do
        (./$file < input_$a.txt) > output_test_$a.txt
        if diff -q output_$a.txt output_test_$a.txt; then
            gum style --foreground  $color "[✓] Test $(gum style --foreground 212 "#$a") $(gum style --foreground $color "Passed")"
        else
            gum style --foreground $color2 "[•] Test $(gum style --foreground 212 "#$a") $(gum style --foreground $color2 "Failed")"
            gum style --foreground 99 "Expected Output:"
            gum style --border normal --border-foreground 212 --padding '0 1' --margin '0 1' $(cat output_$a.txt)
            gum style --foreground 99 "Your Output:"
            gum style --border normal --border-foreground 210 --padding '0 1' --margin '0 1' $(cat output_test_$a.txt)
        fi
        a=`expr $a + 1`
    done
else
    echo "Failed to Compile $file.cpp"
fi
