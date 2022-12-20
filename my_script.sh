#!/bin/sh

folderName=$(gum input --placeholder "Enter folder name")
echo "Folder name is $(gum style --foreground 212 $folderName)"
number=$(gum input --placeholder "Enter number of files to create")
echo "Number of files is $(gum style --foreground 212 $number)"

gum spin -s line --title "work in progress..." -- sleep 1
go run /home/ankit/CP/golang/script.go $folderName $number

if gum confirm "Want to open editor?"; then
    gum style --border normal --margin "1" --padding "1 2" --border-foreground 212 "Choose Your Editor"
    
    editor=$(gum choose "vs code" "vim")
    
    case $editor in
        "vs code")
            code $folderName
        ;;
        "vim")
            vim $folderName
        ;;
    esac
fi


clear