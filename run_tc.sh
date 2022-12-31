contest=$1
pb=$2
color="#FEC260"
color2="#E94560"

cd $contest/$pb

g++ $pb.cpp -o $pb &

cp_pid=$!

while kill -0 $cp_pid 2>/dev/null; do
    gum spin --spinner minidot --title "compiling $pb.cpp" --title.foreground 99 -- sleep 1
done

if [ -x $pb ]; then
    gum style --foreground 69 " ✓ Compiled Succssfully"
    a=0
    count=`ls -l input_*.txt | wc -l`
    while [ $a -lt $count ]
    do
        (./$pb < input_$a.txt) > output_test_$a.txt
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
    rm $pb
else
    echo "failed to compile $pb.cpp. Try Again!"
fi
