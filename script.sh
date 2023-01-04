#!/bin/sh

platform=$(gum choose "Atcoder" "Codeforces" "Codechef" "Custom")
echo "Platform choosen > $(gum style --foreground 212 $platform)"

# Add to gitignore
add_to_gitignore() {
    if grep -F -q "$1/" .gitignore; then
        echo "Already added to gitignore"
    else
        echo "$1/" >> .gitignore
    fi
}

# Open editor
editor_action() {
    case $platform in
        "Atcoder")
            $1 $contestId 
        ;;
        "Codeforces")
            $1 codeforces_contest_$contestId
        ;;
        "Custom")
            $1 $folderName
        ;;
        *)
            $1 $contest
        ;;
    esac
}

loader_action() {
    while kill -0 $1 2>/dev/null; do
        gum spin --spinner minidot --title "work in progress..." --title.foreground 99 -- sleep 1
    done
}

case $platform in
    "Atcoder")
        contestId=$(gum input --placeholder "Enter contest id")
        /home/ankit/development/go/atcoder_crawler/atcoder_crawler --contest=$contestId --dir=pwd &
        ATCODER_PID=$!
        loader_action $ATCODER_PID
        add_to_gitignore $contestId
    ;;
    "Codeforces")
        contestId=$(gum input --placeholder "Enter contest id")
        cd golang
        go run script.go --platform=$platform --contestId=$contestId &
        CODEFORCES_PID=$!
        loader_action $CODEFORCES_PID
        cd ..
        add_to_gitignore codeforces_contest_$contestId
    ;;
    "Codechef")
        contest=$(gum input --placeholder "Enter Codechef contest id Id e.g START71B")
        python ~/codechef/codechef_cp_helper/main.py -c contest &
        CODECHEF_PID=$!
        loader_action $CODECHEF_PID
        add_to_gitignore $contest
    ;;
    "Custom")
        folderName=$(gum input --placeholder "Enter folder name")
        echo "Folder name is $(gum style --foreground 212 $folderName)"
        number=$(gum input --placeholder "Enter number of files to create")
        echo "Number of files is $(gum style --foreground 212 $number)"
        cd golang
        go run script.go --platform=$platform --name=$folderName --count=$number
        cd ..
        add_to_gitignore $folderName
    ;;
esac

gum style --foreground '#FE987B' "(>‿◠)✌ Yay! Done"

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

