#!/bin/bash
printerError() {
    error_counter=0
    allowed_letters=('a' 'b' 'c' 'd' 'e' 'f' 'g' 'h' 'i' 'j' 'k' 'l' 'm')
    read -a letters <<< "$(echo $1 | sed 's/./& /g')"

    for letter in "${letters[@]}"; do
        is_found=false
        for allowed_letter in "${allowed_letters[@]}"; do
            [ $letter == $allowed_letter ] && is_found=true && break
        done
        [ $is_found != true ] && ((error_counter++))
    done

    echo "$error_counter/${#letters[@]}"
}

printerError $1