#!/bin/sh

platform=$(gum choose "Codeforces" "Custom")
echo "Platform choosen > $(gum style --foreground 212 $platform)"

add_to_gitignore() {
    if grep -F -q "$1/" .gitignore; then
        echo "Already added to gitignore"
    else
        echo "$1/" >> .gitignore
    fi
}

case $platform in
    "Codeforces")
        contestId=$(gum input --placeholder "Enter contest id")
        cd golang
        go run /home/ankit/CP/golang/script.go $platform $contestId &
        BG_PID=$!
        while kill -0 $BG_PID 2>/dev/null; do
            gum spin -s line --title "work in progress..." -- sleep 1
        done
        cd ..
        add_to_gitignore "codeforces_contest_$contestId"
    ;;
    "Custom")
        folderName=$(gum input --placeholder "Enter folder name")
        echo "Folder name is $(gum style --foreground 212 $folderName)"
        number=$(gum input --placeholder "Enter number of files to create")
        echo "Number of files is $(gum style --foreground 212 $number)"
        cd golang
        go run /home/ankit/CP/golang/script.go $platform $folderName $number
        cd ..
        add_to_gitignore $folderName
    ;;
esac

gum style --foreground '#FE987B' "(>‿◠)✌ Yay! Done"

editor_action() {
    cmd=$1
    case $platform in
        "Codeforces")
            $cmd codeforces_contest_$contestId
        ;;
        "Custom")
            $cmd $folderName
        ;;
    esac
}

if gum confirm "Want to open editor?"; then
    editor=$(gum choose "Vs Code" "Vim" "Neovim")
    case $editor in
        "Vs Code")
            editor_action code
        ;;
        "Vim")
            editor_action vim
        ;;
        "Neovim")
            editor_action nvim
    esac
fi

