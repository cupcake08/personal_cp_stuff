#!/bin/sh

platform=$(gum choose "codeforces" "custom")
echo "Platform choosen > $(gum style --foreground 212 $platform)"

case $platform in
    "codeforces")
        contestId=$(gum input --placeholder "Enter contest id")
        cd golang
        go run /home/ankit/CP/golang/script.go $platform $contestId &
        BG_PID=$!
        while kill -0 $BG_PID 2>/dev/null; do
            gum spin -s line --title "work in progress..." -- sleep 1
        done
    ;;
    "custom")
        folderName=$(gum input --placeholder "Enter folder name")
        echo "Folder name is $(gum style --foreground 212 $folderName)"
        number=$(gum input --placeholder "Enter number of files to create")
        echo "Number of files is $(gum style --foreground 212 $number)"
        cd golang
        go run /home/ankit/CP/golang/script.go $platform $folderName $number
    ;;
esac

gum style --foreground '#FE987B' "(>‿◠)✌ Yay! Done"

if gum confirm "Want to open editor?"; then
    editor=$(gum choose "vs code" "vim")
    cd ..
    case $editor in
        "vs code")
            case $platform in
                "codeforces")
                    code codeforces_contest_$contestId
                ;;
                "custom")
                    code $folderName
                ;;
            esac
        ;;
        "vim")
            case $platform in
                "codeforces")
                    vim codeforces_contest_$contestId
                ;;
                "custom")
                    vim $folderName
                ;;
            esac
        ;;
    esac
fi
clear