#!/bin/bash
function odd_or_even() {
  if [ $(($1%2)) == 0 ]; then
      echo 'Even'
  else
      echo 'Odd'
  fi
}

odd_or_even "$@"